package libada

import (
	"encoding/hex"
	"errors"
)

type Input struct {
	_      struct{} `cbor:",toarray"`
	TxHash []byte   // blake2b256
	Index  uint32
}

func NewInput(txid string, index uint32) (*Input, error) {
	hash, err := hex.DecodeString(txid)
	if err != nil || len(hash) != 32 {
		return nil, errors.New("invalid txid")
	}
	return &Input{TxHash: hash, Index: index}, nil
}

func MustInput(txid string, index uint32) *Input {
	input, err := NewInput(txid, index)
	if err != nil {
		panic(err)
	}
	return input
}
