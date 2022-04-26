// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package cronjob

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/oracle"
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
			for _, domain := range c.cronBase.base.GetConfig().GasPriceDomainsConfig() {
				gasPriceData, err := oracleOperator.Run(domain)
				if err != nil {
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
