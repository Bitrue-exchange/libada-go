package libada

import (
	"encoding/hex"

	"github.com/fxamacker/cbor/v2"
	"golang.org/x/crypto/blake2b"
)

type Tx struct {
	_        struct{} `cbor:",toarray"`
	Body     *TransactionBody
	Witness  *Witness
	Metadata interface{} // TODO: supports it
}

// NewTx creates new Tx
func NewTx() *Tx {
	return &Tx{
		Body: &TransactionBody{
			Inputs:  make([]*Input, 0, 2),
			Outputs: make([]*Output, 0, 2),
			Fee:     100_0000,
		},
		Witness: &Witness{
			Keys:      make([]*KeysWitness, 0, 2),
			Bootstrap: make([]*BootstrapWitness, 0, 2),
		},
	}
}

// AddInputs add inputs to tx
func (t *Tx) AddInputs(i ...*Input) *Tx {
	t.Body.Inputs = append(t.Body.Inputs, i...)
	return t
}

// AddInputs add outputs to tx
func (t *Tx) AddOutputs(o ...*Output) *Tx {
	t.Body.Outputs = append(t.Body.Outputs, o...)
	return t
}

// SetInvalidBefore sets InvalidBefore for tx
func (t *Tx) SetInvalidBefore(v uint32) *Tx {
	t.Body.InvalidBefore = v
	return t
}

// SetInvalidAfter sets InvalidAfter for tx
func (t *Tx) SetInvalidAfter(v uint32) *Tx {
	t.Body.InvalidAfter = v
	return t
}

// SetFee sets fee for tx
func (t *Tx) SetFee(fee uint64) *Tx {
	t.Body.Fee = fee
	return t
}

// AddKeyWitness adds witness for Shelly address
func (t *Tx) AddKeyWitness(k *KeysWitness) *Tx {
	t.Witness.Keys = append(t.Witness.Keys, k)
	return t
}

// AddBootstrapWitness adds witness for legacy(Byron) address
func (t *Tx) AddBootstrapWitness(k *BootstrapWitness) *Tx {
	t.Witness.Bootstrap = append(t.Witness.Bootstrap, k)
	return t
}

// Bytes returns raw transaction bytes, it can submmit by cardano-graphql
func (t *Tx) Bytes() []byte {
	bytes, _ := cbor.Marshal(t)
	return bytes
}

// Hash returns message to sign
func (t *Tx) Hash() []byte {
	data, _ := cbor.Marshal(t.Body)
	hash := blake2b.Sum256(data)
	return hash[:]
}

// ID returns tx hash(AKA txid)
func (t *Tx) ID() string {
	return hex.EncodeToString(t.Hash())
}
