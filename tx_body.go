package libada

import (
	"github.com/fxamacker/cbor/v2"
)

type TransactionBody struct {
	Inputs        []*Input  `cbor:"0,keyasint"`
	Outputs       []*Output `cbor:"1,keyasint"`
	Fee           uint64    `cbor:"2,keyasint"`
	InvalidAfter  uint32    `cbor:"3,keyasint,omitempty"`
	InvalidBefore uint32    `cbor:"8,keyasint,omitempty"`

	// TODO: implements other fields
}

func (b *TransactionBody) Bytes() []byte {
	bytes, _ := cbor.Marshal(b)
	return bytes
}
