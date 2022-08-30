// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ChainSafe/sygma-fee-oracle/base"
	"github.com/ChainSafe/sygma-fee-oracle/oracle"
	mockOracle "github.com/ChainSafe/sygma-fee-oracle/oracle/mock"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type GasPriceOracleTestSuite struct {
	suite.Suite
	appBase          *base.FeeOracleAppBase
	gasPriceOperator *oracle.GasPriceOracleOperator
	oracle           *mockOracle.MockGasPriceOracle
	testdata         *types.GasPrices
}

func TestRunGasPriceOracleTestSuite(t *testing.T) {
	suite.Run(t, new(GasPriceOracleTestSuite))
}

func (s *GasPriceOracleTestSuite) SetupSuite() {
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

func (s *GasPriceOracleTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
	_ = os.RemoveAll("./keyfile.priv")
}

func (s *GasPriceOracleTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.oracle = mockOracle.NewMockGasPriceOracle(gomockController)
	s.appBase = base.NewFeeOracleAppBase("../config/config.template.yaml", "../domain.json", "../resource.json", "./keyfile.priv", "secp256k1")
	s.gasPriceOperator = oracle.NewGasPriceOracleOperator(s.appBase.GetLogger(), s.oracle)
	s.testdata = &types.GasPrices{
		SafeGasPrice:    "1",
		ProposeGasPrice: "2",
		FastGasPrice:    "3",
		OracleName:      "etherscan",
		DomainName:      "ethereum",
		Time:            time.Now().UnixMilli(),
	}
}

func (s *GasPriceOracleTestSuite) TearDownTest() {
	_ = s.appBase.GetStore().Close()
}

func (s *GasPriceOracleTestSuite) TestInquiryGasPrice_Failure() {
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainName).Return(nil, errors.New("error"))

	_, err := s.gasPriceOperator.Run(s.testdata.DomainName)

	s.NotNil(err)
}

func (s *GasPriceOracleTestSuite) TestInquiryGasPrice_Success() {
	s.oracle.EXPECT().InquiryGasPrice(s.testdata.DomainName).Return(s.testdata, nil)

	_, err := s.gasPriceOperator.Run(s.testdata.DomainName)

	s.Nil(err)
}
