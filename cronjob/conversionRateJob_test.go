package cronjob_test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/ChainSafe/chainbridge-fee-oracle/cronjob"
	"github.com/ChainSafe/chainbridge-fee-oracle/oracle"
	mockOracle "github.com/ChainSafe/chainbridge-fee-oracle/oracle/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	mockStore "github.com/ChainSafe/chainbridge-fee-oracle/store/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type ConversionRateJobTestSuite struct {
	suite.Suite
	appBase                *base.FeeOracleAppBase
	conversionRateOperator *oracle.ConversionRateOracleOperator
	oracle                 *mockOracle.MockConversionRateOracle
	conversionRateStore    *store.ConversionRateStore
	db                     *mockStore.MockStore
	job                    *cronjob.Job
	testdata               *types.ConversionRateResp
}

func TestRunConversionRateJobTestSuite(t *testing.T) {
	suite.Run(t, new(ConversionRateJobTestSuite))
}

func (s *ConversionRateJobTestSuite) SetupSuite() {
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

func (s *ConversionRateJobTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
	_ = os.RemoveAll("./keyfile.priv")
}

func (s *ConversionRateJobTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.oracle = mockOracle.NewMockConversionRateOracle(gomockController)
	s.appBase = base.NewFeeOracleAppBase("../config/config.template.yaml", "./keyfile.priv", "secp256k1")
	s.conversionRateOperator = oracle.NewConversionRateOracleOperator(s.appBase.GetLogger(), s.oracle)
	s.db = mockStore.NewMockStore(gomockController)
	s.conversionRateStore = store.NewConversionRateStore(s.db)
	s.testdata = &types.ConversionRateResp{
		Base:       "eth",
		Foreign:    "usd",
		Rate:       3000,
		OracleName: "cooinmarketcap",
		Time:       time.Time{},
	}

	conversionRateOracle := oracle.NewConversionRateOracleOperator(s.appBase.GetLogger(), s.oracle)
	conversionRateOracles := make(map[string]*oracle.ConversionRateOracleOperator)
	conversionRateOracles["test oracle"] = conversionRateOracle

	conversionRateStore := store.NewConversionRateStore(s.db)

	cronJobServer := cronjob.NewCronJobs(s.appBase, conversionRateOracles, nil, conversionRateStore, nil, s.appBase.GetLogger())

	s.job = cronjob.NewJob(cronJobServer,
		s.appBase.GetConfig().CronJobConfig().UpdateConversionRateJob.Name,
		s.appBase.GetConfig().CronJobConfig().UpdateConversionRateJob.CheckFrequency,
		s.appBase.GetConfig().CronJobConfig().UpdateConversionRateJob.Enable)
	s.job.SetOperationFunc(cronjob.ConversionRateJobOperation(s.job))
}

func (s *ConversionRateJobTestSuite) TearDownTest() {
	_ = s.appBase.GetStore().Close()
}

func (s *ConversionRateJobTestSuite) TestJobOperation_Oracle_Disabled() {
	s.oracle.EXPECT().IsEnabled().Return(false)
	s.oracle.EXPECT().InquiryConversionRate(s.testdata.Base, s.testdata.Foreign).Return(nil, errors.New("error")).Times(0)
	s.oracle.EXPECT().Name().Times(0)

	cronjob.ConversionRateJobOperation(s.job)()
}

func (s *ConversionRateJobTestSuite) TestJobOperation_Run_Failure() {
	s.oracle.EXPECT().IsEnabled().Return(true)
	s.oracle.EXPECT().InquiryConversionRate(s.testdata.Base, s.testdata.Foreign).Return(nil, errors.New("error")).Times(1)
	s.oracle.EXPECT().Name().Return("test oracle").Times(1)
	s.db.EXPECT().Set([]byte(fmt.Sprintf("%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)), []byte("")).Return(nil).Times(0)

	cronjob.ConversionRateJobOperation(s.job)()
}

func (s *ConversionRateJobTestSuite) TestJobOperation_Run_Success_StoreGasPrice_Failure() {
	s.oracle.EXPECT().IsEnabled().Return(true)
	s.oracle.EXPECT().InquiryConversionRate(s.testdata.Base, s.testdata.Foreign).Return(s.testdata, nil).Times(1)
	s.oracle.EXPECT().Name().Return("test oracle").Times(0)

	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)), dataBytes).Return(errors.New("error")).Times(1)
	cronjob.ConversionRateJobOperation(s.job)()
}

func (s *ConversionRateJobTestSuite) TestJobOperation_Run_Success_StoreGasPrice_Success() {
	s.oracle.EXPECT().IsEnabled().Return(true)
	s.oracle.EXPECT().InquiryConversionRate(s.testdata.Base, s.testdata.Foreign).Return(s.testdata, nil).Times(1)
	s.oracle.EXPECT().Name().Return("test oracle").Times(0)

	dataBytes, err := json.Marshal(s.testdata)
	s.Nil(err)

	s.db.EXPECT().Set([]byte(fmt.Sprintf("%s:%s:%s", s.testdata.OracleName, s.testdata.Base, s.testdata.Foreign)), dataBytes).Return(nil).Times(1)
	cronjob.ConversionRateJobOperation(s.job)()
}
