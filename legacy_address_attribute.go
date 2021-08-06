package libada

import "github.com/fxamacker/cbor/v2"

type LegacyAddressAttribute struct {
	HDPayload []byte
	Network   *uint32
}

type attributeCborType struct {
	HDPayload []byte `cbor:"1,keyasint,omitempty"`
	Network   []byte `cbor:"2,keyasint,omitempty"`
}

func (a LegacyAddressAttribute) MarshalCBOR() ([]byte, error) {
	var t = &attributeCborType{HDPayload: a.HDPayload}
	if a.Network != nil {
		network, err := cbor.Marshal(a.Network)
		if err != nil {
			return nil, err
		}
		t.Network = network
	}
	return cbor.Marshal(t)
}

func (a *LegacyAddressAttribute) UnmarshalCBOR(data []byte) error {
	var t attributeCborType
	if err := cbor.Unmarshal(data, &t); err != nil {
		return err
	}

	var network *uint32
	if len(t.Network) != 0 {
		if err := cbor.Unmarshal(t.Network, &network); err != nil {
			return err
		}
	}

	*a = LegacyAddressAttribute{
		HDPayload: t.HDPayload,
		Network:   network,
	}
	return nil
}
