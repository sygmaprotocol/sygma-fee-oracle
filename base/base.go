// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package base

import (
	"fmt"

	"github.com/ChainSafe/chainbridge-fee-oracle/config"
	"github.com/ChainSafe/chainbridge-fee-oracle/identity"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/ChainSafe/chainbridge-fee-oracle/store/db"
	"github.com/sirupsen/logrus"
)

type FeeOracleAppBase struct {
	log  *logrus.Entry
	conf *config.Config

	store          store.Store
	oracleIdentity identity.Keypair

	envInUse string
}

func NewFeeOracleAppBase(configPath, domainConfigPath, resourceConfigPath, keyPath, keyType string) *FeeOracleAppBase {
	conf, logger := config.LoadConfig(configPath, domainConfigPath, resourceConfigPath)

	base := &FeeOracleAppBase{
		log:      logger.WithField("base", "base"),
		conf:     conf,
		envInUse: conf.WorkingEnvConfig(),
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

func (a *FeeOracleAppBase) GetEnv() string {
	return a.envInUse
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

	a.oracleIdentity = kp
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
