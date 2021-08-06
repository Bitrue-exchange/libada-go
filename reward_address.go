package libada

import (
	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/fxamacker/cbor/v2"
)

type Reward struct {
	Network Network
	Payment StakeCredential
}

func (r *Reward) Bytes() []byte {
	buf := make([]byte, 29)
	buf[0] = 0b1110_0000 | (byte(r.Payment.Kind) << 4) | (byte(r.Network) & 0xf)
	copy(buf[1:], r.Payment.Data)
	return buf
}

func (r *Reward) String() string {
	res, _ := bech32.Encode(r.Prefix(), r.Bytes())
	return res
}

func (r *Reward) Kind() AddressKind {
	return RewardAddressKind
}

func (r *Reward) MarshalCBOR() ([]byte, error) {
	if len(r.Payment.Data) != Hash28Size {
		return nil, ErrInvalidCredSize
	}
	return cbor.Marshal(r.Bytes())
}

func (r *Reward) GetNetwork() Network {
	return r.Network
}

func (r *Reward) Prefix() string {
	if r.Network == Testnet {
		return StakePrefix + TestnetSuffix
	}
	return StakePrefix
}
