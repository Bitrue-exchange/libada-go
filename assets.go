package libada

import (
	"encoding/hex"
	"errors"

	"github.com/fxamacker/cbor/v2"
)

type Asset struct {
	PolicyId [28]byte
	Name     AssetName
}

func NewAsset(policyId [28]byte, name AssetName) *Asset {
	return &Asset{PolicyId: policyId, Name: name}
}

func ParseAssetId(assetId string) (*Asset, error) {
	bytes, err := hex.DecodeString(assetId)
	if err != nil {
		return nil, err
	}

	if len(bytes) < 28 || len(bytes) > 60 {
		return nil, errors.New("invalid asset id length")
	}

	var asset Asset
	copy(asset.PolicyId[:], bytes[:28])
	if len(bytes) > 28 {
		asset.Name = NewAssetName(bytes[28:])
	}
	return &asset, nil
}

func MustParseAssetId(assetId string) *Asset {
	asset, err := ParseAssetId(assetId)
	if err != nil {
		panic(err)
	}
	return asset
}

type AssetName struct {
	data   [32]byte
	length int
}

func NewAssetName(b []byte) AssetName {
	if len(b) > 32 {
		panic("asset name size should less than 32 bytes")
	}
	var a [32]byte
	copy(a[:], b[:])
	return AssetName{length: len(b), data: a}
}

func (a AssetName) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(a.data[:a.length])
}

func (a *AssetName) UnmarshalCBOR(raw []byte) error {
	if a == nil {
		return errors.New("unmarshal to nil value")
	}
	var d []byte
	if err := cbor.Unmarshal(raw, &d); err != nil {
		return err
	}
	*a = NewAssetName(d)
	return nil
}

func (a AssetName) String() string {
	return string(a.data[:][:a.length])
}
