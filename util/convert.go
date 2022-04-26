// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package util

import (
	"errors"
	"math/big"
	"strings"
)

// Str2BigInt converts from string to big.Int
func Str2BigInt(small string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(small, 10)
	if !ok {
		return nil, errors.New("failed to convert from string to bigInt")
	}
	return n, nil
}

func Large2SmallUnitConverter(large string, decimal uint) *big.Int {
	curFloat, ok := zero().SetString(large)
	if !ok {
		return nil
	}
	exp, _ := zero().SetString("1" + strings.Repeat("0", int(decimal)) + ".0")
	weiFloat := zero().Mul(curFloat, exp)
	weiInt, _ := weiFloat.Int(nil)
	return weiInt
}

func zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(512)
	return r
}
