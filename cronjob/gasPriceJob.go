// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package cronjob

import (
	"github.com/ChainSafe/sygma-fee-oracle/oracle"
	"github.com/pkg/errors"
)

func GasPriceJobOperation(c *Job) func() {
	return func() {
		c.log.Debug("checking gas price")

		// load gas price domains from config and run through each registered oracles
		for _, oracleOperator := range c.cronBase.gasPriceOracles {
			if !oracleOperator.IsOracleEnabled() {
				continue
			}
			for _, domain := range c.cronBase.base.GetConfig().GasPriceDomains {
				gasPriceData, err := oracleOperator.Run(domain)
				if err != nil || gasPriceData == nil {
					if err != oracle.ErrNotSupported {
						c.log.Error(errors.Wrapf(err, "failed to fetch data from oracle: %s", oracleOperator.GetOracleName()))
					}
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
}
