package types

const (
	GWei = 1e9
)

// GasPricesResp is the type defined for general gas price
// all fields are in wei when referring to EVM based chain
type GasPricesResp struct {
	SafeGasPrice    string `json:"safeGasPrice" mapstructure:"safeGasPrice"`
	ProposeGasPrice string `json:"proposeGasPrice" mapstructure:"proposeGasPrice"`
	FastGasPrice    string `json:"fastGasPrice" mapstructure:"fastGasPrice"`
	OracleName      string `json:"oracleName" mapstructure:"oracleName"`
	DomainId        string `json:"domainId" mapstructure:"domainId"`
	Time            string `json:"time" mapstructure:"time"`
}

// ConversionRateResp is the type defined for general conversion rate
// Rate = how much of Foreign is for 1 Base
type ConversionRateResp struct {
	Base       string  `json:"base" mapstructure:"base"`
	Foreign    string  `json:"foreign" mapstructure:"foreign"`
	Rate       float64 `json:"rate" mapstructure:"rate"`
	OracleName string  `json:"oracleName" mapstructure:"oracleName"`
	Time       string  `json:"time" mapstructure:"time"`
}
