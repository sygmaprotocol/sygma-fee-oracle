package base

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/ChainSafe/chainbridge-fee-oracle/store/db"
	"github.com/sirupsen/logrus"
)

type FeeOracleAppBase struct {
	log  *logrus.Entry
	conf *config.Config

	store store.Store

	envInUse string
}

func NewFeeOracleAppBase(configPath string) *FeeOracleAppBase {
	conf, logger := config.LoadConfig(configPath)

	base := &FeeOracleAppBase{
		log:      logger.WithField("base", "base"),
		conf:     conf,
		envInUse: conf.WorkingEnvConfig(),
	}

	base.initStore()

	base.verifyBase()
	return base
}

func (a *FeeOracleAppBase) GetConfig() *config.Config {
	return a.conf
}

func (a *FeeOracleAppBase) GetLogger() *logrus.Entry {
	return a.log
}

func (a *FeeOracleAppBase) GetEnv() string {
	return a.envInUse
}

func (a *FeeOracleAppBase) GetStore() store.Store {
	return a.store
}

func (a *FeeOracleAppBase) initStore() {
	storeDB, err := db.NewLvlDB(a.GetConfig().StoreConfig().Path)
	if err != nil {
		panic(err)
	}

	a.store = storeDB
}

// verifyBase preforms all essential checks for the app base
func (a *FeeOracleAppBase) verifyBase() {
	err := a.conf.EssentialConfigCheck()
	if err != nil {
		panic(err)
	}
}
