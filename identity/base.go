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

func (o *OracleIdentityOperator) GetOracleIdentityAlgoName() string {
	return o.identity.Type()
}

func (o *OracleIdentityOperator) GetOracleIdentityAddress() string {
	return o.identity.Address()
}
