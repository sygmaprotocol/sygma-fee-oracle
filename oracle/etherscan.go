// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package oracle

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/ChainSafe/chainbridge-fee-oracle/oracle/client"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/ChainSafe/chainbridge-fee-oracle/util"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ GasPriceOracle = (*Etherscan)(nil)

type Etherscan struct {
	log *logrus.Entry

	name   string
	apiKey string
	enable bool
	apis   EtherscanApis
}

type EtherscanApis struct {
	GasPriceRequest string
}

type EtherscanResp struct {
	Statue  string      `json:"statue"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type EtherscanGasPricesResp struct {
	FastGasPrice    string `mapstructure:"FastGasPrice"`
	LastBlock       string `mapstructure:"LastBlock"`
	ProposeGasPrice string `mapstructure:"ProposeGasPrice"`
	SafeGasPrice    string `mapstructure:"SafeGasPrice"`
	GasUsedRatio    string `mapstructure:"gasUsedRatio"`
	SuggestBaseFee  string `mapstructure:"suggestBaseFee"`
}

func NewEtherscan(conf *config.Config, log *logrus.Entry) *Etherscan {
	return &Etherscan{
		log:    log.WithField("services", "etherscan"),
		name:   "etherscan",
		apiKey: conf.OracleConfig().Etherscan.ApiKey,
		enable: conf.OracleConfig().Etherscan.Enable,
		apis: EtherscanApis{
			GasPriceRequest: fmt.Sprintf("%s%s", conf.OracleConfig().Etherscan.Apis.GasPriceApiUrl, conf.OracleConfig().Etherscan.ApiKey),
		},
	}
}

func (e *Etherscan) InquiryGasPrice(domainName string) (*types.GasPrices, error) {
	if strings.ToLower(domainName) != "ethereum" {
		return nil, ErrNotSupported
	}

	statusCode, body, err := client.NewHttpRequestMessage(http.MethodGet, e.apis.GasPriceRequest,
		nil, nil, e.log).Request()
	if err != nil || statusCode != http.StatusOK {
		e.log.Errorf("query gas price err: %v, statusCode: %d", err, statusCode)
		return nil, errors.Wrap(err, "failed to query gas price")
	}

	er := &EtherscanResp{}
	erResult, err := er.parseEtherscanResp(body)
	if err != nil {
		e.log.Errorf("unmarshal response body err: %s ", err.Error())
		return nil, errors.Wrap(err, "failed to unmarshal resp body of querying gas price")
	}

	if erResult.Message != "OK" {
		return nil, errors.Errorf("failed to fetch gas price: %s", erResult.Result)
	}

	var egp EtherscanGasPricesResp
	err = mapstructure.Decode(er.Result, &egp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode gas price response")
	}

	safeGasPriceValue, err := util.Str2BigInt(egp.SafeGasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert gas price response value")
	}
	proposeGasPriceValue, err := util.Str2BigInt(egp.ProposeGasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert gas price response value")
	}
	fastGasPriceValue, err := util.Str2BigInt(egp.FastGasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert gas price response value")
	}

	return &types.GasPrices{
		SafeGasPrice:    new(big.Int).Mul(safeGasPriceValue, big.NewInt(types.GWei)).String(),
		ProposeGasPrice: new(big.Int).Mul(proposeGasPriceValue, big.NewInt(types.GWei)).String(),
		FastGasPrice:    new(big.Int).Mul(fastGasPriceValue, big.NewInt(types.GWei)).String(),
		OracleName:      e.name,
		DomainName:      domainName,
		Time:            time.Now().Unix(),
	}, nil
}

func (e *Etherscan) Name() string {
	return e.name
}

func (e *Etherscan) IsEnabled() bool {
	return e.enable
}

func (er *EtherscanResp) parseEtherscanResp(body []byte) (EtherscanResp, error) {
	err := json.Unmarshal(body, er)
	if err != nil {
		return EtherscanResp{}, err
	}
	return *er, nil
}
