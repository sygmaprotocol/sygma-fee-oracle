// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package api

import (
	"errors"

	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/gin-gonic/gin"
)

type ReturnErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type ReturnSuccessResponse interface{}

func newReturnErrorResp(fe *config.FeeOracleError, err error) *ReturnErrorResponse {
	if err == nil {
		err = errors.New("")
	}
	return &ReturnErrorResponse{
		Code:    fe.Code,
		Message: fe.Message,
		Detail:  fe.Wrap(err).Error(),
	}
}

func ginErrorReturn(c *gin.Context, httpCode int, resp *ReturnErrorResponse) {
	c.JSON(httpCode, gin.H{
		"error": resp,
	})
}

func ginSuccessReturn(c *gin.Context, httpCode int, resp ReturnSuccessResponse) {
	c.JSON(httpCode, gin.H{
		"response": resp,
	})
}

type FetchRateResp struct {
	BaseRate                 string `json:"baseEffectiveRate"`
	TokenRate                string `json:"tokenEffectiveRate"`
	DestinationChainGasPrice string `json:"dstGasPrice"`
	Signature                string `json:"signature"`
	FromDomainID             int    `json:"fromDomainID"`
	ToDomainID               int    `json:"toDomainID"`
	ResourceID               int    `json:"resourceID"`
	DataTimestamp            int64  `json:"dataTimestamp"`
	SignatureTimestamp       int64  `json:"signatureTimestamp"`
}
