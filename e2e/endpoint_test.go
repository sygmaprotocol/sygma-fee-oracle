// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package e2e

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ChainSafe/sygma-fee-oracle/identity/secp256k1"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ChainSafe/sygma-fee-oracle/e2e/setup"
	"github.com/stretchr/testify/suite"

	"github.com/ChainSafe/sygma-fee-oracle/util"
)

type apiRespGeneral struct {
	Response types.Rate `json:"response"`
}

type SignatureVerificationTestSuite struct {
	suite.Suite
	contractSetup *setup.ContractsSetupResp
}

func TestRuSignatureVerificationTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(SignatureVerificationTestSuite))
}

func (s *SignatureVerificationTestSuite) SetupSuite() {
	s.contractSetup = setup.ContractsSetup()
}

func (s *SignatureVerificationTestSuite) TearDownSuite() {}

func (s *SignatureVerificationTestSuite) SetupTest() {}

func (s *SignatureVerificationTestSuite) TearDownTest() {}

func (s *SignatureVerificationTestSuite) TestSignatureVerification_CalculateFee() {
	// test related params define
	amount := big.NewInt(1000000000000000000) // 1 ether
	finalAmount := util.PaddingZero(amount.Bytes(), 32)
	emptySetResourceData := []byte{}

	// fee oracle endpoints request
	resp, err := http.Get("http://127.0.0.1:8091/v1/rate/from/1/to/2/resourceid/0x0000000000000000000000000000000000000000000000000000000000000300")
	s.Nil(err)

	body, err := ioutil.ReadAll(resp.Body)
	s.Nil(err)

	var response *apiRespGeneral
	err = json.Unmarshal(body, &response)
	s.Nil(err)

	// prepare the test data
	ber, err := util.Large2SmallUnitConverter(response.Response.BaseRate, 18)
	s.Nil(err)
	finalBaseEffectiveRate := util.PaddingZero(ber.Bytes(), 32)
	ter, err := util.Large2SmallUnitConverter(response.Response.TokenRate, 8)
	s.Nil(err)
	finalTokenEffectiveRate := util.PaddingZero(ter.Bytes(), 32)

	gasPrice, err := util.Str2BigInt(response.Response.DestinationChainGasPrice)
	s.Nil(err)

	finalGasPrice := util.PaddingZero(gasPrice.Bytes(), 32)
	finalTimestamp := fmt.Sprintf("%064x", response.Response.ExpirationTimestamp+1000000000)

	finalTimestampBytes, err := hex.DecodeString(finalTimestamp)
	s.Nil(err)

	finalFromDomainId := util.PaddingZero([]byte{uint8(response.Response.FromDomainID)}, 32)
	finalToDomainId := util.PaddingZero([]byte{uint8(response.Response.ToDomainID)}, 32)

	finalResourceId, err := hex.DecodeString(response.Response.ResourceID[2:])
	s.Nil(err)

	gasLimit, _ := strconv.ParseUint(response.Response.MsgGasLimit, 10, 64)
	finalMsgGasLimit := fmt.Sprintf("%064x", gasLimit)
	finalMsgGasLimitBytes, err := hex.DecodeString(finalMsgGasLimit)
	s.Nil(err)

	_, err = s.contractSetup.BridgeInstance.AdminSetResource(setup.IncreaseNonce(s.contractSetup.Auth), s.contractSetup.ERC20HandlerAddress, util.Byte32Converter(finalResourceId), s.contractSetup.ERC20PresetMinterPauserAddress, emptySetResourceData)
	s.Nil(err)

	_, err = s.contractSetup.FeeHandlerRouterInstance.AdminSetResourceHandler(setup.IncreaseNonce(s.contractSetup.Auth), uint8(setup.FromDomainId), util.Byte32Converter(finalResourceId), s.contractSetup.DynamicFeeHandlerAddress)
	s.Nil(err)

	// assembly data
	feeDataMessageByte := bytes.Buffer{}
	feeDataMessageByte.Write(finalBaseEffectiveRate)
	feeDataMessageByte.Write(finalTokenEffectiveRate)
	feeDataMessageByte.Write(finalGasPrice)
	feeDataMessageByte.Write(finalTimestampBytes)
	feeDataMessageByte.Write(finalFromDomainId)
	feeDataMessageByte.Write(finalToDomainId)
	feeDataMessageByte.Write(finalResourceId)
	feeDataMessageByte.Write(finalMsgGasLimitBytes)
	finalFeeDataMessage := feeDataMessageByte.Bytes()

	// prepare to sign
	prikeyBytes, err := hex.DecodeString(setup.FeeOracleHexPriKey)
	s.Nil(err)

	feeOracleKey, err := secp256k1.NewKeypairFromPrivateKey(prikeyBytes)
	s.Nil(err)

	s.EqualValues(setup.FeeOracleAddress.String(), feeOracleKey.Address(), "feeOracleAddress does not match feeOracleKey")

	sig, err := feeOracleKey.Sign(finalFeeDataMessage)
	s.Nil(err)
	originalSig := sig

	// modify v
	sigb := bytes.Buffer{}
	sigb.Write(sig[:64])
	sigb.WriteByte(byte(int8(sig[64]) + 27)) // openzepplin is verifying if v is 27/28, need to manually add 27 to v
	sig = sigb.Bytes()

	// assembly feeData
	// 256 + 65 + 32
	feeData := bytes.Buffer{}
	feeData.Write(finalFeeDataMessage)
	feeData.Write(sig)
	feeData.Write(finalAmount) // not used in feeData, moved to depositData
	s.EqualValues(353, feeData.Len(), "invalid feeData")

	s.True(bytes.Equal(feeData.Bytes()[:256], finalFeeDataMessage))
	s.True(bytes.Equal(feeData.Bytes()[256:321], sig))
	s.True(bytes.Equal(feeData.Bytes()[321:], finalAmount))

	// deposit data
	depositData := bytes.Buffer{}
	depositData.Write(finalAmount)

	// ecrecover verify signer
	signer, err := crypto.Ecrecover(crypto.Keccak256(finalFeeDataMessage), originalSig)
	if err != nil {
		panic(err)
	}
	signerPubKey, err := crypto.UnmarshalPubkey(signer)
	if err != nil {
		panic(err)
	}
	if crypto.PubkeyToAddress(*signerPubKey).String() != feeOracleKey.Address() {
		panic("signer does not match")
	}

	// calling actual test: CalculateFee in Fee handler contract
	result, err := s.contractSetup.DynamicFeeHandlerInstance.CalculateFee(&bind.CallOpts{From: s.contractSetup.Auth.From}, setup.SenderAddress, uint8(setup.FromDomainId), uint8(setup.ToDomainID), util.Byte32Converter(finalResourceId), depositData.Bytes(), feeData.Bytes())
	s.Nil(err)

	// test result verification
	s.EqualValues(result.TokenAddress.String(), s.contractSetup.ERC20PresetMinterPauserAddress.String(), "resource token address does not match")
	s.EqualValues("50000000000000000", result.Fee.String(), "fee does not match")
}
