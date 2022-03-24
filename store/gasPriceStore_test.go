package store_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	mockStore "github.com/ChainSafe/chainbridge-fee-oracle/store/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type GasPriceStoreTestSuite struct {
	suite.Suite
	gasPriceStore *store.GasPriceStore
	db            *mockStore.MockStore
	testdata      *types.GasPricesResp
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
	s.testdata = &types.GasPricesResp{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "etherscan",
		DomainId:        "ethereum",
		Time:            time.Now(),
	}
}

func (s *GasPriceStoreTestSuite) TearDownTest() {}

func (s *GasPriceStoreTestSuite) TestStoreGasPrice_Failure() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("%s:%s", s.testdata.OracleName, s.testdata.DomainId)), dataBytes).Return(errors.New("error"))

	err = s.gasPriceStore.StoreGasPrice(s.testdata)

	s.NotNil(err)
}

func (s *GasPriceStoreTestSuite) TestStoreGasPrice_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("%s:%s", s.testdata.OracleName, s.testdata.DomainId)), dataBytes).Return(nil)

	err = s.gasPriceStore.StoreGasPrice(s.testdata)

	s.Nil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_Failure() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("%s:%s", s.testdata.OracleName, s.testdata.DomainId))).Return(nil, errors.New("error"))

	_, err := s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainId)

	s.NotNil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Get([]byte(fmt.Sprintf("%s:%s", s.testdata.OracleName, s.testdata.DomainId))).Return(dataBytes, nil)

	_, err = s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainId)

	s.Nil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_NotFound() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("%s:%s", s.testdata.OracleName, s.testdata.DomainId))).Return(nil, store.ErrNotFound)

	_, err := s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainId)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *GasPriceStoreTestSuite) TestGetGasPriceByDomain_NotFound() {
	s.db.EXPECT().GetBySuffix([]byte(fmt.Sprintf(":%s", s.testdata.DomainId))).Return(nil, store.ErrNotFound)

	_, err := s.gasPriceStore.GetGasPriceByDomain(s.testdata.DomainId)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *GasPriceStoreTestSuite) TestGetGasPriceByDomain_Success() {
	re := make([][]byte, 0)

	d1 := &types.GasPricesResp{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "",
		DomainId:        "ethereum",
		Time:            time.Time{},
	}
	jd1, err := json.Marshal(d1)
	s.Nil(err)
	re = append(re, jd1)

	d2 := &types.GasPricesResp{
		SafeGasPrice:    "2",
		ProposeGasPrice: "3",
		FastGasPrice:    "4",
		OracleName:      "",
		DomainId:        "ethereum",
		Time:            time.Time{},
	}

	jd2, err := json.Marshal(d2)
	s.Nil(err)
	re = append(re, jd2)

	s.db.EXPECT().GetBySuffix([]byte(fmt.Sprintf(":%s", s.testdata.DomainId))).Return(re, nil)

	fetchedData, err := s.gasPriceStore.GetGasPriceByDomain(s.testdata.DomainId)
	s.Nil(err)

	s.Equal(2, len(fetchedData))
	s.EqualValues(d1, fetchedData[0])
	s.EqualValues(d2, fetchedData[1])
}
