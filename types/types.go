package types

import "time"

// GasPricesResp is the type defined for general gas price
// all fields are in gwei when referring to EVM based chain
type GasPricesResp struct {
	SafeGasPrice    string    `json:"safeGasPrice"`
	ProposeGasPrice string    `json:"proposeGasPrice"`
	FastGasPrice    string    `json:"fastGasPrice"`
	OracleName      string    `json:"oracleName"`
	DomainId        string    `json:"domainId"`
	Time            time.Time `json:"time"`
}

// ConversionRateResp is the type defined for general conversion rate
// Rate = how much of Foreign is for 1 Base
type ConversionRateResp struct {
	Base       string    `json:"base"`
	Foreign    string    `json:"foreign"`
	Rate       float64   `json:"rate"`
	OracleName string    `json:"oracleName"`
	Time       time.Time `json:"time"`
}
