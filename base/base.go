// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package base

import (
	"fmt"

	"github.com/ChainSafe/sygma-fee-oracle/config"
	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/ChainSafe/sygma-fee-oracle/store/db"
	"github.com/sirupsen/logrus"
)

type FeeOracleAppBase struct {
	log  *logrus.Entry
	conf *config.Config

	store          store.Store
	oracleIdentity identity.Keypair

	env config.AppEvm
}

func NewFeeOracleAppBase(configPath, domainConfigPath, keyPath, keyType string) *FeeOracleAppBase {
	conf := config.LoadConfig(configPath, domainConfigPath)

	logger := logrus.New()
	logLvl, err := conf.LogLevel()
	if err != nil {
		panic("invalid log level")
	}
	logger.SetLevel(logLvl)
	logger.Infof("log level: %s", logger.Level)

	base := &FeeOracleAppBase{
		log:  logger.WithField("base", "base"),
		conf: conf,
		env:  conf.WorkingEnvConfig(),
	}

	base.initKeyPair(keyPath, keyType)
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

func (a *FeeOracleAppBase) GetEnv() config.AppEvm {
	return a.env
}

func (a *FeeOracleAppBase) GetStore() store.Store {
	return a.store
}

func (a *FeeOracleAppBase) GetOracleIdentity() identity.Keypair {
	return a.oracleIdentity
}

func (a *FeeOracleAppBase) initKeyPair(keyPath, keyType string) {
	kp, err := config.LoadOracleIdentityKey(keyPath, keyType)
	if err != nil {
		panic(fmt.Sprintf("failed to load key config: %s", err))
	}

	a.log.Infof("fee oracle indentity address: %s\n", kp.Address())

	a.oracleIdentity = kp
}

func (a *FeeOracleAppBase) initStore() {
	storeDB, err := db.NewLvlDB(a.GetConfig().Store.Path)
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
