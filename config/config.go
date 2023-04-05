// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"fmt"

	"github.com/pkg/errors"

	"os"
	"strconv"

	oracleErrors "github.com/ChainSafe/sygma-fee-oracle/errors"

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
)

type Config struct {
	logLevel logrus.Level `mapstructure:"log_level"`

	AppMode             string           `mapstructure:"app_mode"`
	Env                 string           `mapstructure:"env"`
	HttpServer          httpServerConfig `mapstructure:"http_server"`
	FinishUpTime        int64            `mapstructure:"finish_up_time"`
	CronJob             cronJobConfig    `mapstructure:"cron_job"`
	Store               store            `mapstructure:"store"`
	ConversionRateApis  []ApiService     `mapstructure:"conversion_rate_apis"`
	DomainsList         []domainList     `mapstructure:"domain_list"`
	ConversionRatePairs []string         `mapstructure:"conversion_rate_pairs"`
	Strategy            strategyConfig   `mapstructure:"strategy"`
	DataValidInterval   int64            `mapstructure:"data_valid_interval"`
	Domains             map[int]Domain
	Resources           map[string]*Resource
}

type strategyConfig struct {
	Local string `mapstructure:"local"`
}

type domainList struct {
	DomainID     int          `mapstructure:"domain_id"`
	GasPriceApis []ApiService `mapstructure:"gas_price_apis"`
}

type ApiService struct {
	Implementation string `mapstructure:"implementation"`
	Source         string `mapstructure:"source"`
	Enable         bool   `mapstructure:"enable"`
	URL            string `mapstructure:"url"`
	ApiKey         string `mapstructure:"api_key"`
	Decimals       int    `mapstructure:"decimals"`
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

func (c *Config) LogLevel() (logrus.Level, error) {
	logLvl := os.Getenv("LOG_LEVEL")
	if logLvl == "" {
		return c.logLevel, nil
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
		env = c.Env
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
		mode = c.AppMode
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
	httpConfig := c.HttpServer

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

// GasPriceApikeyReload dynamically builds the env var key string and replaces the api key if the env var is set
// dynamic env var key string format: <SOURCE>_API_KEY_<DOMAIN_ID>
func (c *Config) GasPriceApikeyReload(domainID int, apiService ApiService) ApiService {
	loadedApiKey := os.Getenv(fmt.Sprintf("%s_API_KEY_%d", strings.ToUpper(apiService.Source), domainID))
	if loadedApiKey != "" {
		apiService.ApiKey = loadedApiKey
	}

	return apiService
}

// ConversionRateApikeyReload dynamically builds the env var key string and replaces the api key if the env var is set
// dynamic env var key string format: <SOURCE>_API_KEY
func (c *Config) ConversionRateApikeyReload(apiService ApiService) ApiService {
	loadedApiKey := os.Getenv(fmt.Sprintf("%s_API_KEY", strings.ToUpper(apiService.Source)))
	if loadedApiKey != "" {
		apiService.ApiKey = loadedApiKey
	}

	return apiService
}

func (c *Config) CronJobConfig() cronJobConfig {
	cronjobConfig := c.CronJob

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

func (c *Config) Domain(domainId int) (*Domain, error) {
	d, ok := c.Domains[domainId]
	if !ok {
		return nil, fmt.Errorf("no domain registered")
	}
	return &d, nil
}

func (c *Config) Resource(resourceId string) (*Resource, error) {
	r, ok := c.Resources[strings.ToLower(resourceId)]
	if !ok {
		return nil, fmt.Errorf("no registered resource")
	}
	return r, nil
}

func (c *Config) ResourceDomainInfo(resourceId string, domainID int) (*DomainInfo, error) {
	r, err := c.Resource(resourceId)
	if err != nil {
		return nil, err
	}

	di, ok := r.DomainInfo[domainID]
	if !ok {
		return nil, fmt.Errorf("no domain info")
	}

	return di, nil
}

func (c *Config) ConversionRatePairsChecker() error {
	if len(c.ConversionRatePairs)%2 != 0 {
		return errors.New("conversion_rate_pairs is invalid, must be pairs")
	}
	for _, e := range c.ConversionRatePairs {
		if e == "" {
			return errors.New("conversion_rate_pairs is invalid, element of pair is empty")
		}
	}

	return nil
}

func (c *Config) conversionRatePairsConfigLoad() []string {
	conversionRatePairs := os.Getenv("CONVERSION_RATE_PAIRS")
	if conversionRatePairs != "" {
		return strings.Split(conversionRatePairs, ",")
	}

	return c.ConversionRatePairs
}

func (c *Config) dataValidIntervalConfigLoad() int64 {
	dataValidInterval := os.Getenv("DATA_VALID_INTERVAL")
	if dataValidInterval != "" {
		i, err := strconv.ParseUint(dataValidInterval, 10, 64)
		if err != nil {
			panic(oracleErrors.LoadConfig.Wrap(errors.Wrap(err, "invalid DATA_VALID_INTERVAL from env param")))
		}
		return int64(i)
	}
	return c.DataValidInterval
}

func (c *Config) ConversionRatePairsConfig() [][]string {
	pricePairs := make([][]string, 0)

	for i := 0; i < len(c.ConversionRatePairs); i++ {
		m := make([]string, 0)
		m = append(m, c.ConversionRatePairs[i])
		m = append(m, c.ConversionRatePairs[i+1])
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

func LoadConfig(configPath, domainConfigPath string) *Config {
	conf, err := loadConfigFromFile(configPath)
	if err != nil {
		panic(oracleErrors.LoadConfig.Wrap(err))
	}

	if strings.HasPrefix(domainConfigPath, "http") {
		conf.Domains, conf.Resources, err = loadDomainsFromNetwork(domainConfigPath)
		if err != nil {
			panic(oracleErrors.LoadConfig.Wrap(err))
		}
	} else {
		conf.Domains, conf.Resources, err = loadDomainsFromFile(domainConfigPath)
		if err != nil {
			panic(oracleErrors.LoadConfig.Wrap(err))
		}
	}

	conf.DataValidInterval = conf.dataValidIntervalConfigLoad()
	conf.ConversionRatePairs = conf.conversionRatePairsConfigLoad()

	return conf
}

func loadConfigFromFile(configPath string) (*Config, error) {
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

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
