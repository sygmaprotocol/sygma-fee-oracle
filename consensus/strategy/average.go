// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package strategy

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/ChainSafe/sygma-fee-oracle/store"
	"github.com/ChainSafe/sygma-fee-oracle/types"
)

var _ Strategy = (*Average)(nil)

type Average struct{}

func (a *Average) Name() string {
	return "average"
}

func (a *Average) GasPrice(store *store.GasPriceStore, domainID int) (*types.GasPrices, error) {
	re, err := store.GetGasPricesByDomain(domainID)
	if err != nil {
		return nil, err
	}
	if len(re) == 0 {
		return nil, errors.New("no gas price data found")
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

	return &types.GasPrices{
		SafeGasPrice:    fmt.Sprintf("%d", int(math.Round(safe/dataSize))),
		ProposeGasPrice: fmt.Sprintf("%d", int(propose/dataSize)),
		FastGasPrice:    fmt.Sprintf("%d", int(fast/dataSize)),
		DomainID:        domainID,
		Time:            re[0].Time, // use the first data time for now
	}, nil

}

func (a *Average) ConversionRate(store *store.ConversionRateStore, base, foreign string) (*types.ConversionRate, error) {
	if base == foreign {
		return &types.ConversionRate{
			Base:    base,
			Foreign: foreign,
			Rate:    1,
			Time:    time.Now().Unix(),
		}, nil
	}

	re, err := store.GetConversionRatesByCurrencyPair(base, foreign)
	if err != nil {
		return nil, err
	}
	if len(re) == 0 {
		return nil, errors.New("no conversion rate data found")
	}

	rate := float64(0)
	dataSize := float64(len(re))
	for _, resp := range re {
		rate += resp.Rate
	}

	return &types.ConversionRate{
		Base:    base,
		Foreign: foreign,
		Rate:    rate / dataSize,
		Time:    re[0].Time,
	}, nil
}
