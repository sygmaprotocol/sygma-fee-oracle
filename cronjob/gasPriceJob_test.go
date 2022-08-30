// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package cronjob_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ChainSafe/sygma-fee-oracle/base"
	"github.com/ChainSafe/sygma-fee-oracle/cronjob"
	"github.com/ChainSafe/sygma-fee-oracle/oracle"
	mockOracle "github.com/ChainSafe/sygma-fee-oracle/oracle/mock"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	mockStore "github.com/ChainSafe/sygma-fee-oracle/store/mock"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type GasPriceJobTestSuite struct {
	suite.Suite
	appBase          *base.FeeOracleAppBase
	gasPriceOperator *oracle.GasPriceOracleOperator
	oracle           *mockOracle.MockGasPriceOracle
	gasPriceStore    *store.GasPriceStore
	db               *mockStore.MockStore
	job              *cronjob.Job
	testdata         *types.GasPrices
}

func TestRunGasPriceJobTestSuite(t *testing.T) {
	suite.Run(t, new(GasPriceJobTestSuite))
}

func (s *GasPriceJobTestSuite) SetupSuite() {
	priv, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	crypto.FromECDSA(priv)
	err = os.WriteFile("./keyfile.priv", crypto.FromECDSA(priv), 0666)
	if err != nil {
		panic(err)
	}
}

func (s *GasPriceJobTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
	_ = os.RemoveAll("./keyfile.priv")
}

func (s *GasPriceJobTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.oracle = mockOracle.NewMockGasPriceOracle(gomockController)
	s.appBase = base.NewFeeOracleAppBase("../config/config.template.yaml", "../domain.json", "../resource.json", "./keyfile.priv", "secp256k1")
	s.gasPriceOperator = oracle.NewGasPriceOracleOperator(s.appBase.GetLogger(), s.oracle)
	s.db = mockStore.NewMockStore(gomockController)
	s.gasPriceStore = store.NewGasPriceStore(s.db)
	s.testdata = &types.GasPrices{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "test oracle",
		DomainName:      "ethereum",
		Time:            time.Now().UnixMilli(),
	}

	gasPriceOracle := oracle.NewGasPriceOracleOperator(s.appBase.GetLogger(), s.oracle)
	gasPriceOracles := make(map[string]*oracle.GasPriceOracleOperator)
	gasPriceOracles["test oracle"] = gasPriceOracle

	gasPriceStore := store.NewGasPriceStore(s.db)

	cronJobServer := cronjob.NewCronJobs(s.appBase, nil, gasPriceOracles, nil, gasPriceStore, s.appBase.GetLogger())

	s.job = cronjob.NewJob(cronJobServer,
		s.appBase.GetConfig().CronJobConfig().UpdateGasPriceJob.Name,
		s.appBase.GetConfig().CronJobConfig().UpdateGasPriceJob.CheckFrequency,
		s.appBase.GetConfig().CronJobConfig().UpdateGasPriceJob.Enable)
	s.job.SetOperationFunc(cronjob.GasPriceJobOperation(s.job))
}

func (s *GasPriceJobTestSuite) TearDownTest() {
	_ = s.appBase.GetStore().Close()
}

func (s *GasPriceJobTestSuite) TestJobOperation_Oracle_Disabled() {
	s.oracle.EXPECT().IsEnabled().Return(false)
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainName).Return(nil, errors.New("error")).Times(0)
	s.oracle.EXPECT().Name().Times(0)

	cronjob.GasPriceJobOperation(s.job)()
}

func (s *GasPriceJobTestSuite) TestJobOperation_Run_Failure() {
	s.oracle.EXPECT().IsEnabled().Return(true)
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainName).Return(nil, errors.New("error")).Times(1)
	s.oracle.EXPECT().Name().Return("test oracle").Times(1)
	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName)), []byte("")).Return(nil).Times(0)

	cronjob.GasPriceJobOperation(s.job)()
}

func (s *GasPriceJobTestSuite) TestJobOperation_Run_Success_StoreGasPrice_Failure() {
	s.oracle.EXPECT().IsEnabled().Return(true)
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainName).Return(s.testdata, nil).Times(1)
	s.oracle.EXPECT().Name().Return("test oracle").Times(0)

	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName)), dataBytes).Return(errors.New("error")).Times(1)
	cronjob.GasPriceJobOperation(s.job)()
}

func (s *GasPriceJobTestSuite) TestJobOperation_Run_Success_StoreGasPrice_Success() {
	s.oracle.EXPECT().IsEnabled().Return(true)
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainName).Return(s.testdata, nil).Times(1)
	s.oracle.EXPECT().Name().Return("test oracle").Times(0)

	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("gasprice:%s:%s", s.testdata.OracleName, s.testdata.DomainName)), dataBytes).Return(nil).Times(1)
	cronjob.GasPriceJobOperation(s.job)()
}
