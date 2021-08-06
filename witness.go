package libada

import (
	"errors"

	"github.com/fxamacker/cbor/v2"
)

type Witness struct {
	Keys      []*KeysWitness      `cbor:"0,keyasint,omitempty"`
	Scripts   []*ScriptsWitness   `cbor:"1,keyasint,omitempty"`
	Bootstrap []*BootstrapWitness `cbor:"2,keyasint,omitempty"`
}

// KeysWithness the witness for Shelly address
type KeysWitness struct {
	_    struct{} `cbor:",toarray"`
	VKey []byte   // 32bytes pubkey
	Sign []byte   // ed25519 sign
}

func NewKeysWitness(vkey, sign []byte) *KeysWitness {
	if len(vkey) > 32 {
		vkey = vkey[:32]
	}
	return &KeysWitness{VKey: vkey, Sign: sign}
}

// BootstrapWitness the witness for Byron address
type BootstrapWitness struct {
	_         struct{} `cbor:",toarray"`
	VKey      []byte   // pubkey(32 bytes)
	Sign      []byte   // ed25519 signature
	Chaincode []byte   // chaincode
	Attribute []byte   // byron address attributes
}

func NewBootstrapWitness(addr *LegacyAddress, xpub, sig []byte) (*BootstrapWitness, error) {
	if len(xpub) != XPubSize {
		return nil, errors.New("xpub should 64 bytes")
	}

	rawAttribute, err := cbor.Marshal(addr.Attrs)
	if err != nil {
		return nil, err
	}

	result := &BootstrapWitness{
		VKey:      xpub[:32],
		Sign:      sig,
		Chaincode: xpub[32:],
		Attribute: rawAttribute,
	}
	return result, nil
}

type ScriptsWitness struct {
	_ struct{} `cbor:",toarray"`
}

func (s *ScriptsWitness) MarshalCBOR() ([]byte, error) {
	return nil, errors.New("unimplemented")
}
