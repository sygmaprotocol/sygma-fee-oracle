// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package util

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
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
	bf, ok := new(big.Float).SetString(large)
	if !ok {
		return nil, fmt.Errorf("failed conversion from string to big float")
	}
	exp := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(decimal)), nil)
	fexp, _ := new(big.Float).SetString(exp.String())
	bf = bf.Mul(bf, fexp)
	bi, _ := bf.Int(nil)
	return bi, nil
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
