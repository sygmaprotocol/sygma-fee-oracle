// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/sirupsen/logrus"
)

type ConversionRateOracle interface {
	Oracle
	InquiryConversionRate(baseCurrency, foreignCurrency string) (*types.ConversionRate, error)
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

func (c *ConversionRateOracleOperator) Run(baseCurrency, foreignCurrency string) (*types.ConversionRate, error) {
	return c.oracle.InquiryConversionRate(baseCurrency, foreignCurrency)
}

func (c *ConversionRateOracleOperator) GetOracleSource() string {
	return c.oracle.Source()
}

func (c *ConversionRateOracleOperator) IsOracleEnabled() bool {
	return c.oracle.IsEnabled()
}
