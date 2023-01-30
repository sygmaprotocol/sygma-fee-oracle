// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package api

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ChainSafe/sygma-fee-oracle/util"
	"github.com/pkg/errors"
)

func (h *Handler) rateSignature(result *FetchRateResp, fromDomainID int, resourceID string) (string, error) {
	fromDomainBaseCurrencyResourceId := config.ResourceIDBuilder(config.NativeCurrencyAddr, fromDomainID)
	fromDomainBaseCurrencyDomainInfo, err := h.conf.ResourceDomainInfo(fromDomainBaseCurrencyResourceId, fromDomainID)
	if err == nil {
		return "", err
	}
	baseRate, err := util.Large2SmallUnitConverter(result.BaseRate, uint(fromDomainBaseCurrencyDomainInfo.Decimals))
	if err != nil {
		return "", errors.Wrap(err, "failed to convert BaseRate")
	}
	finalBaseEffectiveRate := util.PaddingZero(baseRate.Bytes(), 32)

	tokenRateCurrencyDomainInfo, err := h.conf.ResourceDomainInfo(resourceID, fromDomainID)
	if err == nil {
		return "", err
	}
	tokenRate, err := util.Large2SmallUnitConverter(result.TokenRate, uint(tokenRateCurrencyDomainInfo.Decimals))
	if err != nil {
		return "", errors.Wrap(err, "failed to convert TokenRate")
	}
	finalTokenEffectiveRate := util.PaddingZero(tokenRate.Bytes(), 32)

	gasPrice, err := util.Str2BigInt(result.DestinationChainGasPrice)
	if err != nil {
		return "", errors.Wrap(err, "failed to convert DestinationChainGasPrice")
	}
	finalGasPrice := util.PaddingZero(gasPrice.Bytes(), 32)

	finalTimestamp := fmt.Sprintf("%064x", result.ExpirationTimestamp)
	finalTimestampBytes, err := hex.DecodeString(finalTimestamp)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode timestamp")
	}

	finalFromDomainId := util.PaddingZero([]byte{uint8(result.FromDomainID)}, 32)
	finalToDomainId := util.PaddingZero([]byte{uint8(result.ToDomainID)}, 32)

	finalResourceId, err := hex.DecodeString(resourceID[2:])
	if err != nil {
		return "", errors.Wrap(err, "failed to decode resourceID")
	}

	finalMsgGasLimit := fmt.Sprintf("%064x", result.MsgGasLimit)
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

	signature, err := h.identity.Sign(feeDataRaw)
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

// TODO: this is the placeholder for the algorithm of calculating the data expiration
func (h *Handler) dataExpirationManager(baseTimestamp int64) int64 {
	return baseTimestamp + h.conf.DataValidIntervalConfig()
}

func (h *Handler) parseDomains(fromID string, toID string) (from *config.Domain, to *config.Domain, err error) {
	fromDomainID, err := strconv.Atoi(fromID)
	if err != nil {
		return from, to, err
	}
	toDomainID, err := strconv.Atoi(toID)
	if err != nil {
		return from, to, err
	}
	if fromDomainID == toDomainID {
		return from, to, fmt.Errorf("from and to domain equal")
	}

	toDomain := h.conf.GetRegisteredDomains(toDomainID)
	if toDomain == nil {
		return from, to, fmt.Errorf("to domain not registered")
	}
	fromDomain := h.conf.GetRegisteredDomains(fromDomainID)
	if fromDomain == nil {
		return from, to, fmt.Errorf("from domain not registered")
	}

	return from, to, nil
}

func (h *Handler) parseResource(resourceID string) (*config.Resource, error) {
	if !strings.HasPrefix(resourceID, "0x") || len(resourceID) != 66 {
		return &config.Resource{}, fmt.Errorf("resource invalid")
	}

	resource := h.conf.GetRegisteredResources(resourceID)
	if resource == nil {
		return &config.Resource{}, fmt.Errorf("resource not registerred")
	}

	return resource, nil
}

func (h *Handler) calculateTokenRate(resource *config.Resource, ber *types.ConversionRate, from *config.Domain, to *config.Domain) (*types.ConversionRate, error) {
	ter := &types.ConversionRate{}
	// if resource is the base currency of the fromDomain
	if resource.Symbol == from.BaseCurrencySymbol {
		ter = ber
	} else {

		if resource.Symbol == to.BaseCurrencySymbol {
			ter.Rate = 1.0
		} else {
			ter, err := h.consensus.FilterLocalConversionRateData(h.conversionRateStore, to.BaseCurrencySymbol, resource.Symbol)
			if err != nil {
				return ter, err
			}
		}
	}

	return ter, nil
}
