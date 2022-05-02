// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package util

import (
	"bytes"
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

func Large2SmallUnitConverter(large string, decimal uint) (*big.Int, error) {
	curFloat, ok := zero().SetString(large)
	if !ok {
		return big.NewInt(0), errors.New("failed to convert on the given decimal")
	}
	exp, _ := zero().SetString("1" + strings.Repeat("0", int(decimal)) + ".0")
	weiFloat := zero().Mul(curFloat, exp)
	weiInt, _ := weiFloat.Int(nil)
	return weiInt, nil
}

func zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(512)
	return r
}

func PaddingZero(data []byte, length int) []byte {
	b := bytes.Buffer{}

	for i := 0; i < length-len(data); i++ {
		b.WriteByte(0)
	}
	b.Write(data)

	return b.Bytes()
}

// Byte32Converter converts the first 32 bytes of given []byte to [32]byte
func Byte32Converter(b []byte) [32]byte {
	f := [32]byte{}
	copy(f[:], b)
	return f
}
