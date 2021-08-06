package libada

import (
	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/fxamacker/cbor/v2"
)

type BaseAddress struct {
	Network Network
	Payment StakeCredential
	Stake   StakeCredential
}

func (b *BaseAddress) Bytes() []byte {
	buf := make([]byte, 57)
	buf[0] = (byte(b.Payment.Kind) << 4) | (byte(b.Stake.Kind) << 5) | (byte(b.Network) & 0xf)
	copy(buf[1:29], b.Payment.Data)
	copy(buf[29:], b.Stake.Data)
	return buf
}

func (b *BaseAddress) String() string {
	res, _ := bech32.Encode(b.Prefix(), b.Bytes())
	return res
}

func (b *BaseAddress) Kind() AddressKind {
	return BaseAddressKind
}

func (b *BaseAddress) MarshalCBOR() ([]byte, error) {
	if len(b.Payment.Data) != Hash28Size || len(b.Stake.Data) != Hash28Size {
		return nil, ErrInvalidCredSize
	}
	return cbor.Marshal(b.Bytes())
}

func (b *BaseAddress) GetNetwork() Network {
	return b.Network
}

func (b *BaseAddress) Prefix() string {
	if b.Network == Testnet {
		return B32Prefix + TestnetSuffix
	}
	return B32Prefix
}
