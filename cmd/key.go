// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ChainSafe/sygma-fee-oracle/identity/secp256k1"
	"github.com/pkg/errors"

	"github.com/ChainSafe/sygma-fee-oracle/identity"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

const (
	defaultKeyFilePath = "keyfile.priv"
)

var (
	keyGenerator = map[identity.KeyType]keyGenerationFunc{
		identity.Secp256k1Type: secp256k1KeyGenerator,
	}
)

type keyGenerationFunc func(string) string

type oracleKeyGeneration struct {
	KeyFilePath string
	KeyType     string
}

var oracleKeyGenerationCmd = &cobra.Command{
	Use:   "key-generate",
	Short: "Start Sygma fee oracle identity key generation",
	RunE:  generateKey,
}

var osKey oracleKeyGeneration

func init() {
	oracleKeyGenerationCmd.Flags().StringVarP(&osKey.KeyFilePath, "keyfile_path", "k", "keyfile.priv", "Output dir for generated key file, filename is required with .priv as file extension")
	oracleKeyGenerationCmd.Flags().StringVarP(&osKey.KeyType, "key_type", "t", "secp256k1", "Support: secp256k1")
}

func generateKey(_ *cobra.Command, _ []string) error {
	fmt.Printf("key file path: %s\n", osKey.KeyFilePath)
	fmt.Printf("key type: %s\n", osKey.KeyType)

	if filepath.Clean(osKey.KeyFilePath) == defaultKeyFilePath {
		_, err := os.Open(defaultKeyFilePath)
		if err == nil {
			fmt.Println("Warning: default keyfile.priv exists already")
			return nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}

	generationFunc, ok := keyGenerator[strings.ToLower(osKey.KeyType)]
	if !ok {
		panic("invalid key type")
	}
	address := generationFunc(osKey.KeyFilePath)

	fmt.Printf("address of the generated key is: %s\n", address)

	return nil
}

func secp256k1KeyGenerator(keyPath string) string {
	priv, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(keyPath, crypto.FromECDSA(priv), 0600)
	if err != nil {
		panic(err)
	}

	kp, err := secp256k1.NewKeypairFromPrivateKey(crypto.FromECDSA(priv))
	if err != nil {
		panic(err)
	}

	return kp.Address()
}
