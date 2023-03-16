// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/consensus"
	oracleErrors "github.com/ChainSafe/sygma-fee-oracle/errors"
	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/signature"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ChainSafe/sygma-fee-oracle/util"
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
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&oracleErrors.InvalidRequestInput, errors.New("invalid domainID")))
		return
	}

	resource, err := h.parseResource(c.Param("resourceID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&oracleErrors.InvalidRequestInput, errors.New("invalid resourceID")))
		return
	}
	h.log.Debugf("new request with params fromDomainID: %d, toDomainID: %d, resourceID: %s\n", fromDomain.ID, toDomain.ID, resource.ID)

	msgGasLimitParam := c.DefaultQuery("msgGasLimit", "0")
	if !util.CheckInteger(msgGasLimitParam) {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&oracleErrors.InvalidRequestInput, errors.New("invalid msgGasLimit")))
		return
	}

	gp, err := h.consensus.FilterLocalGasPriceData(h.gasPriceStore, toDomain.Name)
	if err != nil {
		h.log.Errorf("get gasprice process failed: %v", err)
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&oracleErrors.InternalServerError, err))
		return
	}
	h.log.Debugf("aggregatedGasPriceData: %v\n", gp)

	ber, err := h.consensus.FilterLocalConversionRateData(
		h.conversionRateStore,
		toDomain.BaseCurrencySymbol,
		fromDomain.BaseCurrencySymbol)
	if err != nil {
		h.log.Errorf("calculate ber process failed: %v", err)
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&oracleErrors.InternalServerError, err))
		return
	}
	h.log.Debugf("base rate calculation: to: %s, from: %s\n", toDomain.BaseCurrencySymbol, fromDomain.BaseCurrencySymbol)

	ter, err := h.calculateTokenRate(resource, ber, fromDomain, toDomain)
	if err != nil {
		h.log.Errorf("calculate ter process failed: %v", err)
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&oracleErrors.InternalServerError, err))
		return
	}
	h.log.Debugf("token rate calculation: to: %s, from: %s\n", toDomain.BaseCurrencySymbol, resource.Symbol)

	dataTime := ter.Time
	signedTime := time.Now().Unix()
	rate := &types.Rate{
		BaseRate:                 fmt.Sprintf("%f", ber.Rate),
		TokenRate:                fmt.Sprintf("%f", ter.Rate),
		DestinationChainGasPrice: gp.SafeGasPrice,
		FromDomainID:             fromDomain.ID,
		ToDomainID:               toDomain.ID,
		MsgGasLimit:              msgGasLimitParam,
		DataTimestamp:            dataTime,
		SignatureTimestamp:       signedTime,
		ExpirationTimestamp:      h.dataExpirationManager(dataTime),
	}
	rate.Signature, err = signature.RateSignature(h.conf, rate, h.identity, fromDomain.ID, resource.ID)
	if err != nil {
		h.log.Errorf("signature process failed: %v", err)
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&oracleErrors.InternalServerError, err))
		return
	}

	rate.ResourceID = resource.ID
	ginSuccessReturn(c, http.StatusOK, rate)
}
