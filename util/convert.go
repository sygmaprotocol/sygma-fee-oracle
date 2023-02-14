// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package util

import (
	"bytes"
	"errors"
	"fmt"
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
	a := strings.Split(large, ".")
	for i := 0; i < int(decimal); i++ {
		if len(a) > 1 && len(a[1]) > i {
			a[0] += string(a[1][i])
		} else {
			a[0] += "0"
		}
	}

	b, ok := new(big.Int).SetString(a[0], 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert string to big int")
	}

	return b, nil
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
