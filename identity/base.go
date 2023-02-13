// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package identity

type OracleIdentityOperator struct {
	identity Keypair
}

func NewOracleIdentityOperator(key Keypair) *OracleIdentityOperator {
	return &OracleIdentityOperator{
		identity: key,
	}
}

func (o *OracleIdentityOperator) Sign(data []byte) ([]byte, error) {
	return o.identity.Sign(data)
}

func (o *OracleIdentityOperator) IdentityAddress() string {
	return o.identity.Address()
}
