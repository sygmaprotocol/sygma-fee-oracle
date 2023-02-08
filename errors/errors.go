// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package errors

import (
	"errors"
	"fmt"
)

const (
	loadConfigErrCode          = 10000
	invalidRequestInputErrCode = 10001
	parseRequestErrCode        = 10002
)

var (
	LoadConfig          = FeeOracleError{loadConfigErrCode, "failed to load app config", ""}
	ParseRequest        = FeeOracleError{parseRequestErrCode, "failed to parse request data", ""}
	InvalidRequestInput = FeeOracleError{invalidRequestInputErrCode, "invalid request input", ""}
	InternalServerError = FeeOracleError{}
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
