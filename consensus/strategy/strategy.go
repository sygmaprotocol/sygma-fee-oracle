// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package strategy

import (
	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/ChainSafe/sygma-fee-oracle/types"
)

type Strategy interface {
	Name() string
	GasPrice(*store.GasPriceStore, int) (*types.GasPrices, error)
	ConversionRate(*store.ConversionRateStore, string, string) (*types.ConversionRate, error)
}
