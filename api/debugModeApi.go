// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/util"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *Handler) debugGetRate(c *gin.Context) {
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
	resourceID := c.Param("resourceID")

	resource := h.conf.GetRegisteredResources(resourceID)
	if resource == nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid resourceID, make "+
			"sure you config token address under the corresponding resource in resource.json")))
		return
	}

	msgGasLimitParam := c.DefaultQuery("msgGasLimit", "0")
	finalMsgGasLimit, err := util.MsgGasLimitChecker(msgGasLimitParam)
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&config.ErrInvalidRequestInput, errors.New("invalid msgGasLimit")))
		return
	}

	endpointRespData := &FetchRateResp{
		BaseRate:                 "0.000316",
		TokenRate:                "0.485081",
		DestinationChainGasPrice: "100000000000",
		FromDomainID:             fromDomainID,
		ToDomainID:               toDomainID,
		DataTimestamp:            time.Now().Unix(),
		SignatureTimestamp:       time.Now().Unix(),
		ExpirationTimestamp:      h.dataExpirationManager(time.Now().Unix() + 100000000),
		Debug:                    true,
		ResourceID:               resourceID,
		MsgGasLimit:              finalMsgGasLimit,
	}

	endpointRespData.Signature, err = h.rateSignature(endpointRespData, fromDomainID, resourceID)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrIdentityStampFail, err))
		return
	}

	ginSuccessReturn(c, http.StatusOK, endpointRespData)
}
