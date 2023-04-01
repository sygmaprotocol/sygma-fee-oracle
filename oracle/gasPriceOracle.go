// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ChainSafe/sygma-fee-oracle/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
)

type GasPriceOracle interface {
	Oracle
	InquiryGasPrice(domainID int) (*types.GasPrices, error)
	SupportedGasPriceDomainIds() []int
}

type GasPriceConverter func(gasPrices *types.GasPrices) (*types.GasPrices, error)

type GasPriceOracleOperator struct {
	log *logrus.Entry

	oracle GasPriceOracle

	gasPriceConverterFns map[string]GasPriceConverter
}

func NewGasPriceOracleOperator(log *logrus.Entry, oracle GasPriceOracle) *GasPriceOracleOperator {
	return &GasPriceOracleOperator{
		log:    log.WithField("oracle", "gas price"),
		oracle: oracle,
		gasPriceConverterFns: map[string]GasPriceConverter{
			"etherscan":   etherscanGasPriceConverter,
			"polygonscan": polygonscanGasPriceConverter,
			"moonscan":    moonscanGasPriceConverter,
		},
	}
}

func (g *GasPriceOracleOperator) Run(domainID int) (*types.GasPrices, error) {
	gasPriceData, err := g.oracle.InquiryGasPrice(domainID)
	if err != nil {
		return nil, err
	}
	return g.GasPriceUnitConverter(gasPriceData)
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

func (g *GasPriceOracleOperator) GasPriceUnitConverter(gasPrices *types.GasPrices) (*types.GasPrices, error) {
	return g.gasPriceConverterFns[g.oracle.Name()](gasPrices)
}

func etherscanGasPriceConverter(gasPrices *types.GasPrices) (*types.GasPrices, error) {
	safeGasPriceValue, err := util.Str2BigInt(gasPrices.SafeGasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert gas price response value")
	}
	proposeGasPriceValue, err := util.Str2BigInt(gasPrices.ProposeGasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert gas price response value")
	}
	fastGasPriceValue, err := util.Str2BigInt(gasPrices.FastGasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert gas price response value")
	}

	return &types.GasPrices{
		SafeGasPrice:    new(big.Int).Mul(safeGasPriceValue, big.NewInt(types.GWei)).String(),
		ProposeGasPrice: new(big.Int).Mul(proposeGasPriceValue, big.NewInt(types.GWei)).String(),
		FastGasPrice:    new(big.Int).Mul(fastGasPriceValue, big.NewInt(types.GWei)).String(),
		OracleName:      gasPrices.OracleName,
		DomainID:        gasPrices.DomainID,
		Time:            gasPrices.Time,
	}, nil
}

func polygonscanGasPriceConverter(gasPrices *types.GasPrices) (*types.GasPrices, error) {
	safeGasPriceValue, err := util.Large2SmallUnitConverter(gasPrices.SafeGasPrice, types.GWeiDecimal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert safe gasprice")
	}
	proposeGasPriceValue, err := util.Large2SmallUnitConverter(gasPrices.ProposeGasPrice, types.GWeiDecimal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert propose gasprice")
	}
	fastGasPriceValue, err := util.Large2SmallUnitConverter(gasPrices.FastGasPrice, types.GWeiDecimal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert fast gasprice")
	}

	return &types.GasPrices{
		SafeGasPrice:    safeGasPriceValue.String(),
		ProposeGasPrice: proposeGasPriceValue.String(),
		FastGasPrice:    fastGasPriceValue.String(),
		OracleName:      gasPrices.OracleName,
		DomainID:        gasPrices.DomainID,
		Time:            gasPrices.Time,
	}, nil
}

func moonscanGasPriceConverter(gasPrices *types.GasPrices) (*types.GasPrices, error) {
	return gasPrices, nil
}
