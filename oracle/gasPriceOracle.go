// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/sirupsen/logrus"
)

type GasPriceOracle interface {
	Oracle
	InquiryGasPrice(domainID int) (*types.GasPrices, error)
	SupportedGasPriceDomainIds() []int
}

type GasPriceOracleOperator struct {
	log *logrus.Entry

	oracle GasPriceOracle
}

func NewGasPriceOracleOperator(log *logrus.Entry, oracle GasPriceOracle) *GasPriceOracleOperator {
	return &GasPriceOracleOperator{
		log:    log.WithField("oracle", "gas price"),
		oracle: oracle,
	}
}

func (g *GasPriceOracleOperator) Run(domainID int) (*types.GasPrices, error) {
	return g.oracle.InquiryGasPrice(domainID)
}

func (g *GasPriceOracleOperator) GetOracleName() string {
	return g.oracle.Name()
}

func (g *GasPriceOracleOperator) IsOracleEnabled() bool {
	return g.oracle.IsEnabled()
}

func (g *GasPriceOracleOperator) GetOracle() GasPriceOracle {
	return g.oracle
}
