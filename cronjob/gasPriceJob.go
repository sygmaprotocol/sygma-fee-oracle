// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package cronjob

import (
	"github.com/pkg/errors"
)

func GasPriceJobOperation(c *Job) func() {
	return func() {
		c.log.Debug("checking gas price")

		// run through each registered oracles and fetch gas price data for its associated supported domain
		for _, oracleOperator := range c.cronBase.gasPriceOracles {
			if !oracleOperator.IsOracleEnabled() {
				continue
			}

			gasPriceData, err := oracleOperator.Run()
			if err != nil || gasPriceData == nil {
				c.log.Error(errors.Wrapf(err, "failed to fetch data from oracle: %s", oracleOperator.GetOracleSource()))
				continue
			}

			c.log.Debugf("gas price data: %+v\n", gasPriceData)

			err = c.cronBase.gasPriceStore.StoreGasPrice(gasPriceData)
			if err != nil {
				c.log.Error(errors.Wrap(err, "failed to store data into store"))
				continue
			}
		}
	}
}
