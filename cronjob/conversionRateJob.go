package cronjob

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
	"github.com/pkg/errors"
)

func ConversionRateJobOperation(c *Job) func() {
	return func() {
		c.log.Debug("checking conversion rate")

		pricePairs := c.cronBase.base.GetConfig().ConversionRatePairsConfig()

		// load conversion pairs from config and run through each registered oracles
		for _, oracleOperator := range c.cronBase.conversionRateOracles {
			if !oracleOperator.IsOracleEnabled() {
				continue
			}
			for _, pricePair := range pricePairs {
				rateData, err := oracleOperator.Run(pricePair[0], pricePair[1])
				if err != nil {
					c.log.Error(errors.Wrapf(err, "failed to fetch data from oracle: %s", oracleOperator.GetOracleName()))
					continue
				}

				c.log.Debugf("conversion rate data: %+v\n", rateData)

				err = c.cronBase.conversionRateStore.StoreConversionRate(rateData)
				if err != nil {
					c.log.Error(errors.Wrap(err, "failed to store data into store"))
					continue
				}

				reverseRateData := &types.ConversionRate{
					Base:       pricePair[1],
					Foreign:    pricePair[0],
					Rate:       1 / rateData.Rate,
					OracleName: rateData.OracleName,
					Time:       rateData.Time,
				}

				err = c.cronBase.conversionRateStore.StoreConversionRate(reverseRateData)
				if err != nil {
					c.log.Error(errors.Wrap(err, "failed to store data into store"))
					continue
				}
			}
		}
	}
}
