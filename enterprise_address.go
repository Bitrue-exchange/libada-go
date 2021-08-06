package libada

import (
	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/fxamacker/cbor/v2"
)

type EnterpriseAddress struct {
	Network Network
	Payment StakeCredential
}

// NewEnterpriseAddress creates Enterprise type address with public key
// the pubkey param can be pubkeyHash or 32 bytes pubkey or 64 bytes xpub
func NewEnterpriseAddress(pubkey []byte, kind StakeCredentialType, network Network) *EnterpriseAddress {
	//  KeyStakeCred.Data is Blake2b224 result of public key 32bytes
	if len(pubkey) >= 32 {
		pubkey = Blake2b224(pubkey[:32])
	}
	return &EnterpriseAddress{
		Network: network,
		Payment: StakeCredential{
			Kind: kind,
			Data: pubkey,
		},
	}
}

func NewKeyedEnterpriseAddress(key []byte, network Network) *EnterpriseAddress {
	return NewEnterpriseAddress(key, KeyStakeCredentialType, network)
}

func (e *EnterpriseAddress) Bytes() []byte {
	buf := make([]byte, 29)
	buf[0] = 0b0110_0000 | (byte(e.Payment.Kind) << 4) | (byte(e.Network) & 0xf)
	copy(buf[1:], e.Payment.Data)
	return buf
}

func (e *EnterpriseAddress) String() string {
	res, _ := bech32.Encode(e.Prefix(), e.Bytes())
	return res
}

func (e *EnterpriseAddress) Kind() AddressKind {
	return EnterpriseAddressKind
}

func (e *EnterpriseAddress) MarshalCBOR() ([]byte, error) {
	if len(e.Payment.Data) != Hash28Size {
		return nil, ErrInvalidCredSize
	}
	return cbor.Marshal(e.Bytes())
}

func (e *EnterpriseAddress) GetNetwork() Network {
	return e.Network
}

func (e *EnterpriseAddress) Prefix() string {
	if e.Network == Testnet {
		return B32Prefix + TestnetSuffix
	}
	return B32Prefix
}
