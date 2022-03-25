package store_test

import (
	"encoding/json"
	"fmt"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	mockStore "github.com/ChainSafe/chainbridge-fee-oracle/store/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
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
	testdata            *types.ConversionRateResp
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
	s.testdata = &types.ConversionRateResp{
		Base:       "eth",
		Foreign:    "usdt",
		Rate:       3000,
		OracleName: "cooinmarketcap",
		Time:       time.Time{}.String(),
	}
}

func (s *ConversionRateStoreTestSuite) TearDownTest() {}

func (s *ConversionRateStoreTestSuite) TestStoreConversionRate_Failure() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)), dataBytes).Return(errors.New("error"))

	err = s.conversionRateStore.StoreConversionRate(s.testdata)

	s.NotNil(err)
}

func (s *ConversionRateStoreTestSuite) TestStoreConversionRate_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)), dataBytes).Return(nil)

	err = s.conversionRateStore.StoreConversionRate(s.testdata)

	s.Nil(err)
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRate_Failure() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign))).Return(nil, errors.New("error"))

	_, err := s.conversionRateStore.GetConversionRate(s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)

	s.NotNil(err)
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRate_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Get([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign))).Return(dataBytes, nil)

	_, err = s.conversionRateStore.GetConversionRate(s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)

	s.Nil(err)
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRate_NotFound() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("conversionrate:%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign))).Return(nil, store.ErrNotFound)

	_, err := s.conversionRateStore.GetConversionRate(s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRateByCurrencyPair_NotFound() {
	var dataReceiver *types.ConversionRateResp
	s.db.EXPECT().GetByPrefix([]byte(fmt.Sprintf("conversionrate:")), dataReceiver).Return(nil, store.ErrNotFound)

	_, err := s.conversionRateStore.GetConversionRateByCurrencyPair(s.testdata.Base, s.testdata.Foreign)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *ConversionRateStoreTestSuite) TestGetConversionRateByCurrencyPair_Success() {
	re := make([]interface{}, 0)
	var dataReceiverInterface interface{}

	d1 := &types.ConversionRateResp{
		Base:       "eth",
		Foreign:    "usdt",
		Rate:       2345.987,
		OracleName: "",
		Time:       time.Time{}.String(),
	}
	jd1, err := json.Marshal(d1)
	s.Nil(err)

	err = json.Unmarshal(jd1, &dataReceiverInterface)
	s.Nil(err)

	re = append(re, dataReceiverInterface)

	d2 := &types.ConversionRateResp{
		Base:       "eth",
		Foreign:    "usdt",
		Rate:       2350.234,
		OracleName: "",
		Time:       time.Time{}.String(),
	}

	jd2, err := json.Marshal(d2)
	s.Nil(err)
	err = json.Unmarshal(jd2, &dataReceiverInterface)
	s.Nil(err)
	re = append(re, dataReceiverInterface)

	var dataReceiver *types.ConversionRateResp
	s.db.EXPECT().GetByPrefix([]byte(fmt.Sprintf("conversionrate:")), dataReceiver).Return(re, nil)

	fetchedData, err := s.conversionRateStore.GetConversionRateByCurrencyPair(s.testdata.Base, s.testdata.Foreign)
	s.Nil(err)

	s.Equal(2, len(fetchedData))
	s.EqualValues(*d1, fetchedData[0])
	s.EqualValues(*d2, fetchedData[1])
}
