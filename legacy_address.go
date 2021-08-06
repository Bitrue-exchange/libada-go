package libada

import (
	"errors"
	"fmt"
	"hash/crc32"

	"github.com/fxamacker/cbor/v2"
	"github.com/islishude/base58"
)

const (
	XPubSize              = 64
	ByronPubKeyTag uint64 = 0
	ByronTag       uint64 = 24
)

type LegacyAddress struct {
	Hashed []byte
	Attrs  LegacyAddressAttribute
	Tag    uint64 // Byron only supports PubKeyTag(0)
}

func NewSimpleLegacyAddress(xpub []byte, networks ...uint32) (*LegacyAddress, error) {
	var attr LegacyAddressAttribute
	if len(networks) > 0 {
		attr.Network = &networks[0]
	}
	hashed, err := newHashedSpending(xpub, attr)
	if err != nil {
		return nil, err
	}
	return &LegacyAddress{Hashed: hashed, Attrs: attr, Tag: ByronPubKeyTag}, nil
}

func newHashedSpending(xpub []byte, attr LegacyAddressAttribute) ([]byte, error) {
	if len(xpub) != XPubSize {
		return nil, errors.New("xpub size should be 64 bytes")
	}
	spend, err := cbor.Marshal([]interface{}{ByronPubKeyTag, xpub})
	if err != nil {
		return nil, err
	}
	buf, err := cbor.Marshal([]interface{}{ByronPubKeyTag, cbor.RawMessage(spend), attr})
	if err != nil {
		return nil, err
	}
	return Sha3AndBlake2b224(buf), nil
}

func (a *LegacyAddress) MarshalCBOR() ([]byte, error) {
	if len(a.Hashed) != Hash28Size {
		return nil, errors.New("Invalid hash28 data")
	}
	raw, err := cbor.Marshal([]interface{}{a.Hashed, a.Attrs, a.Tag})
	if err != nil {
		return nil, err
	}
	return cbor.Marshal([]interface{}{
		cbor.Tag{Number: ByronTag, Content: raw},
		uint64(crc32.ChecksumIEEE(raw)),
	})
}

func (a *LegacyAddress) UnmarshalCBOR(data []byte) error {
	type RawAddress struct {
		_        struct{} `cbor:",toarray"`
		Tag      cbor.Tag
		Checksum uint64
	}

	var rawAddress RawAddress
	if err := cbor.Unmarshal(data, &rawAddress); err != nil {
		return fmt.Errorf("mashal raw: %s", err)
	}

	rawTag, ok := rawAddress.Tag.Content.([]byte)
	if !ok || rawAddress.Tag.Number != ByronTag {
		return errors.New("not a valid byron address")
	}

	checksum := crc32.ChecksumIEEE(rawTag)
	if rawAddress.Checksum != uint64(checksum) {
		return errors.New("checksum unmatched")
	}
	var got struct {
		_      struct{} `cbor:",toarray"`
		Hashed []byte
		Attrs  LegacyAddressAttribute
		Tag    uint64
	}
	if err := cbor.Unmarshal(rawTag, &got); err != nil {
		return err
	}
	if len(got.Hashed) != Hash28Size || got.Tag != ByronPubKeyTag {
		return errors.New("Invalid byron hashed or type")
	}
	if a == nil {
		return errors.New("unmarshal to nil value")
	}
	*a = LegacyAddress{Hashed: got.Hashed, Attrs: got.Attrs, Tag: got.Tag}
	return nil
}

func (a *LegacyAddress) Bytes() []byte {
	raw, _ := a.MarshalCBOR()
	return raw
}

func (a *LegacyAddress) Kind() AddressKind {
	return LegacyAddressKind
}

func (a *LegacyAddress) String() string {
	return base58.Encode(a.Bytes())
}

func (a *LegacyAddress) GetNetwork() Network {
	if a.Attrs.Network == nil {
		return Mainnet
	}
	return Testnet
}

func (a *LegacyAddress) Prefix() string {
	return ""
}
