// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package signature_test

import (
	"encoding/hex"
	"testing"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/identity/secp256k1"
	"github.com/ChainSafe/sygma-fee-oracle/signature"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/stretchr/testify/suite"
)

type SignatureTestSuite struct {
	suite.Suite
	oracleIdentity *identity.OracleIdentityOperator
}

func TestRunSignatureTestSuite(t *testing.T) {
	suite.Run(t, new(SignatureTestSuite))
}

func (s *SignatureTestSuite) SetupTest() {
	key, _ := hex.DecodeString("21e153aaaaa05d76aed18a8217d36df36a6179931f09a37a33972227b270e1b8")
	keypair, _ := secp256k1.NewKeypairFromPrivateKey(key)
	s.oracleIdentity = identity.NewOracleIdentityOperator(keypair)
}

func (s *SignatureTestSuite) Test_RateSignature_InvalidConfig() {
	fromDomainID := 1
	conf := &config.Config{}
	rate := &types.Rate{
		FromDomainID:             1,
		ToDomainID:               2,
		BaseRate:                 "1.345",
		TokenRate:                "0.112",
		ExpirationTimestamp:      1000,
		DestinationChainGasPrice: "1000",
		MsgGasLimit:              "1000000",
	}

	_, err := signature.RateSignature(conf, rate, s.oracleIdentity, fromDomainID, "")

	s.NotNil(err)
}

func (s *SignatureTestSuite) Test_RateSignature_MissingDomainInfo() {
	fromDomainID := 1
	nativeResourceID := config.ResourceIDBuilder(config.NativeCurrencyAddr, fromDomainID)
	resourceID := "0x0000000000000000000000000000000000000000000000000000000000000001"
	nativeResource := config.Resource{
		ID: nativeResourceID,
	}
	targetResource := config.Resource{
		ID: resourceID,
	}
	resources := make(map[string]*config.Resource, 0)
	resources[nativeResourceID] = &nativeResource
	resources[resourceID] = &targetResource
	conf := &config.Config{
		Resources: resources,
	}

	rate := &types.Rate{
		FromDomainID:             1,
		ToDomainID:               2,
		BaseRate:                 "1.345",
		TokenRate:                "0.112",
		ExpirationTimestamp:      1000,
		DestinationChainGasPrice: "1000",
		MsgGasLimit:              "1000000",
	}

	_, err := signature.RateSignature(conf, rate, s.oracleIdentity, fromDomainID, resourceID)

	s.NotNil(err)
}

func (s *SignatureTestSuite) Test_RateSignature_ValidSignature() {
	fromDomainID := 1
	domainMap := make(map[int]*config.DomainInfo, 0)
	domainMap[fromDomainID] = &config.DomainInfo{
		Decimals: 18,
	}
	nativeResourceID := config.ResourceIDBuilder(config.NativeCurrencyAddr, fromDomainID)
	resourceID := "0x0000000000000000000000000000000000000000000000000000000000000001"
	nativeResource := config.Resource{
		ID:         nativeResourceID,
		DomainInfo: domainMap,
	}
	targetResource := config.Resource{
		ID:         resourceID,
		DomainInfo: domainMap,
	}
	resources := make(map[string]*config.Resource, 0)
	resources[nativeResourceID] = &nativeResource
	resources[resourceID] = &targetResource
	conf := &config.Config{
		Resources: resources,
	}
	rate := &types.Rate{
		FromDomainID:             1,
		ToDomainID:               2,
		BaseRate:                 "1.345",
		TokenRate:                "0.112",
		ExpirationTimestamp:      1000,
		DestinationChainGasPrice: "1000",
		MsgGasLimit:              "1000000",
	}

	signature, err := signature.RateSignature(conf, rate, s.oracleIdentity, fromDomainID, resourceID)

	s.Nil(err)
	s.Equal(signature, "7f628bdc854780f8ca8e64a8a75122a905d0538518d6ca64ef5a9077f14bd2c9460f0ecffce957a347cd6b6db0c212456f533d5f397599a1ad420c87a5eb49321c")
}
