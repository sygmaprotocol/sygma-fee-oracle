package oracle

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/sirupsen/logrus"
)

type ConversionRateOracle interface {
	Oracle
	InquiryConversionRate(baseCurrency, foreignCurrency string) (*types.ConversionRateResp, error)
}

type ConversionRateOracleOperator struct {
	log *logrus.Entry

	oracle ConversionRateOracle
}

func NewConversionRateOracleOperator(log *logrus.Entry, oracle ConversionRateOracle) *ConversionRateOracleOperator {
	return &ConversionRateOracleOperator{
		log:    log.WithField("oracle", "conversion rate"),
		oracle: oracle,
	}
}

func (c *ConversionRateOracleOperator) Run(baseCurrency, foreignCurrency string) (*types.ConversionRateResp, error) {
	return c.oracle.InquiryConversionRate(baseCurrency, foreignCurrency)
}

func (c *ConversionRateOracleOperator) GetOracleName() string {
	return c.oracle.Name()
}

func (c *ConversionRateOracleOperator) IsOracleEnabled() bool {
	return c.oracle.IsEnabled()
}
