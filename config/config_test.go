// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config_test

import (
	"errors"
	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	mockremoteparam "github.com/ChainSafe/chainbridge-fee-oracle/config/mock"
	"github.com/ChainSafe/chainbridge-fee-oracle/remoteParam"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type RemoteParamOperatorTestSuite struct {
	suite.Suite
	remoteParamOperator *mockremoteparam.MockRemoteParamOperator
	appBase             *base.FeeOracleAppBase
	conf                *config.Config
}

func TestRunConversionRateStoreTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteParamOperatorTestSuite))
}

func (s *RemoteParamOperatorTestSuite) SetupSuite() {
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

func (s *RemoteParamOperatorTestSuite) TearDownSuite() {
	_ = os.RemoveAll("./lvldbdata")
	_ = os.RemoveAll("./keyfile.priv")
}

func (s *RemoteParamOperatorTestSuite) SetupTest() {
	gomockController := gomock.NewController(s.T())
	s.remoteParamOperator = mockremoteparam.NewMockRemoteParamOperator(gomockController)
	s.appBase = base.NewFeeOracleAppBase("./config.template.yaml", "../domain.json", "../resource.json", "./keyfile.priv", "secp256k1")
	s.conf = s.appBase.GetConfig()
}

func (s *RemoteParamOperatorTestSuite) TearDownTest() {
	_ = s.appBase.GetStore().Close()
}

func (s *RemoteParamOperatorTestSuite) TestSetRemoteParams_operatorIsNil() {
	s.remoteParamOperator.EXPECT().LoadParameter("/chainbridge/fee-oracle/domainData").Times(0)
	s.NotPanics(func() { s.conf.SetRemoteParams(nil) }, "should not panic")
}

func (s *RemoteParamOperatorTestSuite) TestSetRemoteParams_Failure() {
	s.remoteParamOperator.EXPECT().LoadParameter("/chainbridge/fee-oracle/domainData").Return(nil, errors.New("error"))
	s.Panics(func() { s.conf.SetRemoteParams(s.remoteParamOperator) }, "should panic on non existing param name")
}

func (s *RemoteParamOperatorTestSuite) TestSetRemoteParams_Success() {
	s.Nil(s.conf.GetRegisteredDomains(10))
	s.Nil(s.conf.GetRegisteredResources(config.ResourceIDBuilder("0xfC2e2618147813E510CFc92747f0D09C14A653c5", 3)))

	// mock domain data
	// {
	//  "domains": [
	//    {
	//      "id": 10,
	//      "name": "ethereum",
	//      "baseCurrencyFullName": "ether",
	//      "baseCurrencySymbol": "eth",
	//      "addressPrefix": "0x"
	//    }
	//  ]
	//}
	domainData := "{\n  \"domains\": [\n    {\n      \"id\": 10,\n      \"name\": \"ethereum\",\n      \"baseCurrencyFullName\": \"ether\",\n      \"baseCurrencySymbol\": \"eth\",\n      \"addressPrefix\": \"0x\"\n    }\n  ]\n}"

	// mock resource data
	// {
	//  "resources": [
	//    {
	//      "symbol": "eth",
	//      "decimal": 18,
	//      "tokenAddress": "0xfC2e2618147813E510CFc92747f0D09C14A653c5",
	//      "domainId": 3
	//    }
	//  ]
	//}
	resourceData := "{\n  \"resources\": [\n    {\n      \"symbol\": \"eth\",\n      \"decimal\": 18,\n      \"tokenAddress\": \"0xfC2e2618147813E510CFc92747f0D09C14A653c5\",\n      \"domainId\": 3\n    }\n  ]\n}"

	s.remoteParamOperator.EXPECT().LoadParameter("/chainbridge/fee-oracle/domainData").Times(1).Return(&remoteParam.RemoteParamResult{Value: domainData}, nil)
	s.remoteParamOperator.EXPECT().LoadParameter("/chainbridge/fee-oracle/resourceData").Times(1).Return(&remoteParam.RemoteParamResult{Value: resourceData}, nil)

	s.NotPanics(func() { s.conf.SetRemoteParams(s.remoteParamOperator) }, "should not panic")

	s.NotNil(s.conf.GetRegisteredDomains(10))
	s.NotNil(s.conf.GetRegisteredResources(config.ResourceIDBuilder("0xfC2e2618147813E510CFc92747f0D09C14A653c5", 3)))
}
