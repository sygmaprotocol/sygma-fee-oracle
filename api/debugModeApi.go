// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/config"
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
	}

	endpointRespData.Signature, err = h.rateSignature(endpointRespData, fromDomainID, resourceID)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrIdentityStampFail, err))
		return
	}

	ginSuccessReturn(c, http.StatusOK, endpointRespData)
}
