// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/hex"
	"github.com/ChainSafe/sygma-fee-oracle/remoteParam"
	"github.com/pkg/errors"

	"io/ioutil"
	"os"
	"strconv"

	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/identity/secp256k1"

	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppEvm string
type AppMode string

var (
	AppEvmDev  AppEvm = "dev"
	AppEvmProd AppEvm = "production"

	AppModeRelease AppMode = "release"
	AppModeDebug   AppMode = "debug"

	remoteParamDomainData   = "/chainbridge/fee-oracle/domainData"
	remoteParamResourceData = "/chainbridge/fee-oracle/resourceData"
)

type Config struct {
	config configData
}

type configData struct {
	AppMode string `mapstructure:"app_mode"`

	Env string `mapstructure:"env"`

	LogLevel logrus.Level `mapstructure:"log_level"`

	HttpServer httpServerConfig `mapstructure:"http_server"`

	FinishUpTime int64 `mapstructure:"finish_up_time"`

	CronJob cronJobConfig `mapstructure:"cron_job"`

	Store store `mapstructure:"store"`

	Oracle oracle `mapstructure:"oracle"`

	GasPriceDomains []string `mapstructure:"gas_price_domains"`

	ConversionRatePairs []string `mapstructure:"conversion_rate_pairs"`

	Strategy strategyConfig `mapstructure:"strategy"`

	DataValidInterval int64 `mapstructure:"data_valid_interval"`

	Domains map[int]domain

	Resources map[string]resource
}

type strategyConfig struct {
	Local string `mapstructure:"local"`
}

type oracle struct {
	Etherscan     etherscan     `mapstructure:"etherscan"`
	Polygonscan   polygonscan   `mapstructure:"polygonscan"`
	CoinMarketCap coinMarketCap `mapstructure:"coinmarketcap"`
}

type store struct {
	Path string `mapstructure:"path"`
}

type cronJobConfig struct {
	UpdateConversionRateJob cronJob `mapstructure:"update_conversion_rate_job"`
	UpdateGasPriceJob       cronJob `mapstructure:"update_gas_price_job"`
}

type cronJob struct {
	Name           string `mapstructure:"name"`
	Enable         bool   `mapstructure:"enable"`
	CheckFrequency string `mapstructure:"check_frequency"`
	ProcessNumber  int    `mapstructure:"process_number"`
}

type httpServerConfig struct {
	Mode string `mapstructure:"mode"`
	Port string `mapstructure:"port"`
}

type polygonscan struct {
	Enable bool    `mapstructure:"enable"`
	ApiKey string  `mapstructure:"api_key"`
	Apis   apiUrls `mapstructure:"apis"`
}

type etherscan struct {
	Enable bool    `mapstructure:"enable"`
	ApiKey string  `mapstructure:"api_key"`
	Apis   apiUrls `mapstructure:"apis"`
}

type coinMarketCap struct {
	Enable bool    `mapstructure:"enable"`
	ApiKey string  `mapstructure:"api_key"`
	Apis   apiUrls `mapstructure:"apis"`
}

type apiUrls struct {
	GasPriceApiUrl string `mapstructure:"gas_price"`
	QueryRate      string `mapstructure:"query_rate"`
}

func (c *Config) logLevel() (logrus.Level, error) {
	logLvl := os.Getenv("LOG_LEVEL")
	if logLvl == "" {
		return c.config.LogLevel, nil
	}

	lvl, err := strconv.ParseUint(logLvl, 10, 64)
	if err != nil {
		return logrus.InfoLevel, err
	}

	return logrus.Level(lvl), nil
}

func (c *Config) WorkingEnvConfig() AppEvm {
	env := os.Getenv("WORKING_ENV")
	if env == "" {
		env = c.config.Env
	}

	switch env {
	case "dev":
		return AppEvmDev
	case "production":
		return AppEvmProd
	default:
		return AppEvmDev
	}
}

func (c *Config) AppModeConfig() AppMode {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = c.config.AppMode
	}

	switch mode {
	case "debug":
		return AppModeDebug
	case "release":
		return AppModeRelease
	default:
		return AppModeRelease
	}
}

