// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package strategy

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
)

type Strategy interface {
	Name() string
	GasPrice(*store.GasPriceStore, string) (*types.GasPrices, error)
	ConversionRate(*store.ConversionRateStore, string, string) (*types.ConversionRate, error)
}
