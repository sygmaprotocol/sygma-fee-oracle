// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ChainSafe/sygma-fee-oracle/util"

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
		v1RouterGroups["rate"].GET("/from/:fromDomainID/to/:toDomainID/resourceid/:resourceID", apiHandler.debugGetRate)
	} else {
		v1RouterGroups["rate"].GET("/from/:fromDomainID/to/:toDomainID/resourceid/:resourceID", apiHandler.getRate)
	}
}

// endpoint: /{version}/rate/from/{fromDomainID}/to/{toDomainID}/resourceid/{resourceID}?msgGasLimit={msgGasLimit}
// api call example:   /rate/from/0/to/1/resourceid/0x0000000000000000000000000000000000000000000000000000000000000001
// calculation of transferring native resource: from ethereum to polygon transfer eth    => ber = matic / eth, ter = matic / eth, gas price is from polygon
// calculation of transferring standard token : from ethereum to polygon transfer usdt   => ber = matic / eth, ter = matic / usdt, gas price is from polygon
func (h *Handler) getRate(c *gin.Context) {
	fromDomain, toDomain, err := h.parseDomains(c.Param("fromDomainID"), c.Param("toDomainID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid resourceID")))
	}

	resource, err := h.parseResource(c.Param("resourceID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid resourceID")))
	}
	h.log.Debugf("new request with params fromDomainID: %d, toDomainID: %d, resourceID: %s\n", fromDomain.ID, toDomain.ID, resourceID)

	msgGasLimitParam := c.DefaultQuery("msgGasLimit", "0")
	if util.CheckInteger(msgGasLimitParam) {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid msgGasLimit")))
		return
	}

	gp, err := h.consensus.FilterLocalGasPriceData(h.gasPriceStore, toDomain.Name)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrInternalServerError, err))
		return
	}
	h.log.Debugf("aggregatedGasPriceData: %v\n", gp)

	ter, err := h.consensus.FilterLocalConversionRateData(
		h.conversionRateStore,
		toDomain.BaseCurrencySymbol,
		fromDomain.BaseCurrencySymbol)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrInternalServerError, err))
		return
	}
	h.log.Debugf("base rate calculation: to: %s, from: %s\n", toDomain.BaseCurrencySymbol, fromDomain.BaseCurrencySymbol)

	aggregatedTokenRateData, err := h.calculateTokenRate(resource, ter, fromDomain, toDomain)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrInternalServerError, err))
		return
	}
	h.log.Debugf("token rate calculation: to: %s, from: %s\n", toDomain.BaseCurrencySymbol, resource.Symbol)

	dataTime := ter.Time
	signedTime := time.Now().Unix()
	endpointRespData := &FetchRateResp{
		BaseRate:                 fmt.Sprintf("%f", ter.Rate),
		TokenRate:                fmt.Sprintf("%f", aggregatedTokenRateData.Rate),
		DestinationChainGasPrice: gp.SafeGasPrice,
		FromDomainID:             fromDomain.ID,
		ToDomainID:               toDomain.ID,
		MsgGasLimit:              msgGasLimitParam,
		DataTimestamp:            dataTime,
		SignatureTimestamp:       signedTime,
		ExpirationTimestamp:      h.dataExpirationManager(dataTime),
	}
	endpointRespData.Signature, err = h.rateSignature(endpointRespData, fromDomain.ID, resource.ID)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrInternalServerError, err))
		return
	}

	endpointRespData.ResourceID = resource.ID
	ginSuccessReturn(c, http.StatusOK, endpointRespData)
}

func (h *Handler) parseDomains(fromID string, toID string) (from *config.Domain, to *config.Domain, err error) {
	fromDomainID, err := strconv.Atoi(fromID)
	if err != nil {
		return from, to, err
	}
	toDomainID, err := strconv.Atoi(toID)
	if err != nil {
		return from, to, err
	}
	if fromDomainID == toDomainID {
		return from, to, fmt.Errorf("from and to domain equal")
	}

	toDomain := h.conf.GetRegisteredDomains(toDomainID)
	if toDomain == nil {
		return from, to, fmt.Errorf("to domain not registered")
	}
	fromDomain := h.conf.GetRegisteredDomains(fromDomainID)
	if fromDomain == nil {
		return from, to, fmt.Errorf("from domain not registered")
	}

	return from, to, nil
}

func (h *Handler) parseResource(resourceID string) (*config.Resource, error) {
	if !strings.HasPrefix(resourceID, "0x") || len(resourceID) != 66 {
		return &config.Resource{}, fmt.Errorf("resource invalid")
	}

	resource := h.conf.GetRegisteredResources(resourceID)
	if resource == nil {
		return &config.Resource{}, fmt.Errorf("resource not registerred")
	}

	return resource, nil
}

func (h *Handler) calculateTokenRate(resource *config.Resource, ber *types.ConversionRate, from *config.Domain, to *config.Domain) (*types.ConversionRate, error) {
	ter := &types.ConversionRate{}
	// if resource is the base currency of the fromDomain
	if resource.Symbol == from.BaseCurrencySymbol {
		ter = ber
	} else {

		if resource.Symbol == to.BaseCurrencySymbol {
			ter.Rate = 1.0
		} else {
			ter, err := h.consensus.FilterLocalConversionRateData(h.conversionRateStore, to.BaseCurrencySymbol, resource.Symbol)
			if err != nil {
				return ter, err
			}
		}
	}

	return ter, nil
}
