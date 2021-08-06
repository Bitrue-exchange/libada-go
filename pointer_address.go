package libada

import (
	"bytes"

	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/fxamacker/cbor/v2"
)

type StakePoint struct {
	Slot      uint64
	TxIndex   uint64
	CertIndex uint64
}

type PointerAddress struct {
	Network Network
	Payment StakeCredential
	Stake   StakePoint
}

func VariableNatEncode(n uint64) []byte {
	o := []byte{byte(n) & 0x7f}
	n /= 128
	for n > 0 {
		o = append(o, byte(n&0x7f)|0x80)
		n /= 128
	}
	for i, j := 0, len(o)-1; i < j; i, j = i+1, j-1 {
		o[i], o[j] = o[j], o[i]
	}
	return o
}

func VariableNatDecode(bytes []byte) (uint64, int, bool) {
	var output uint64
	var bytes_read int

	for _, b := range bytes {
		output = (output << 7) | uint64(b&0x7f)
		bytes_read += 1
		if (b & 0x80) == 0 {
			return output, bytes_read, true
		}
	}
	return 0, 0, false
}

func (p *PointerAddress) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	_ = buf.WriteByte(0b0100_0000 | (byte(p.Payment.Kind) << 4) | (byte(p.Network) & 0xf))
	_, _ = buf.Write(p.Payment.Data)

	_, _ = buf.Write(VariableNatEncode(p.Stake.Slot))
	_, _ = buf.Write(VariableNatEncode(p.Stake.TxIndex))
	_, _ = buf.Write(VariableNatEncode(p.Stake.CertIndex))
	return buf.Bytes()
}

func (p *PointerAddress) String() string {
	res, _ := bech32.Encode(p.Prefix(), p.Bytes())
	return res
}

func (p *PointerAddress) Kind() AddressKind {
	return PointerAddressKind
}

func (p *PointerAddress) MarshalCBOR() ([]byte, error) {
	if len(p.Payment.Data) != Hash28Size {
		return nil, ErrInvalidCredSize
	}
	return cbor.Marshal(p.Bytes())
}

func (p *PointerAddress) GetNetwork() Network {
	return p.Network
}

func (p *PointerAddress) Prefix() string {
	if p.Network == Testnet {
		return B32Prefix + TestnetSuffix
	}
	return B32Prefix
}
