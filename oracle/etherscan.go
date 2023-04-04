// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/oracle/client"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ GasPriceOracle = (*Etherscan)(nil)

type Etherscan struct {
	log *logrus.Entry

	source   string
	apiKey   string
	enable   bool
	apis     EtherscanApis
	domainID int
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

func NewEtherscan(oracleSource string, oracle config.Oracle, domainID int, log *logrus.Entry) *Etherscan {
	return &Etherscan{
		log:    log.WithField("services", oracleSource),
		source: oracleSource,
		apiKey: oracle.ApiKey,
		enable: oracle.Enable,
		apis: EtherscanApis{
			GasPriceRequest: fmt.Sprintf("%s%s", oracle.URL, oracle.ApiKey),
		},
		domainID: domainID,
	}
}

func (e *Etherscan) InquiryGasPrice() (*types.GasPrices, error) {
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

	return &types.GasPrices{
		SafeGasPrice:    egp.SafeGasPrice,
		ProposeGasPrice: egp.ProposeGasPrice,
		FastGasPrice:    egp.FastGasPrice,
		OracleSource:    e.source,
		DomainID:        e.domainID,
		Time:            time.Now().Unix(),
	}, nil
}

func (e *Etherscan) Source() string {
	return e.source
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
