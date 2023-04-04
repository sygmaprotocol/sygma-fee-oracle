// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package types

const (
	GWei        = 1e9
	GWeiDecimal = 9
)

// GasPrices is the type defined for general gas price
// all fields are in wei when referring to EVM based chain
type GasPrices struct {
	SafeGasPrice    string `json:"safeGasPrice" mapstructure:"safeGasPrice"`
	ProposeGasPrice string `json:"proposeGasPrice" mapstructure:"proposeGasPrice"`
	FastGasPrice    string `json:"fastGasPrice" mapstructure:"fastGasPrice"`
	OracleSource    string `json:"oracleSource" mapstructure:"oracleSource"`
	DomainID        int    `json:"domainID" mapstructure:"domainID"`
	Time            int64  `json:"time" mapstructure:"time"`
}

// ConversionRate is the type defined for general conversion rate
// Rate = how much of Foreign is for 1 Base
type ConversionRate struct {
	Base         string  `json:"base" mapstructure:"base"`
	Foreign      string  `json:"foreign" mapstructure:"foreign"`
	Rate         float64 `json:"rate" mapstructure:"rate"`
	OracleSource string  `json:"oracleSource" mapstructure:"oracleSource"`
	Time         int64   `json:"time" mapstructure:"time"`
}

type Rate struct {
	BaseRate                 string `json:"baseEffectiveRate"`
	TokenRate                string `json:"tokenEffectiveRate"`
	DestinationChainGasPrice string `json:"dstGasPrice"`
	Signature                string `json:"signature"`
	FromDomainID             int    `json:"fromDomainID"`
	ToDomainID               int    `json:"toDomainID"`
	ResourceID               string `json:"resourceID"`
	MsgGasLimit              string `json:"msgGasLimit"`
	// DataTimestamp represents the timestamp that the data is fetched from external services
	DataTimestamp int64 `json:"dataTimestamp"`
	// SignatureTimestamp represents the timestamp that the endpoint responses with the signature
	SignatureTimestamp int64 `json:"signatureTimestamp"`
	// ExpirationTimestamp represents the timestamp that the response data expires on the fee handler contract
	ExpirationTimestamp int64 `json:"expirationTimestamp"`
	// Indicator if the app is running in debug mode
	Debug bool `json:"debug"`
}
