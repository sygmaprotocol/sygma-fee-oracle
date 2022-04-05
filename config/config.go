package config

import (
	"errors"
	"github.com/ChainSafe/chainbridge-fee-oracle/identity"
	"github.com/ChainSafe/chainbridge-fee-oracle/identity/secp256k1"
	"io/ioutil"

	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	config configData
}

type configData struct {
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

	Domains map[int]domain

	Resources map[int]resource
}

type strategyConfig struct {
	Local string `mapstructure:"local"`
}

type oracle struct {
	Etherscan     etherscan     `mapstructure:"etherscan"`
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

type etherscan struct {
	Enable bool   `mapstructure:"enable"`
	ApiKey string `mapstructure:"api_key"`
}

type coinMarketCap struct {
	Enable bool   `mapstructure:"enable"`
	ApiKey string `mapstructure:"api_key"`
}

func (c *Config) logLevel() logrus.Level {
	return c.config.LogLevel
}

func (c *Config) WorkingEnvConfig() string {
	return c.config.Env
}

func (c *Config) PrepareHttpServer() *gin.Engine {
	gin.SetMode(c.config.HttpServer.Mode)
	instance := gin.New()

	return instance
}

func (c *Config) HttpServerConfig() httpServerConfig {
	return c.config.HttpServer
}

func (c *Config) FinishUpTimeConfig() int64 {
	return c.config.FinishUpTime
}

func (c *Config) OracleConfig() oracle {
	return c.config.Oracle
}

func (c *Config) CronJobConfig() cronJobConfig {
	return c.config.CronJob
}

func (c *Config) StoreConfig() store {
	return c.config.Store
}

func (c *Config) GasPriceDomainsConfig() []string {
	return c.config.GasPriceDomains
}

func (c *Config) GetRegisteredDomains(domainId int) *domain {
	d, ok := c.config.Domains[domainId]
	if !ok {
		return nil
	}
	return &d
}

func (c *Config) GetRegisteredResources(resourceId int) *resource {
	r, ok := c.config.Resources[resourceId]
	if !ok {
		return nil
	}
	return &r
}

func (c *Config) ConversionRatePairsChecker() error {
	if len(c.config.ConversionRatePairs)%2 != 0 {
		return errors.New("conversion_rate_pairs is invalid, must be pairs")
	}
	return nil
}

func (c *Config) StrategyConfig() strategyConfig {
	return c.config.Strategy
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

func LoadConfig(configPath string) (*Config, *logrus.Logger) {
	conf, err := newConfig(configPath)
	if err != nil {
		panic(ErrLoadConfig.Wrap(err))
	}
	log := logrus.New()
	log.SetLevel(conf.logLevel())

	// load domains and resources
	conf.config.Domains = loadDomains()
	conf.config.Resources = loadResources(conf.config.Domains)

	return conf, log
}

func LoadOracleIdentityKey(keyPath, keyType string) (identity.Keypair, error) {
	privBytes, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return nil, err
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
