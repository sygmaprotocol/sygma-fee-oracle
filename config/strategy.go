// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import "github.com/ChainSafe/chainbridge-fee-oracle/consensus/strategy"

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
