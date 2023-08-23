// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package consensus

import (
	"github.com/ChainSafe/sygma-fee-oracle/consensus/strategy"
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/ChainSafe/sygma-fee-oracle/types"
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

func (c *Consensus) FilterLocalGasPriceData(store *store.GasPriceStore, domainID int) (*types.GasPrices, error) {
	return c.strategy.GasPrice(store, domainID)
}

func (c *Consensus) FilterLocalConversionRateData(store *store.ConversionRateStore, base, foreign string) (*types.ConversionRate, error) {
	return c.strategy.ConversionRate(store, base, foreign)
}
