// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/types"

	"github.com/ChainSafe/sygma-fee-oracle/config"

	"github.com/ChainSafe/sygma-fee-oracle/consensus"
	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	conf *config.Config

	identity  *identity.OracleIdentityOperator
	consensus *consensus.Consensus

	gasPriceStore       *store.GasPriceStore
	conversionRateStore *store.ConversionRateStore

	log *logrus.Entry
}

// AddRouterPathsV1 registers api routers
func AddRouterPathsV1(v1RouterGroups map[string]*gin.RouterGroup, apiHandler *Handler, log *logrus.Entry) {
	apiHandler.log = log.WithField("api", "v1")

	if apiHandler.conf.AppModeConfig() == config.AppModeDebug {
		v1RouterGroups["rate"].GET("/from/:fromDomainID/to/:toDomainID/token/:address", apiHandler.debugGetRate)
	} else {
		v1RouterGroups["rate"].GET("/from/:fromDomainID/to/:toDomainID/token/:address", apiHandler.getRate)
	}
}

// endpoint: /{version}/rate/from/{fromDomainID}/to/{toDomainID}/token/{address}
// example for transferring native token:   /rate/from/1/to/2/token/0x0000000000000000000000000000000000000000 : from ethereum to polygon transfer eth    => ber = matic / eth, ter = matic / eth, gas price is from polygon
// example for transferring standard token: /rate/from/1/to/2/token/0x8A953CfE442c5E8855cc6c61b1293FA648BAE472 : from ethereum to polygon transfer doge   => ber = matic / eth, ter = matic / doge, gas price is from polygon
func (h *Handler) getRate(c *gin.Context) {
	fromDomainID, err := strconv.Atoi(c.Param("fromDomainID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid fromDomainID")))
		return
	}
	toDomainID, err := strconv.Atoi(c.Param("toDomainID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid toDomainID")))
		return
	}
	resourceTokenAddr := c.Param("address")

	h.log.Debugf("new request with params fromDomainID: %d, toDomainID: %d, tokenAddress: %s\n", fromDomainID, toDomainID, resourceTokenAddr)

	toDomain := h.conf.GetRegisteredDomains(toDomainID)
	if toDomain == nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid toDomainID")))
		return
	}
	fromDomain := h.conf.GetRegisteredDomains(fromDomainID)
	if fromDomain == nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid fromDomain")))
		return
	}

	resourceID := config.ResourceIDBuilder(resourceTokenAddr, fromDomainID)
	resource := h.conf.GetRegisteredResources(resourceID)
	if resource == nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid token address")))
		return
	}

	aggregatedGasPriceData, err := h.consensus.FilterLocalGasPriceData(h.gasPriceStore, toDomain.Name)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrConsensusFail, err))
		return
	}
	h.log.Debugf("aggregatedGasPriceData: %v\n", aggregatedGasPriceData)

	baseSymbol := toDomain.BaseCurrencySymbol
	foreignSymbol := fromDomain.BaseCurrencySymbol
	h.log.Debugf("base rate calculation: base: %s, foreign: %s\n", baseSymbol, foreignSymbol)

	aggregatedBaseRateData, err := h.consensus.FilterLocalConversionRateData(h.conversionRateStore, baseSymbol, foreignSymbol)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrConsensusFail, err))
		return
	}

	// if resourceID is the base currency of the fromDomain
	aggregatedTokenRateData := aggregatedBaseRateData
	if resource.Symbol != foreignSymbol {
		fromDomainResource := h.conf.GetRegisteredResources(resource.ID)
		if fromDomainResource == nil {
			ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput,
				errors.New("no resource registered with given the fromDomainID")))
			return
		}

		h.log.Debugf("token rate calculation: base: %s, foreign: %s\n", baseSymbol, fromDomainResource.Symbol)

		aggregatedTokenRateData = &types.ConversionRate{}
		if fromDomainResource.Symbol == baseSymbol {
			aggregatedTokenRateData.Rate = 1.0
		} else {
			aggregatedTokenRateData, err = h.consensus.FilterLocalConversionRateData(h.conversionRateStore, baseSymbol, fromDomainResource.Symbol)
			if err != nil {
				ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrConsensusFail, err))
				return
			}
		}
	}

	dataTime := aggregatedBaseRateData.Time
	signedTime := time.Now().Unix()

	endpointRespData := &FetchRateResp{
		BaseRate:                 fmt.Sprintf("%f", aggregatedBaseRateData.Rate),
		TokenRate:                fmt.Sprintf("%f", aggregatedTokenRateData.Rate),
		DestinationChainGasPrice: aggregatedGasPriceData.SafeGasPrice,
		FromDomainID:             fromDomainID,
		ToDomainID:               toDomainID,
		DataTimestamp:            dataTime,
		SignatureTimestamp:       signedTime,
		ExpirationTimestamp:      h.dataExpirationManager(dataTime),
	}

	endpointRespData.Signature, err = h.rateSignature(endpointRespData, fromDomainID, resource.TokenAddress, resource.DomainId)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrIdentityStampFail, err))
		return
	}

	endpointRespData.ResourceID = resourceTokenAddr[len(h.conf.GetRegisteredDomains(fromDomainID).AddressPrefix):] +
		fmt.Sprintf("%02s", strconv.FormatInt(int64(fromDomainID), 16))

	ginSuccessReturn(c, http.StatusOK, endpointRespData)
}
