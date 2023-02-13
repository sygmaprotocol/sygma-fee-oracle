// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"errors"

	oracleErrors "github.com/ChainSafe/sygma-fee-oracle/errors"
	"github.com/gin-gonic/gin"
)

type ReturnErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type ReturnSuccessResponse interface{}

func newReturnErrorResp(fe *oracleErrors.FeeOracleError, err error) *ReturnErrorResponse {
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
