// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package store_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/store"
	mockStore "github.com/ChainSafe/sygma-fee-oracle/store/mock"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type GasPriceStoreTestSuite struct {
	suite.Suite
	gasPriceStore *store.GasPriceStore
	db            *mockStore.MockStore
	testdata      *types.GasPrices
}

func TestRunGasPriceStoreTestSuite(t *testing.T) {
	suite.Run(t, new(GasPriceStoreTestSuite))
}

func (s *GasPriceStoreTestSuite) SetupSuite() {}

func (s *GasPriceStoreTestSuite) TearDownSuite() {}

func (s *GasPriceStoreTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.db = mockStore.NewMockStore(gomockController)
	s.gasPriceStore = store.NewGasPriceStore(s.db)
	s.testdata = &types.GasPrices{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "etherscan",
		DomainName:      "ethereum",
		Time:            time.Now().UnixMilli(),
	}
}

func (s *GasPriceStoreTestSuite) TearDownTest() {}

func (s *GasPriceStoreTestSuite) TestStoreGasPrice_Failure() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName)), dataBytes).Return(errors.New("error"))

	err = s.gasPriceStore.StoreGasPrice(s.testdata)

	s.NotNil(err)
}

func (s *GasPriceStoreTestSuite) TestStoreGasPrice_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName)), dataBytes).Return(nil)

	err = s.gasPriceStore.StoreGasPrice(s.testdata)

	s.Nil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_Failure() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName))).Return(nil, errors.New("error"))

	_, err := s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainName)

	s.NotNil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Get([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName))).Return(dataBytes, nil)

	_, err = s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainName)

	s.Nil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_NotFound() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName))).Return(nil, store.ErrNotFound)

	_, err := s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainName)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *GasPriceStoreTestSuite) TestGetGasPriceByDomain_NotFound() {
	var dataReceiver *types.GasPrices
	s.db.EXPECT().GetByPrefix([]byte("gasprice:"), dataReceiver).Return(nil, store.ErrNotFound)

	_, err := s.gasPriceStore.GetGasPricesByDomain(s.testdata.DomainName)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *GasPriceStoreTestSuite) TestGetGasPriceByDomain_Success() {
	re := make([]interface{}, 0)
	var dataReceiverInterface interface{}

	d1 := &types.GasPrices{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "",
		DomainName:      "ethereum",
		Time:            time.Time{}.UnixMilli(),
	}
	jd1, err := json.Marshal(d1)
	s.Nil(err)
	err = json.Unmarshal(jd1, &dataReceiverInterface)
	s.Nil(err)
	re = append(re, dataReceiverInterface)

	d2 := &types.GasPrices{
		SafeGasPrice:    "2",
		ProposeGasPrice: "3",
		FastGasPrice:    "4",
		OracleName:      "",
		DomainName:      "ethereum",
		Time:            time.Time{}.UnixMilli(),
	}

	jd2, err := json.Marshal(d2)
	s.Nil(err)
	err = json.Unmarshal(jd2, &dataReceiverInterface)
	s.Nil(err)
	re = append(re, dataReceiverInterface)

	var dataReceiver *types.GasPrices
	s.db.EXPECT().GetByPrefix([]byte("gasprice:"), dataReceiver).Return(re, nil)

	fetchedData, err := s.gasPriceStore.GetGasPricesByDomain(s.testdata.DomainName)
	s.Nil(err)

	s.Equal(2, len(fetchedData))
	s.EqualValues(*d1, fetchedData[0])
	s.EqualValues(*d2, fetchedData[1])
}
