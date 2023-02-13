// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package signature

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ChainSafe/sygma-fee-oracle/util"
	"github.com/pkg/errors"
)

func RateSignature(conf *config.Config, rate *types.Rate, identity *identity.OracleIdentityOperator, fromDomainID int, resourceID string) (string, error) {
	fromDomainBaseCurrencyResourceId := config.ResourceIDBuilder(config.NativeCurrencyAddr, fromDomainID)
	fromDomainBaseCurrencyDomainInfo, err := conf.ResourceDomainInfo(fromDomainBaseCurrencyResourceId, fromDomainID)
	if err != nil {
		return "", err
	}
	baseRate, err := util.Large2SmallUnitConverter(rate.BaseRate, uint(fromDomainBaseCurrencyDomainInfo.Decimals))
	if err != nil {
		return "", errors.Wrap(err, "failed to convert BaseRate")
	}
	finalBaseEffectiveRate := util.PaddingZero(baseRate.Bytes(), 32)

	tokenRateCurrencyDomainInfo, err := conf.ResourceDomainInfo(resourceID, fromDomainID)
	if err != nil {
		return "", err
	}
	tokenRate, err := util.Large2SmallUnitConverter(rate.TokenRate, uint(tokenRateCurrencyDomainInfo.Decimals))
	if err != nil {
		return "", errors.Wrap(err, "failed to convert TokenRate")
	}
	finalTokenEffectiveRate := util.PaddingZero(tokenRate.Bytes(), 32)

	gasPrice, err := util.Str2BigInt(rate.DestinationChainGasPrice)
	if err != nil {
		return "", errors.Wrap(err, "failed to convert DestinationChainGasPrice")
	}
	finalGasPrice := util.PaddingZero(gasPrice.Bytes(), 32)

	finalTimestamp := fmt.Sprintf("%064x", rate.ExpirationTimestamp)
	finalTimestampBytes, err := hex.DecodeString(finalTimestamp)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode timestamp")
	}

	finalFromDomainId := util.PaddingZero([]byte{uint8(rate.FromDomainID)}, 32)
	finalToDomainId := util.PaddingZero([]byte{uint8(rate.ToDomainID)}, 32)

	finalResourceId, err := hex.DecodeString(resourceID[2:])
	if err != nil {
		return "", errors.Wrap(err, "failed to decode resourceID")
	}

	gasLimit, err := strconv.ParseUint(rate.MsgGasLimit, 10, 64)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse MsgGasLimit")
	}
	finalMsgGasLimit := fmt.Sprintf("%064x", gasLimit)
	finalMsgGasLimitBytes, err := hex.DecodeString(finalMsgGasLimit)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode MsgGasLimit")
	}

	feeDataMessageByte := bytes.Buffer{}
	feeDataMessageByte.Write(finalBaseEffectiveRate)
	feeDataMessageByte.Write(finalTokenEffectiveRate)
	feeDataMessageByte.Write(finalGasPrice)
	feeDataMessageByte.Write(finalTimestampBytes)
	feeDataMessageByte.Write(finalFromDomainId)
	feeDataMessageByte.Write(finalToDomainId)
	feeDataMessageByte.Write(finalResourceId)
	feeDataMessageByte.Write(finalMsgGasLimitBytes)
	feeDataRaw := feeDataMessageByte.Bytes()

	signature, err := identity.Sign(feeDataRaw)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign")
	}

	// modify v
	// openzepplin is verifying if v is 27/28, need to manually add 27 to v
	sigb := bytes.Buffer{}
	sigb.Write(signature[:64])
	sigb.WriteByte(byte(int8(signature[64]) + 27))
	signature = sigb.Bytes()

	return hex.EncodeToString(signature), nil
}
