// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"errors"
	"math/big"

	"github.com/ChainSafe/sygma-fee-oracle/config"
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
	BaseRate                 string   `json:"baseEffectiveRate"`
	TokenRate                string   `json:"tokenEffectiveRate"`
	DestinationChainGasPrice string   `json:"dstGasPrice"`
	Signature                string   `json:"signature"`
	FromDomainID             int      `json:"fromDomainID"`
	ToDomainID               int      `json:"toDomainID"`
	ResourceID               string   `json:"resourceID"`
	MsgGasLimit              *big.Int `json:"msgGasLimit"`
	// DataTimestamp represents the timestamp that the data is fetched from external services
	DataTimestamp int64 `json:"dataTimestamp"`
	// SignatureTimestamp represents the timestamp that the endpoint responses with the signature
	SignatureTimestamp int64 `json:"signatureTimestamp"`
	// ExpirationTimestamp represents the timestamp that the response data expires on the fee handler contract
	ExpirationTimestamp int64 `json:"expirationTimestamp"`
	// Indicator if the app is running in debug mode
	Debug bool `json:"debug"`
}
