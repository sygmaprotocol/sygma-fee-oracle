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
		Time:            time.Now().String(),
	}
}

func (s *GasPriceStoreTestSuite) TearDownTest() {}

func (s *GasPriceStoreTestSuite) TestStoreGasPrice_Failure() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainId)), dataBytes).Return(errors.New("error"))

	err = s.gasPriceStore.StoreGasPrice(s.testdata)

	s.NotNil(err)
}

func (s *GasPriceStoreTestSuite) TestStoreGasPrice_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainId)), dataBytes).Return(nil)

	err = s.gasPriceStore.StoreGasPrice(s.testdata)

	s.Nil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_Failure() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainId))).Return(nil, errors.New("error"))

	_, err := s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainId)

	s.NotNil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_Success() {
	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Get([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainId))).Return(dataBytes, nil)

	_, err = s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainId)

	s.Nil(err)
}

func (s *GasPriceStoreTestSuite) TestGetGasPrice_NotFound() {
	s.db.EXPECT().Get([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainId))).Return(nil, store.ErrNotFound)

	_, err := s.gasPriceStore.GetGasPrice(s.testdata.OracleName, s.testdata.DomainId)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *GasPriceStoreTestSuite) TestGetGasPriceByDomain_NotFound() {
	var dataReceiver *types.GasPricesResp
	s.db.EXPECT().GetByPrefix([]byte(fmt.Sprintf("gasprice:")), dataReceiver).Return(nil, store.ErrNotFound)

	_, err := s.gasPriceStore.GetGasPriceByDomain(s.testdata.DomainId)

	s.EqualError(err, store.ErrNotFound.Error())
}

func (s *GasPriceStoreTestSuite) TestGetGasPriceByDomain_Success() {
	re := make([]interface{}, 0)
	var dataReceiverInterface interface{}

	d1 := &types.GasPricesResp{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "",
		DomainId:        "ethereum",
		Time:            time.Time{}.String(),
	}
	jd1, err := json.Marshal(d1)
	s.Nil(err)
	err = json.Unmarshal(jd1, &dataReceiverInterface)
	s.Nil(err)
	re = append(re, dataReceiverInterface)

	d2 := &types.GasPricesResp{
		SafeGasPrice:    "2",
		ProposeGasPrice: "3",
		FastGasPrice:    "4",
		OracleName:      "",
		DomainId:        "ethereum",
		Time:            time.Time{}.String(),
	}

	jd2, err := json.Marshal(d2)
	s.Nil(err)
	err = json.Unmarshal(jd2, &dataReceiverInterface)
	s.Nil(err)
	re = append(re, dataReceiverInterface)

	fmt.Println(re)

	var dataReceiver *types.GasPricesResp
	s.db.EXPECT().GetByPrefix([]byte(fmt.Sprintf("gasprice:")), dataReceiver).Return(re, nil)

	fetchedData, err := s.gasPriceStore.GetGasPriceByDomain(s.testdata.DomainId)
	s.Nil(err)

	s.Equal(2, len(fetchedData))
	s.EqualValues(*d1, fetchedData[0])
	s.EqualValues(*d2, fetchedData[1])
}
