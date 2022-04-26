// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package identity

type KeyType = string

const Secp256k1Type KeyType = "secp256k1"

type Keypair interface {
	// Type is the key algorithm type
	Type() string
	// Sign is used to sign the given message and return signature
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
