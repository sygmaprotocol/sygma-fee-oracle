// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package secp256k1

import (
	"crypto/ecdsa"

	"github.com/ChainSafe/sygma-fee-oracle/identity"

	"github.com/ethereum/go-ethereum/common/hexutil"
	secp256k1 "github.com/ethereum/go-ethereum/crypto"
)

var _ identity.Keypair = (*KeyPair)(nil)

type KeyPair struct {
	public  *ecdsa.PublicKey
	private *ecdsa.PrivateKey
}

func NewKeypairFromPrivateKey(priv []byte) (identity.Keypair, error) {
	pk, err := secp256k1.ToECDSA(priv)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		public:  pk.Public().(*ecdsa.PublicKey),
		private: pk,
	}, nil
}

func (e *KeyPair) Type() string {
	return "secp256k1"
}

func (e *KeyPair) Address() string {
	return secp256k1.PubkeyToAddress(*e.public).String()
}

func (e *KeyPair) PublicKey() string {
	return hexutil.Encode(secp256k1.CompressPubkey(e.public))
}

func (e *KeyPair) Sign(data []byte) ([]byte, error) {
	return secp256k1.Sign(secp256k1.Keccak256(data), e.private)
}

func (e *KeyPair) Encode() []byte {
	return secp256k1.FromECDSA(e.private)
}

func (e *KeyPair) Decode(bytes []byte) error {
	key, err := secp256k1.ToECDSA(bytes)
	if err != nil {
		return err
	}

	e.public = key.Public().(*ecdsa.PublicKey)
	e.private = key

	return nil
}
