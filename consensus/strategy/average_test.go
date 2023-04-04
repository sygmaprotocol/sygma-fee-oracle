// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package strategy_test

import (
	"errors"
	"testing"

	"github.com/ChainSafe/sygma-fee-oracle/consensus/strategy"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	mockStore "github.com/ChainSafe/sygma-fee-oracle/store/mock"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type AverageStrategyTestSuite struct {
	suite.Suite
	strategy            strategy.Strategy
	conversionRateStore *store.ConversionRateStore
	gasPriceStore       *store.GasPriceStore
	db                  *mockStore.MockStore
}

func TestRunAverageStrategyTestSuite(t *testing.T) {
	suite.Run(t, new(AverageStrategyTestSuite))
}

func (a *AverageStrategyTestSuite) SetupSuite() {
}

func (a *AverageStrategyTestSuite) TearDownSuite() {
}

func (a *AverageStrategyTestSuite) SetupTest() {
	a.strategy = &strategy.Average{}
	gomockController := gomock.NewController(a.T())
	a.db = mockStore.NewMockStore(gomockController)
	a.conversionRateStore = store.NewConversionRateStore(a.db)
	a.gasPriceStore = store.NewGasPriceStore(a.db)
}

func (a *AverageStrategyTestSuite) TearDownTest() {
}

func (a *AverageStrategyTestSuite) TestConversionRate_DB_Error() {
	var dataReceiver *types.ConversionRate
	a.db.EXPECT().GetByPrefix([]byte("conversionrate:"), dataReceiver).Return(nil, errors.New("db error"))

	re, err := a.strategy.ConversionRate(a.conversionRateStore, "eth", "matic")
	a.Nil(re)
	a.EqualError(err, "db error")
}

func (a *AverageStrategyTestSuite) TestConversionRate_Without_Data() {
	var dataReceiver *types.ConversionRate
	var emptyResult []interface{}
	a.db.EXPECT().GetByPrefix([]byte("conversionrate:"), dataReceiver).Return(emptyResult, nil)

	re, err := a.strategy.ConversionRate(a.conversionRateStore, "eth", "matic")
	a.Nil(re)
	a.EqualError(err, "no conversion rate data found")
}

func (a *AverageStrategyTestSuite) TestConversionRate_With_Data() {
	var dataReceiver *types.ConversionRate
	result := make([]interface{}, 0)
	result = append(result, types.ConversionRate{
		Base:         "eth",
		Foreign:      "matic",
		Rate:         1000.0,
		OracleSource: "",
		Time:         0,
	})
	result = append(result, types.ConversionRate{
		Base:         "eth",
		Foreign:      "matic",
		Rate:         2000.0,
		OracleSource: "",
		Time:         0,
	})
	result = append(result, types.ConversionRate{
		Base:         "eth",
		Foreign:      "matic",
		Rate:         2100.0,
		OracleSource: "",
		Time:         0,
	})
	a.db.EXPECT().GetByPrefix([]byte("conversionrate:"), dataReceiver).Return(result, nil)

	re, err := a.strategy.ConversionRate(a.conversionRateStore, "eth", "matic")
	a.Nil(err)
	a.EqualValues(1700.0, re.Rate)
}

func (a *AverageStrategyTestSuite) TestGasPrice_DB_Error() {
	var dataReceiver *types.GasPrices
	a.db.EXPECT().GetByPrefix([]byte("gasprice:"), dataReceiver).Return(nil, errors.New("db error"))

	re, err := a.strategy.GasPrice(a.gasPriceStore, 1)
	a.Nil(re)
	a.EqualError(err, "db error")
}

func (a *AverageStrategyTestSuite) TestGasPrice_Without_Data() {
	var dataReceiver *types.GasPrices
	var emptyResult []interface{}
	a.db.EXPECT().GetByPrefix([]byte("gasprice:"), dataReceiver).Return(emptyResult, nil)

	re, err := a.strategy.GasPrice(a.gasPriceStore, 1)
	a.Nil(re)
	a.EqualError(err, "no gas price data found")
}

func (a *AverageStrategyTestSuite) TestGasPrice_With_Data() {
	var dataReceiver *types.GasPrices
	result := make([]interface{}, 0)
	result = append(result, types.GasPrices{
		SafeGasPrice:    "10",
		ProposeGasPrice: "15",
		FastGasPrice:    "20",
		OracleSource:    "",
		DomainID:        1,
		Time:            0,
	})
	result = append(result, types.GasPrices{
		SafeGasPrice:    "20",
		ProposeGasPrice: "25",
		FastGasPrice:    "30",
		OracleSource:    "",
		DomainID:        1,
		Time:            0,
	})
	result = append(result, types.GasPrices{
		SafeGasPrice:    "30",
		ProposeGasPrice: "35",
		FastGasPrice:    "40",
		OracleSource:    "",
		DomainID:        1,
		Time:            0,
	})
	a.db.EXPECT().GetByPrefix([]byte("gasprice:"), dataReceiver).Return(result, nil)

	re, err := a.strategy.GasPrice(a.gasPriceStore, 1)

	a.Nil(err)
	a.EqualValues("20", re.SafeGasPrice)
	a.EqualValues("25", re.ProposeGasPrice)
	a.EqualValues("30", re.FastGasPrice)
}
