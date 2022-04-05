package util

import (
	"errors"
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
