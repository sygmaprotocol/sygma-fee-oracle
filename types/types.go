// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package types

const (
	GWei        = 1e9
	GWeiDecimal = 9
	Wei         = 1e18
)

// GasPrices is the type defined for general gas price
// all fields are in wei when referring to EVM based chain
type GasPrices struct {
	SafeGasPrice    string `json:"safeGasPrice" mapstructure:"safeGasPrice"`
	ProposeGasPrice string `json:"proposeGasPrice" mapstructure:"proposeGasPrice"`
	FastGasPrice    string `json:"fastGasPrice" mapstructure:"fastGasPrice"`
	OracleName      string `json:"oracleName" mapstructure:"oracleName"`
	DomainName      string `json:"domainName" mapstructure:"domainName"`
	Time            int64  `json:"time" mapstructure:"time"`
}

// ConversionRate is the type defined for general conversion rate
// Rate = how much of Foreign is for 1 Base
type ConversionRate struct {
	Base       string  `json:"base" mapstructure:"base"`
	Foreign    string  `json:"foreign" mapstructure:"foreign"`
	Rate       float64 `json:"rate" mapstructure:"rate"`
	OracleName string  `json:"oracleName" mapstructure:"oracleName"`
	Time       int64   `json:"time" mapstructure:"time"`
}