func (c *Config) PrepareHttpServer() *gin.Engine {
	gin.SetMode(c.HttpServerConfig().Mode)
	instance := gin.New()

	return instance
}

func (c *Config) HttpServerConfig() httpServerConfig {
	httpConfig := c.config.HttpServer

	serverMode := os.Getenv("HTTP_SERVER_MODE")
	serverPort := os.Getenv("HTTP_SERVER_PORT")

	if serverMode != "" {
		httpConfig.Mode = serverMode
	}
	if serverPort != "" {
		httpConfig.Port = ":" + serverPort
	}

	return httpConfig
}

func (c *Config) FinishUpTimeConfig() int64 {
	return c.config.FinishUpTime
}

func (c *Config) OracleConfig() oracle {
	oracleConfig := c.config.Oracle

	etherscanAPIKey := os.Getenv("ETHERSCAN_API_KEY")
	polygonscanAPIKey := os.Getenv("POLYGONSCAN_API_KEY")
	coinMarketCapAPIKey := os.Getenv("COINMARKETCAP_API_KEY")

	if etherscanAPIKey != "" {
		oracleConfig.Etherscan.ApiKey = etherscanAPIKey
	}
	if polygonscanAPIKey != "" {
		oracleConfig.Polygonscan.ApiKey = polygonscanAPIKey
	}
	if coinMarketCapAPIKey != "" {
		oracleConfig.CoinMarketCap.ApiKey = coinMarketCapAPIKey
	}

	return oracleConfig
}

func (c *Config) CronJobConfig() cronJobConfig {
	cronjobConfig := c.config.CronJob

	conversionRateJobFrequency := os.Getenv("CONVERSION_RATE_JOB_FREQUENCY")
	gasPriceJobFrequency := os.Getenv("GAS_PRICE_JOB_FREQUENCY")

	if conversionRateJobFrequency != "" {
		cronjobConfig.UpdateConversionRateJob.CheckFrequency = conversionRateJobFrequency
	}
	if gasPriceJobFrequency != "" {
		cronjobConfig.UpdateGasPriceJob.CheckFrequency = gasPriceJobFrequency
	}

	return cronjobConfig
}

func (c *Config) StoreConfig() store {
	return c.config.Store
}

func (c *Config) GasPriceDomainsConfig() []string {
	return c.config.GasPriceDomains
}

func (c *Config) setDomains(domainData string) {
	c.config.Domains = parseDomains([]byte(domainData))
}

func (c *Config) GetRegisteredDomains(domainId int) *domain {
	d, ok := c.config.Domains[domainId]
	if !ok {
		return nil
	}
	return &d
}

func (c *Config) setResources(resourceData string) {
	c.config.Resources = parseResources([]byte(resourceData))
}

func (c *Config) GetRegisteredResources(resourceId string) *resource {
	r, ok := c.config.Resources[strings.ToLower(resourceId)]
	if !ok {
		return nil
	}
	return &r
}

func (c *Config) GetRegisteredResourceDomainInfo(resourceId string, domainId int) *resourceDomainInfo {
	r := c.GetRegisteredResources(resourceId)
	if r == nil {
		return nil
	}

	for _, d := range r.Domains {
		if d.DomainId == domainId {
			return &d
		}
	}
	return nil
}

func (c *Config) ConversionRatePairsChecker() error {
	if len(c.config.ConversionRatePairs)%2 != 0 {
		return errors.New("conversion_rate_pairs is invalid, must be pairs")
	}
	for _, e := range c.config.ConversionRatePairs {
		if e == "" {
			return errors.New("conversion_rate_pairs is invalid, element of pair is empty")
		}
	}

	return nil
}

func (c *Config) StrategyConfig() strategyConfig {
	return c.config.Strategy
}

func (c *Config) conversionRatePairsConfigLoad() []string {
	conversionRatePairs := os.Getenv("CONVERSION_RATE_PAIRS")
	if conversionRatePairs != "" {
		return strings.Split(conversionRatePairs, ",")
	}

	return c.config.ConversionRatePairs
}

