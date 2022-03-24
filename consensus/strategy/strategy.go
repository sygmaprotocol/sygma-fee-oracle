package strategy

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
)

type Strategy interface {
	Name() string
	RunOnGasPrice(*store.GasPriceStore, string) (*types.GasPricesResp, error)
	RunOnConversionRate(*store.ConversionRateStore, string, string) (*types.ConversionRateResp, error)
}
