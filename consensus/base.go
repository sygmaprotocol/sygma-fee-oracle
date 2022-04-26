// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package consensus

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/consensus/strategy"
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/sirupsen/logrus"
)

// Consensus is the type defined for local(centralized) / group(decentralized) oracle to check and verify data in its store
// using the initialized strategy to preform the data checking/verifying logic
type Consensus struct {
	// strategy is used for oracle local store data check and verification and data aggregation
	strategy strategy.Strategy

	log *logrus.Entry
}

func NewConsensus(s strategy.Strategy, logger *logrus.Entry) *Consensus {
	return &Consensus{
		strategy: s,
		log:      logger.WithField("consensus", "base"),
	}
}

func (c *Consensus) GetStrategy() string {
	return c.strategy.Name()
}

func (c *Consensus) FilterLocalGasPriceData(store *store.GasPriceStore, DomainName string) (*types.GasPrices, error) {
	return c.strategy.GasPrice(store, DomainName)
}

func (c *Consensus) FilterLocalConversionRateData(store *store.ConversionRateStore, base, foreign string) (*types.ConversionRate, error) {
	return c.strategy.ConversionRate(store, base, foreign)
}
