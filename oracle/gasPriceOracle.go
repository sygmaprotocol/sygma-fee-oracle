// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/sirupsen/logrus"
)

type GasPriceOracle interface {
	Oracle
	InquiryGasPrice() (*types.GasPrices, error)
}

type GasPriceConverter func(gasPrices *types.GasPrices) (*types.GasPrices, error)

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

func (g *GasPriceOracleOperator) Run() (*types.GasPrices, error) {
	return g.oracle.InquiryGasPrice()
}

func (g *GasPriceOracleOperator) GetOracleSource() string {
	return g.oracle.Source()
}

func (g *GasPriceOracleOperator) IsOracleEnabled() bool {
	return g.oracle.IsEnabled()
}

func (g *GasPriceOracleOperator) GetOracle() GasPriceOracle {
	return g.oracle
}
