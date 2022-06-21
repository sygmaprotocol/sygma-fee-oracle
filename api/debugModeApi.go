// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"fmt"
	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
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
	resourceTokenAddr := c.Param("address")

	endpointRespData := &FetchRateResp{
		BaseRate:                 "0.000445",
		TokenRate:                "8.948864",
		DestinationChainGasPrice: "9000000000",
		FromDomainID:             fromDomainID,
		ToDomainID:               toDomainID,
		DataTimestamp:            time.Now().Unix(),
		SignatureTimestamp:       time.Now().Unix(),
		ExpirationTimestamp:      h.dataExpirationManager(time.Now().Unix()),
		Debug:                    true,
		ResourceID: resourceTokenAddr[len(h.conf.GetRegisteredDomains(fromDomainID).AddressPrefix):] +
			fmt.Sprintf("%02s", strconv.FormatInt(int64(fromDomainID), 16)),
	}
	endpointRespData.Signature, err = h.rateSignature(endpointRespData, fromDomainID, resourceTokenAddr, fromDomainID)
	if err != nil {
		ginErrorReturn(c, http.StatusInternalServerError, newReturnErrorResp(&config.ErrIdentityStampFail, err))
		return
	}

	ginSuccessReturn(c, http.StatusOK, endpointRespData)
}
