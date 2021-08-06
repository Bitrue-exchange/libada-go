package libada

import "errors"

var ErrInvalidCredSize = errors.New("stake credential size is not 28")

type StakeCredentialType byte

const (
	KeyStakeCredentialType StakeCredentialType = iota
	ScriptStakeCredentialype
)

type StakeCredential struct {
	Kind StakeCredentialType `cbor:"0,keyasint"`
	Data []byte              `cbor:"1,keyasint"`
}

func NewKeysStakeCred(d []byte) StakeCredential {
	return StakeCredential{Kind: KeyStakeCredentialType, Data: d}
}
