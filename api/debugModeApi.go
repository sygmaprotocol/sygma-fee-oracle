// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"net/http"
	"strconv"
	"time"

	oracleErrors "github.com/ChainSafe/sygma-fee-oracle/errors"
	"github.com/ChainSafe/sygma-fee-oracle/util"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *Handler) debugGetRate(c *gin.Context) {
	fromDomainID, err := strconv.Atoi(c.Param("fromDomainID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&oracleErrors.InvalidRequestInput, errors.New("invalid fromDomainID")))
		return
	}
	toDomainID, err := strconv.Atoi(c.Param("toDomainID"))
	if err != nil {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&oracleErrors.InvalidRequestInput, errors.New("invalid toDomainID")))
		return
	}
	resourceID := c.Param("resourceID")

	msgGasLimitParam := c.DefaultQuery("msgGasLimit", "0")
	validValue := util.CheckInteger(msgGasLimitParam)
	if !validValue {
		ginErrorReturn(c, http.StatusBadRequest, newReturnErrorResp(&oracleErrors.InvalidRequestInput, errors.New("invalid msgGasLimit")))
		return
	}

	endpointRespData := &FetchRateResp{
		BaseRate:                 "0.000445",
		TokenRate:                "15.948864",
		DestinationChainGasPrice: "2000000000",
		FromDomainID:             fromDomainID,
		ToDomainID:               toDomainID,
		DataTimestamp:            time.Now().Unix(),
		SignatureTimestamp:       time.Now().Unix(),
		ExpirationTimestamp:      h.dataExpirationManager(time.Now().Unix() + 100000000),
		Debug:                    true,
		ResourceID:               resourceID,
		MsgGasLimit:              msgGasLimitParam,
	}

	endpointRespData.Signature, err = h.rateSignature(endpointRespData, fromDomainID, resourceID)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&oracleErrors.InternalServerError, err))
		return
	}

	ginSuccessReturn(c, http.StatusOK, endpointRespData)
}
