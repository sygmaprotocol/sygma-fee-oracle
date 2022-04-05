package oracle_test

import (
	"os"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/ChainSafe/chainbridge-fee-oracle/oracle"
	mockOracle "github.com/ChainSafe/chainbridge-fee-oracle/oracle/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type ConversionRateOracleTestSuite struct {
	suite.Suite
	appBase                *base.FeeOracleAppBase
	conversionRateOperator *oracle.ConversionRateOracleOperator
	oracle                 *mockOracle.MockConversionRateOracle
	testdata               *types.ConversionRate
}

func TestRunConversionRateOracleTestSuite(t *testing.T) {
	suite.Run(t, new(ConversionRateOracleTestSuite))
}

func (s *ConversionRateOracleTestSuite) SetupSuite() {
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

func (s *ConversionRateOracleTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
	_ = os.RemoveAll("./keyfile.priv")
}

func (s *ConversionRateOracleTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.oracle = mockOracle.NewMockConversionRateOracle(gomockController)
	s.appBase = base.NewFeeOracleAppBase("../config/config.template.yaml", "./keyfile.priv", "secp256k1")
	s.conversionRateOperator = oracle.NewConversionRateOracleOperator(s.appBase.GetLogger(), s.oracle)
	s.testdata = &types.ConversionRate{
		Base:       "eth",
		Foreign:    "usdt",
		Rate:       3000,
		OracleName: "cooinmarketcap",
		Time:       time.Time{}.UnixMilli(),
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
