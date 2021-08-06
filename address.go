package libada

import (
	"errors"
	"strings"

	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/fxamacker/cbor/v2"
	"github.com/islishude/base58"
)

const (
	B32Prefix     = "addr"
	TestnetSuffix = "_test"
	StakePrefix   = "stake"
	Hash28Size    = 28
)

//go:generate stringer -type=AddressKind
type AddressKind byte

const (
	LegacyAddressKind AddressKind = iota
	BaseAddressKind
	PointerAddressKind
	EnterpriseAddressKind
	RewardAddressKind
)

type Address interface {
	cbor.Marshaler

	// Raw bytes for transaction outputs
	Bytes() []byte
	// Readable and humanable string
	String() string
	// Bech32 prefix, and empty for legacy address
	Prefix() string
	// address kind
	Kind() AddressKind
	// address network
	GetNetwork() Network
}

func MustDecodeAddress(raw string) Address {
	res, err := DecodeAddress(raw)
	if err != nil {
		panic(err)
	}
	return res
}

// DecodeAddress decode humanable address
// it should be invalid if you get an error
// for an exchange, you should also check address network and address type
func DecodeAddress(raw string) (res Address, err error) {
	var rbytes []byte
	var prefix string

	if strings.HasPrefix(raw, B32Prefix) || strings.HasPrefix(raw, StakePrefix) {
		prefix, rbytes, err = bech32.Decode(raw)
	} else {
		rbytes, err = base58.Decode(raw)
	}

	if err != nil {
		return
	}

	res, err = DecodeRawAddress(rbytes)
	if err != nil {
		return
	}

	if p := res.Prefix(); p != prefix {
		err = errors.New("invalid address prefix")
	}

	return
}

// DecodeRawAddresss decodes raw address bytes
// it decode raw address in tx output, you may don't need it
func DecodeRawAddress(s []byte) (Address, error) {
	if len(s) == 0 {
		return nil, errors.New("empty address")
	}

	header := s[0]
	network := Network(header & 0x0f)

	readAddrCred := func(bit byte, pos int) StakeCredential {
		hashBytes := s[pos : pos+Hash28Size]
		if header&(1<<bit) == 0 {
			return StakeCredential{Kind: KeyStakeCredentialType, Data: hashBytes}
		}
		return StakeCredential{Kind: ScriptStakeCredentialype, Data: hashBytes}
	}

	switch (header & 0xf0) >> 4 {
	// Base type
	case 0b0000, 0b0001, 0b0010, 0b0011:
		// header + keyhash
		if len(s) != 57 {
			return nil, errors.New("Invalid length for base address")
		}
		return &BaseAddress{Network: network, Payment: readAddrCred(4, 1),
			Stake: readAddrCred(5, Hash28Size+1)}, nil
	// Pointer type
	case 0b0100, 0b0101:
		// header + keyhash + 3 natural numbers (min 1 byte each)
		if len(s) < 32 {
			return nil, errors.New("Invalid length for pointer address")
		}
		byteIndex := 1
		paymentCred := readAddrCred(4, 1)
		slot, slotBytes, ok := VariableNatDecode(s[byteIndex:])
		if !ok {
			return nil, errors.New("slot variable decode failed")
		}
		byteIndex += slotBytes

		txIndex, txBytes, ok := VariableNatDecode(s[byteIndex:])
		if !ok {
			return nil, errors.New("txIndex variable decode failed")
		}
		byteIndex += txBytes

		certIndex, certBytes, ok := VariableNatDecode(s[byteIndex:])
		if !ok {
			return nil, errors.New("certIndex variable decode failed")
		}
		byteIndex += certBytes

		if byteIndex > len(s) {
			return nil, errors.New("byte index is out range of pointer lenght")
		}

		return &PointerAddress{
			Network: network, Payment: paymentCred,
			Stake: StakePoint{Slot: slot, TxIndex: txIndex, CertIndex: certIndex},
		}, nil
	// Enterprise type
	case 0b0110, 0b0111:
		// header + keyhash
		if len(s) != 29 {
			return nil, errors.New("invalid length for enterprise address")
		}
		return &EnterpriseAddress{Network: network, Payment: readAddrCred(4, 1)}, nil
	// Reward type
	case 0b1110, 0b1111:
		if len(s) != 29 {
			return nil, errors.New("invalid length for reward address")
		}
		return &Reward{Network: network, Payment: readAddrCred(4, 1)}, nil
	// Legacy byron type
	case 0b1000:
		var byron LegacyAddress
		if err := cbor.Unmarshal(s, &byron); err != nil {
			return nil, err
		}
		return &byron, nil
	}
	return nil, errors.New("unsupports address type")
}
