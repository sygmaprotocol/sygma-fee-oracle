package strategy

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ChainSafe/chainbridge-fee-oracle/store"
	"github.com/ChainSafe/chainbridge-fee-oracle/types"
)

var _ Strategy = (*Average)(nil)

type Average struct{}

func (a *Average) Name() string {
	return "average"
}

func (a *Average) RunOnGasPrice(store *store.GasPriceStore, domainId string) (*types.GasPricesResp, error) {
	re, err := store.GetGasPriceByDomain(domainId)
	if err != nil {
		return nil, err
	}
	if len(re) == 0 {
		return nil, errors.New("no data found")
	}

	fast, propose, safe := float64(0), float64(0), float64(0)
	dataSize := float64(len(re))
	for _, resp := range re {
		f, err := strconv.ParseFloat(resp.FastGasPrice, 64)
		if err != nil {
			return nil, err
		}
		fast += f
		p, err := strconv.ParseFloat(resp.ProposeGasPrice, 64)
		if err != nil {
			return nil, err
		}
		propose += p
		s, err := strconv.ParseFloat(resp.SafeGasPrice, 64)
		if err != nil {
			return nil, err
		}
		safe += s
	}

	return &types.GasPricesResp{
		SafeGasPrice:    fmt.Sprintf("%d", int(math.Round(fast/dataSize))),
		ProposeGasPrice: fmt.Sprintf("%d", int(propose/dataSize)),
		FastGasPrice:    fmt.Sprintf("%d", int(safe/dataSize)),
		DomainId:        domainId,
		Time:            re[0].Time, // use the first data time for now
	}, nil

}

func (a *Average) RunOnConversionRate(store *store.ConversionRateStore, base, foreign string) (*types.ConversionRateResp, error) {
	re, err := store.GetConversionRateByCurrencyPair(base, foreign)
	if err != nil {
		return nil, err
	}
	if len(re) == 0 {
		return nil, errors.New("no data found")
	}

	rate := float64(0)
	dataSize := float64(len(re))
	for _, resp := range re {
		rate += resp.Rate
	}

	return &types.ConversionRateResp{
		Base:    base,
		Foreign: foreign,
		Rate:    rate / dataSize,
		Time:    re[0].Time,
	}, nil
}