func (c *Config) dataValidIntervalConfigLoad() int64 {
	dataValidInterval := os.Getenv("DATA_VALID_INTERVAL")
	if dataValidInterval != "" {
		i, err := strconv.ParseUint(dataValidInterval, 10, 64)
		if err != nil {
			panic(ErrLoadConfig.Wrap(errors.Wrap(err, "invalid DATA_VALID_INTERVAL from env param")))
		}
		return int64(i)
	}
	return c.config.DataValidInterval
}

func (c *Config) DataValidIntervalConfig() int64 {
	return c.config.DataValidInterval
}

func (c *Config) ConversionRatePairsConfig() [][]string {
	pricePairs := make([][]string, 0)

	for i := 0; i < len(c.config.ConversionRatePairs); i++ {
		m := make([]string, 0)
		m = append(m, c.config.ConversionRatePairs[i])
		m = append(m, c.config.ConversionRatePairs[i+1])
		pricePairs = append(pricePairs, m)
		i++
	}

	return pricePairs
}

func (c *Config) remoteParamsLoad(operator remoteParam.RemoteParamOperator, paramName string) (string, error) {
	out, err := operator.LoadParameter(paramName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to load given param: %s", paramName)
	}

	return out.Value, nil
}

// SetRemoteParams fetches the remote params and override the local ones
// only call this func when init app base
func (c *Config) SetRemoteParams(operator remoteParam.RemoteParamOperator) {
	if operator == nil {
		return
	}

	domains, err := c.remoteParamsLoad(operator, remoteParamDomainData)
	if err != nil {
		panic(err)
	}
	if domains != "" {
		c.setDomains(domains)
	}

	resources, err := c.remoteParamsLoad(operator, remoteParamResourceData)
	if err != nil {
		panic(err)
	}
	if resources != "" {
		c.setResources(resources)
	}
}

func (c *Config) EssentialConfigCheck() error {
	errs := make([]error, 0)

	err := c.ConversionRatePairsChecker()
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		return errs[0]
	}

	return nil
}

func LoadConfig(configPath, domainConfigPath, resourceConfigPath string) (*Config, *logrus.Logger) {
	conf, err := newConfig(configPath)
	if err != nil {
		panic(ErrLoadConfig.Wrap(err))
	}
	log := logrus.New()
	logLvl, err := conf.logLevel()
	if err != nil {
		panic(ErrLoadConfig.Wrap(err))
	}
	log.SetLevel(logLvl)

	// load domains and resources
	conf.config.Domains = loadDomains(domainConfigPath)
	conf.config.Resources = loadResources(resourceConfigPath)

	// load data valid interval
	conf.config.DataValidInterval = conf.dataValidIntervalConfigLoad()
	// load conversion price pair
	conf.config.ConversionRatePairs = conf.conversionRatePairsConfigLoad()

	return conf, log
}

func LoadOracleIdentityKeyFromFile(keyPath string) ([]byte, error) {
	privBytes, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return nil, err
	}

	return privBytes, nil
}

func LoadOracleIdentityKeyFromEvn() (string, string) {
	evnKey := os.Getenv("IDENTITY_KEY")
	evnKeyType := os.Getenv("IDENTITY_KEY_TYPE")

	return evnKey, evnKeyType
}

func LoadOracleIdentityKey(keyPath, keyType string) (identity.Keypair, error) {
	var privBytes []byte
	var err error

	// load key from EVN params first, if not found, load from given key path
	evnKey, envKeyType := LoadOracleIdentityKeyFromEvn()
	if evnKey != "" && envKeyType != "" {
		keyType = envKeyType
		privBytes, err = hex.DecodeString(evnKey)
		if err != nil {
			return nil, err
		}
	} else {
		privBytes, err = LoadOracleIdentityKeyFromFile(keyPath)
		if err != nil {
			return nil, err
		}
	}

	switch strings.ToLower(keyType) {
	case identity.Secp256k1Type:
		return secp256k1.NewKeypairFromPrivateKey(privBytes)
	default:
		return nil, errors.New("unsupported key type")
	}
}

func newConfig(configPath string) (*Config, error) {
	dir, file := path.Split(configPath)
	ext := filepath.Ext(file)[1:]
	fileName := strings.TrimSuffix(file, "."+ext)

	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType(ext)

	// where to look for
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	config := configData{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &Config{config: config}, nil
}
