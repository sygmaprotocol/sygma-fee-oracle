package oracle_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/ChainSafe/chainbridge-fee-oracle/oracle"
	mockOracle "github.com/ChainSafe/chainbridge-fee-oracle/oracle/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type GasPriceOracleTestSuite struct {
	suite.Suite
	appBase          *base.FeeOracleAppBase
	gasPriceOperator *oracle.GasPriceOracleOperator
	oracle           *mockOracle.MockGasPriceOracle
	testdata         *types.GasPricesResp
}

func TestRunGasPriceOracleTestSuite(t *testing.T) {
	suite.Run(t, new(GasPriceOracleTestSuite))
}

func (s *GasPriceOracleTestSuite) SetupSuite() {}

func (s *GasPriceOracleTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
}

func (s *GasPriceOracleTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.oracle = mockOracle.NewMockGasPriceOracle(gomockController)
	s.appBase = base.NewFeeOracleAppBase("../config/config.template.yaml")
	s.gasPriceOperator = oracle.NewGasPriceOracleOperator(s.appBase.GetLogger(), s.oracle)
	s.testdata = &types.GasPricesResp{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "etherscan",
		DomainId:        "ethereum",
		Time:            time.Now(),
	}
}

func (s *GasPriceOracleTestSuite) TearDownTest() {
	_ = s.appBase.GetStore().Close()
}

func (s *GasPriceOracleTestSuite) TestInquiryGasPrice_Failure() {
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainId).Return(nil, errors.New("error"))

	_, err := s.gasPriceOperator.Run(s.testdata.DomainId)

	s.NotNil(err)
}

func (s *GasPriceOracleTestSuite) TestInquiryGasPrice_Success() {
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainId).Return(s.testdata, nil)

	_, err := s.gasPriceOperator.Run(s.testdata.DomainId)

	s.Nil(err)
}
