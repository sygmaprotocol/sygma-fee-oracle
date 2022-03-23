package oracle_test

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/ChainSafe/chainbridge-fee-oracle/oracle"
	mockOracle "github.com/ChainSafe/chainbridge-fee-oracle/oracle/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type ConversionRateOracleTestSuite struct {
	suite.Suite
	appBase                *base.FeeOracleAppBase
	conversionRateOperator *oracle.ConversionRateOracleOperator
	oracle                 *mockOracle.MockConversionRateOracle
	testdata               *types.ConversionRateResp
}

func TestRunConversionRateOracleTestSuite(t *testing.T) {
	suite.Run(t, new(ConversionRateOracleTestSuite))
}

func (s *ConversionRateOracleTestSuite) SetupSuite() {}

func (s *ConversionRateOracleTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
}

func (s *ConversionRateOracleTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.oracle = mockOracle.NewMockConversionRateOracle(gomockController)
	s.appBase = base.NewFeeOracleAppBase("../config/config.template.yaml")
	s.conversionRateOperator = oracle.NewConversionRateOracleOperator(s.appBase.GetLogger(), s.oracle)
	s.testdata = &types.ConversionRateResp{
		Base:       "eth",
		Foreign:    "usdt",
		Rate:       3000,
		OracleName: "cooinmarketcap",
		Time:       time.Time{},
	}
}

func (s *ConversionRateOracleTestSuite) TearDownTest() {
	_ = s.appBase.GetStore().Close()
}

func (s *ConversionRateOracleTestSuite) TestInquiryConversionRate_Failure() {
	s.oracle.EXPECT().InquiryConversionRate(s.testdata.Base, s.testdata.Foreign).Return(nil, errors.New("error"))

	_, err := s.conversionRateOperator.Run(s.testdata.Base, s.testdata.Foreign)

	s.NotNil(err)
}

func (s *ConversionRateOracleTestSuite) TestInquiryConversionRate_Success() {
	s.oracle.EXPECT().InquiryConversionRate(s.testdata.Base, s.testdata.Foreign).Return(s.testdata, nil)

	_, err := s.conversionRateOperator.Run(s.testdata.Base, s.testdata.Foreign)

	s.Nil(err)
}
