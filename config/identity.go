// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package config

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ChainSafe/sygma-fee-oracle/identity/secp256k1"
	"github.com/pkg/errors"
)

func LoadOracleIdentityKeyFromFile(keyPath string) ([]byte, error) {
	privBytes, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return nil, err
	}

	return privBytes, nil
}

func LoadOracleIdentityKeyFromEvn() (string, string) {
	evnKey := os.Getenv("IDENTITY_KEY")
	evnKeyType := os.Getenv("IDENTITY_KEY_TYPE")

	return evnKey, evnKeyType
}

func LoadOracleIdentityKey(keyPath, keyType string) (identity.Keypair, error) {
	var privBytes []byte
	var err error

	// load key from EVN params first, if not found, load from given key path
	evnKey, envKeyType := LoadOracleIdentityKeyFromEvn()
	if evnKey != "" && envKeyType != "" {
		keyType = envKeyType
		privBytes, err = hex.DecodeString(evnKey)
		if err != nil {
			return nil, err
		}
	} else {
		privBytes, err = LoadOracleIdentityKeyFromFile(keyPath)
		if err != nil {
			return nil, err
		}
	}

	switch strings.ToLower(keyType) {
	case identity.Secp256k1Type:
		return secp256k1.NewKeypairFromPrivateKey(privBytes)
	default:
		return nil, errors.New("unsupported key type")
	}
}
