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

func (c *Consensus) FilterLocalGasPriceData(store *store.GasPriceStore, domainId string) (*types.GasPricesResp, error) {
	return c.strategy.RunOnGasPrice(store, domainId)
}

func (c *Consensus) FilterLocalConversionRateData(store *store.ConversionRateStore, base, foreign string) (*types.ConversionRateResp, error) {
	return c.strategy.RunOnConversionRate(store, base, foreign)
}

func (c *Consensus) VoteGroupData() {
	// TODO: this would be used when fee oracle expended in decentralization
}
