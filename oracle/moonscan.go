// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/oracle/client"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ GasPriceOracle = (*Moonscan)(nil)

type Moonscan struct {
	log *logrus.Entry

	name              string
	apiKey            string
	enable            bool
	apis              MoonscanApis
	GasPriceDomainIDs []int
}

type MoonscanApis struct {
	GasPriceRequest string
}

type MoonscanResp struct {
	Statue  string `json:"statue"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func NewMoonscan(conf *config.Config, log *logrus.Entry) *Moonscan {
	return &Moonscan{
		log:    log.WithField("services", "moonscan"),
		name:   "moonscan",
		apiKey: conf.OracleConfig().Moonscan.ApiKey,
		enable: conf.OracleConfig().Moonscan.Enable,
		apis: MoonscanApis{
			GasPriceRequest: fmt.Sprintf("%s%s", conf.OracleConfig().Moonscan.Apis.GasPriceApiUrl, conf.OracleConfig().Moonscan.ApiKey),
		},
		GasPriceDomainIDs: conf.OracleConfig().Moonscan.GasPriceDomainIds,
	}
}

func (m *Moonscan) InquiryGasPrice(domainID int) (*types.GasPrices, error) {
	// Moonscan doesn't support GasTracker API like Etherscan and Polygonscan does,
	// so we will use eth_gasPrice RPC call here
	// the JSON2.0 response will be different: if Message is empty, it means a successful call
	statusCode, body, err := client.NewHttpRequestMessage(http.MethodGet, m.apis.GasPriceRequest,
		nil, nil, m.log).Request()
	if err != nil || statusCode != http.StatusOK {
		m.log.Errorf("query gas price err: %v, statusCode: %d", err, statusCode)
		return nil, errors.Wrap(err, "failed to query gas price")
	}

	pr := &MoonscanResp{}
	prResult, err := pr.parseMoonscanResp(body)
	if err != nil {
		m.log.Errorf("unmarshal response body err: %s ", err.Error())
		return nil, errors.Wrap(err, "failed to unmarshal resp body of querying gas price")
	}

	if prResult.Message != "" {
		return nil, errors.Errorf("failed to fetch gas price: %s", prResult.Result)
	}

	gp, ok := big.NewInt(0).SetString(pr.Result[2:], 16)
	if !ok {
		return nil, errors.Wrap(err, "failed to decode gas price response")
	}

	return &types.GasPrices{
		SafeGasPrice:    gp.String(),
		ProposeGasPrice: gp.String(),
		FastGasPrice:    gp.String(),
		OracleName:      m.name,
		DomainID:        domainID,
		Time:            time.Now().Unix(),
	}, nil
}

func (m *Moonscan) Name() string {
	return m.name
}

func (m *Moonscan) IsEnabled() bool {
	return m.enable
}

func (m *Moonscan) SupportedGasPriceDomainIds() []int {
	return m.GasPriceDomainIDs
}

func (mr *MoonscanResp) parseMoonscanResp(body []byte) (MoonscanResp, error) {
	err := json.Unmarshal(body, mr)
	if err != nil {
		return MoonscanResp{}, err
	}
	return *mr, nil
}
