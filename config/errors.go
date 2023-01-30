// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"errors"
	"fmt"
)

const (
	loadConfigErrCode          = 10000
	invalidRequestInputErrCode = 10001
	parseRequestErrCode        = 10002
	loadDomainConfigErrCode    = 10003
	loadResourceConfigErrCode  = 10004
)

var (
	ErrLoadConfig          = FeeOracleError{loadConfigErrCode, "failed to load app config", ""}
	ErrLoadDomainConfig    = FeeOracleError{loadDomainConfigErrCode, "failed to load domain config", ""}
	ErrLoadResourceConfig  = FeeOracleError{loadResourceConfigErrCode, "failed to load resource config", ""}
	ErrParseRequest        = FeeOracleError{parseRequestErrCode, "failed to parse request data", ""}
	ErrInvalidRequestInput = FeeOracleError{invalidRequestInputErrCode, "invalid request input", ""}
	ErrInternalServerError = FeeOracleError{}
)

type FeeOracleError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// Error returns defined error code with error message
func (f *FeeOracleError) Error() string {
	return fmt.Sprintf("(%d) %s", f.Code, f.Message)
}

// ErrorMsg returns defined error message only
func (f *FeeOracleError) ErrorMsg() string {
	return f.Message
}

// Wrap returns defined error code with error message plus original error message
func (f *FeeOracleError) Wrap(err error) *FeeOracleError {
	if err == nil {
		err = errors.New("")
	}
	return &FeeOracleError{
		f.Code,
		f.Message + ": " + err.Error(),
		f.Detail}
}
