// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package identity

type KeyType = string

const Secp256k1Type KeyType = "secp256k1"

type Keypair interface {
	// Type is the key algorithm type
	Type() string
	// Sign is used to sign the given unhashed message and return signature
	Sign([]byte) ([]byte, error)
	// Encode is used to write the key to a file
	Encode() []byte
	// Decode is used to retrieve a key from a file
	Decode([]byte) error
	// Address provides the address for the keypair
	Address() string
	// PublicKey returns the keypair's public key an encoded a string
	PublicKey() string
}
