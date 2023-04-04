// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package store_test

import (
	"encoding/json"
	"fmt"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	mockStore "github.com/ChainSafe/sygma-fee-oracle/store/mock"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ConversionRateStoreTestSuite struct {
	suite.Suite
	conversionRateStore *store.ConversionRateStore
	db                  *mockStore.MockStore
	testdata            *types.ConversionRate
}

func TestRunConversionRateStoreTestSuite(t *testing.T) {
	suite.Run(t, new(ConversionRateStoreTestSuite))
}

func (s *ConversionRateStoreTestSuite) SetupSuite() {}

func (s *ConversionRateStoreTestSuite) TearDownSuite() {}

func (s *ConversionRateStoreTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.db = mockStore.NewMockStore(gomockController)
	s.conversionRateStore = store.NewConversionRateStore(s.db)
	s.testdata = &types.ConversionRate{
		Base:         "eth",
		Foreign:      "usdt",
		Rate:         3000,
		OracleSource: "cooinmarketcap",
		Time:         time.Time{}.UnixMilli(),
	}
}

func (s *ConversionRateStoreTestSuite) TearDownTest() {}

func (s *ConversionRateStoreTestSuite) TestStoreConversionRate_Failure() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign)), dataBytes).Return(errors.New("error"))

	err = s.conversionRateStore.StoreConversionRate(s.testdata)

	s.NotNil(err)
}

func (s *ConversionRateStoreTestSuite) TestStoreConversionRate_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign)), dataBytes).Return(nil)

	err = s.conversionRateStore.StoreConversionRate(s.testdata)

	s.Nil(err)
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRate_Failure() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign))).Return(nil, errors.New("error"))

	_, err := s.conversionRateStore.GetConversionRate(s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign)

	s.NotNil(err)
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRate_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Get([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign))).Return(dataBytes, nil)

	_, err = s.conversionRateStore.GetConversionRate(s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign)

	s.Nil(err)
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRate_NotFound() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign))).Return(nil, store.ErrNotFound)

	_, err := s.conversionRateStore.GetConversionRate(s.testdata.OracleSource, s.testdata.Base, s.testdata.Foreign)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRateByCurrencyPair_NotFound() {
	var dataReceiver *types.ConversionRate
	s.db.EXPECT().GetByPrefix([]byte("conversionrate:"), dataReceiver).Return(nil, store.ErrNotFound)

	_, err := s.conversionRateStore.GetConversionRatesByCurrencyPair(s.testdata.Base, s.testdata.Foreign)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRateByCurrencyPair_Success() {
	re := make([]interface{}, 0)
	var dataReceiverInterface interface{}

	d1 := &types.ConversionRate{
		Base:         "eth",
		Foreign:      "usdt",
		Rate:         2345.987,
		OracleSource: "",
		Time:         time.Time{}.UnixMilli(),
	}
	jd1, err := json.Marshal(d1)
	s.Nil(err)

	err = json.Unmarshal(jd1, &dataReceiverInterface)
	s.Nil(err)

	re = append(re, dataReceiverInterface)

	d2 := &types.ConversionRate{
		Base:         "eth",
		Foreign:      "usdt",
		Rate:         2350.234,
		OracleSource: "",
		Time:         time.Time{}.UnixMilli(),
	}

	jd2, err := json.Marshal(d2)
	s.Nil(err)
	err = json.Unmarshal(jd2, &dataReceiverInterface)
	s.Nil(err)
	re = append(re, dataReceiverInterface)

	var dataReceiver *types.ConversionRate
	s.db.EXPECT().GetByPrefix([]byte("conversionrate:"), dataReceiver).Return(re, nil)

	fetchedData, err := s.conversionRateStore.GetConversionRatesByCurrencyPair(s.testdata.Base, s.testdata.Foreign)
	s.Nil(err)

	s.Equal(2, len(fetchedData))
	s.EqualValues(*d1, fetchedData[0])
	s.EqualValues(*d2, fetchedData[1])
}
