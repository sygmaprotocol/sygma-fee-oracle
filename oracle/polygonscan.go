// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/oracle/client"
	"github.com/ChainSafe/sygma-fee-oracle/types"
	"github.com/ChainSafe/sygma-fee-oracle/util"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ GasPriceOracle = (*Polygonscan)(nil)

type Polygonscan struct {
	log *logrus.Entry

	name   string
	apiKey string
	enable bool
	apis   PolygonscanApis
}

type PolygonscanApis struct {
	GasPriceRequest string
}

type PolygonscanResp struct {
	Statue  string      `json:"statue"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type PolygonscanGasPricesResp struct {
	FastGasPrice    string `mapstructure:"FastGasPrice"`
	LastBlock       string `mapstructure:"LastBlock"`
	ProposeGasPrice string `mapstructure:"ProposeGasPrice"`
	SafeGasPrice    string `mapstructure:"SafeGasPrice"`
	GasUsedRatio    string `mapstructure:"gasUsedRatio"`
	SuggestBaseFee  string `mapstructure:"suggestBaseFee"`
}

func NewPolygonscan(conf *config.Config, log *logrus.Entry) *Polygonscan {
	return &Polygonscan{
		log:    log.WithField("services", "polygonscan"),
		name:   "polygonscan",
		apiKey: conf.OracleConfig().Polygonscan.ApiKey,
		enable: conf.OracleConfig().Polygonscan.Enable,
		apis: PolygonscanApis{
			GasPriceRequest: fmt.Sprintf("%s%s", conf.OracleConfig().Polygonscan.Apis.GasPriceApiUrl, conf.OracleConfig().Polygonscan.ApiKey),
		},
	}
}

func (p *Polygonscan) InquiryGasPrice(domainName string) (*types.GasPrices, error) {
	if strings.ToLower(domainName) != "polygon" {
		return nil, ErrNotSupported
	}

	statusCode, body, err := client.NewHttpRequestMessage(http.MethodGet, p.apis.GasPriceRequest,
		nil, nil, p.log).Request()
	if err != nil || statusCode != http.StatusOK {
		p.log.Errorf("query gas price err: %v, statusCode: %d", err, statusCode)
		return nil, errors.Wrap(err, "failed to query gas price")
	}

	pr := &PolygonscanResp{}
	prResult, err := pr.parsePolygonscanResp(body)
	if err != nil {
		p.log.Errorf("unmarshal response body err: %s ", err.Error())
		return nil, errors.Wrap(err, "failed to unmarshal resp body of querying gas price")
	}

	if prResult.Message != "OK" {
		return nil, errors.Errorf("failed to fetch gas price: %s", prResult.Result)
	}

	var pgp PolygonscanGasPricesResp
	err = mapstructure.Decode(pr.Result, &pgp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode gas price response")
	}

	safeGasPriceValue, err := util.Large2SmallUnitConverter(pgp.SafeGasPrice, types.GWeiDecimal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert safe gasprice")
	}
	proposeGasPriceValue, err := util.Large2SmallUnitConverter(pgp.ProposeGasPrice, types.GWeiDecimal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert propose gasprice")
	}
	fastGasPriceValue, err := util.Large2SmallUnitConverter(pgp.FastGasPrice, types.GWeiDecimal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert fast gasprice")
	}

	return &types.GasPrices{
		SafeGasPrice:    safeGasPriceValue.String(),
		ProposeGasPrice: proposeGasPriceValue.String(),
		FastGasPrice:    fastGasPriceValue.String(),
		OracleName:      p.name,
		DomainName:      domainName,
		Time:            time.Now().Unix(),
	}, nil
}

func (p *Polygonscan) Name() string {
	return p.name
}

func (p *Polygonscan) IsEnabled() bool {
	return p.enable
}

func (pr *PolygonscanResp) parsePolygonscanResp(body []byte) (PolygonscanResp, error) {
	err := json.Unmarshal(body, pr)
	if err != nil {
		return PolygonscanResp{}, err
	}
	return *pr, nil
}
