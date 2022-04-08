package oracle

import (
	"encoding/json"
	"fmt"
	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/ChainSafe/chainbridge-fee-oracle/oracle/client"

	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var _ ConversionRateOracle = (*CoinMarketCap)(nil)

type CoinMarketCap struct {
	log *logrus.Entry

	name   string
	apiKey string
	enable bool
	apis   CoinMarketCapApis
}

type CoinMarketCapApis struct {
	ConversionRateRequest string
}

type CoinMarketCapResp struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
	Data         interface{} `json:"data"`
}

type CoinMarketCapConversionRateQuote struct {
	Price       float64 `mapstructure:"price"`
	Volume24h   float64 `mapstructure:"volume_24h"`
	MarketCap   float64 `mapstructure:"market_cap"`
	LastUpdated string  `mapstructure:"last_updated"`
}

func NewCoinMarketCap(conf *config.Config, log *logrus.Entry) *CoinMarketCap {
	return &CoinMarketCap{
		log:    log.WithField("services", "coinMarketCap"),
		name:   "coinmarketcap",
		apiKey: conf.OracleConfig().CoinMarketCap.ApiKey,
		enable: conf.OracleConfig().CoinMarketCap.Enable,
		apis: CoinMarketCapApis{
			ConversionRateRequest: conf.OracleConfig().CoinMarketCap.Apis.QueryRate,
		},
	}
}

func (c *CoinMarketCap) InquiryConversionRate(baseCurrency, foreignCurrency string) (conversionRateResp *types.ConversionRate, returnErr error) {
	defer func() {
		if err := recover(); err != nil {
			returnErr = recoverErrorStringifier(err)
		}
	}()

	url := fmt.Sprintf("%ssymbol=%s&convert=%s", c.apis.ConversionRateRequest, baseCurrency, foreignCurrency)

	header := make(map[string]string)
	header["X-CMC_PRO_API_KEY"] = c.apiKey

	statusCode, body, err := client.NewHttpRequestMessage(http.MethodGet, url, header, nil, c.log).Request()
	if err != nil || statusCode != http.StatusOK {
		c.log.Errorf("query conversion rate err: %v, statusCode: %d", err, statusCode)
		return nil, errors.Wrap(err, "failed to query conversion rate")
	}

	cmcp := &CoinMarketCapResp{}
	cmcpResult, err := cmcp.parseCoinMarketCapResp(body)
	if err != nil {
		c.log.Errorf("unmarshal response body err: %s ", err.Error())
		return nil, errors.Wrap(err, "failed to unmarshal resp body of querying conversion rate")
	}

	if cmcpResult.ErrorCode != 0 {
		return nil, errors.Errorf("failed to fetch conversion rate: %s", cmcp.ErrorMessage)
	}

	// parse response data
	data := cmcp.Data.(map[string]interface{})
	baseData := data[strings.ToUpper(baseCurrency)]
	baseDataList := baseData.([]interface{})
	quote := baseDataList[0].(map[string]interface{})
	foreignData := quote["quote"].(map[string]interface{})
	foreignDataQuote := foreignData[strings.ToUpper(foreignCurrency)]

	var cmccr CoinMarketCapConversionRateQuote
	err = mapstructure.Decode(foreignDataQuote, &cmccr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode conversion rate response")
	}

	return &types.ConversionRate{
		Base:       baseCurrency,
		Foreign:    foreignCurrency,
		Rate:       cmccr.Price,
		OracleName: c.name,
		Time:       time.Now().UnixMilli(),
	}, nil
}

func (c *CoinMarketCap) Name() string {
	return c.name
}

func (c *CoinMarketCap) IsEnabled() bool {
	return c.enable
}

func (c *CoinMarketCapResp) parseCoinMarketCapResp(body []byte) (CoinMarketCapResp, error) {
	err := json.Unmarshal(body, c)
	if err != nil {
		return CoinMarketCapResp{}, err
	}
	return *c, nil
}

func recoverErrorStringifier(err interface{}) error {
	switch e := err.(type) {
	case string:
		return errors.New(e)
	case runtime.Error:
		return e
	case error:
		return e
	default:
		return errors.New(e.(string))
	}
}
