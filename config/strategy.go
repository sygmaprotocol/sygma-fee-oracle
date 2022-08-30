// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import "github.com/ChainSafe/sygma-fee-oracle/consensus/strategy"

type RegisteredStrategy string

// all implemented strategies
const (
	Average RegisteredStrategy = "average"
)

func GetStrategy(s strategyConfig) strategy.Strategy {
	var localDataStrategy strategy.Strategy

	switch RegisteredStrategy(s.Local) {
	case Average:
		localDataStrategy = &strategy.Average{}
	default:
		panic("unsupported strategy")
	}

	return localDataStrategy
}
