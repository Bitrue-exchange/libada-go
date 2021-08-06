package libada

import (
	"errors"

	"github.com/fxamacker/cbor/v2"
)

// Output transaction output
type Output struct {
	Address Address
	Amount  uint64
	Assets  map[[28]byte]map[AssetName]uint64
}

func NewOutput(addr Address, amount uint64) *Output {
	return &Output{Address: addr, Amount: amount}
}

func NewOutputWithAsset(addr Address, amount uint64, token *Asset, tokenAmount uint64) *Output {
	return &Output{
		Address: addr,
		Amount:  amount,
		Assets:  map[[28]byte]map[AssetName]uint64{token.PolicyId: {token.Name: tokenAmount}},
	}
}

func MustOutput(addr string, amount uint64) *Output {
	a, err := DecodeAddress(addr)
	if err != nil {
		panic(err)
	}
	return NewOutput(a, amount)
}

func MustOutputWithAssets(addr string, amount uint64, assetId string, tokenAmount uint64) *Output {
	a, err := DecodeAddress(addr)
	if err != nil {
		panic(err)
	}
	asset, err := ParseAssetId(assetId)
	if err != nil {
		panic(err)
	}
	return NewOutputWithAsset(a, amount, asset, tokenAmount)
}

// AddAsset appends asset and asset amount to Outputs
func (o *Output) AddAsset(asset *Asset, tokenAmount uint64) *Output {
	if o.Assets == nil {
		o.Assets = make(map[[28]byte]map[AssetName]uint64)
	}
	if assetList, hasPolicyId := o.Assets[asset.PolicyId]; hasPolicyId {
		assetList[asset.Name] += tokenAmount
	} else {
		o.Assets[asset.PolicyId] = map[AssetName]uint64{asset.Name: tokenAmount}
	}
	return o
}

// AddAssetMap copys raw asset to outputs
func (o *Output) AddAssetMap(raw map[[28]byte]map[AssetName]uint64) {
	for policyId, assets := range raw {
		for assetName, tokenAmount := range assets {
			o.AddAsset(NewAsset(policyId, assetName), tokenAmount)
		}
	}
}

// AddAmount adds ADA for Output
func (o *Output) AddAmount(amount uint64) *Output {
	o.Amount += amount
	return o
}

// SetAmount sets amount for Output
func (o *Output) SetAmount(amount uint64) *Output {
	o.Amount = amount
	return o
}

func (o *Output) MustAddAsset(assetId string, tokenAmount uint64) *Output {
	asset, err := ParseAssetId(assetId)
	if err != nil {
		panic(err)
	}
	return o.AddAsset(asset, tokenAmount)
}

func (o *Output) MarshalCBOR() ([]byte, error) {
	if o.Address == nil {
		return nil, errors.New("nil address to marshal")
	}
	if len(o.Assets) > 0 {
		return cbor.Marshal([]interface{}{o.Address.Bytes(), []interface{}{o.Amount, o.Assets}})
	}
	return cbor.Marshal([]interface{}{o.Address.Bytes(), o.Amount})
}

func (o *Output) UnmarshalCBOR(data []byte) error {
	if o == nil {
		return errors.New("unmarshal to nil output")
	}

	type Raw struct {
		_          struct{} `cbor:",toarray"`
		RawAddress []byte
		RawAmount  cbor.RawMessage
	}

	type RawAmount struct {
		_      struct{} `cbor:",toarray"`
		Amount uint64
		Assets map[[28]byte]map[AssetName]uint64
	}

	var raw Raw
	if err := cbor.Unmarshal(data, &raw); err != nil {
		return err
	}

	addr, err := DecodeRawAddress(raw.RawAddress)
	if err != nil {
		return err
	}

	if len(raw.RawAmount) == 0 {
		return errors.New("no raw amount cbor bytes")
	}

	switch raw.RawAmount[0] {
	case 0x82:
		var amount RawAmount
		if err := cbor.Unmarshal(raw.RawAmount, &amount); err != nil {
			return err
		}
		*o = Output{Address: addr, Amount: amount.Amount, Assets: amount.Assets}
		return nil
	default:
		var amount uint64
		if err := cbor.Unmarshal(raw.RawAmount, &amount); err != nil {
			return err
		}
		*o = Output{Address: addr, Amount: amount}
		return nil
	}
}

// SetMinAda sets min ada required
func (o *Output) SetMinAda() *Output {
	o.Amount = o.MinAdaRequired()
	return o
}

func quot(a, b int) int {
	return (a - (a % b)) / b
}

func roundupBytes2Words(b int) int {
	return quot(b+7, 8)
}

// MinAdaRequired computes required ada for this output
// learn more at https://docs.cardano.org/native-tokens/minimum-ada-value-requirement
func (o *Output) MinAdaRequired() uint64 {
	const MinADA = 1e6

	if len(o.Assets) == 0 {
		return MinADA
	}

	// TODO: coinSize should be 2, it will change in the future
	const coinSize = 0
	const txOutLenWithoutValue = 14
	const txInlength = 7
	const utxoEntrySizeWithoutValue = 6 + txOutLenWithoutValue + txInlength
	const adaOnlyUTxOSize = utxoEntrySizeWithoutValue + coinSize

	var numAssets = 0
	var sumAssetNameLengths = 0
	var sumPolicyIdLengths = 28 * len(o.Assets)
	for _, item := range o.Assets {
		numAssets += len(item)
		for assetName := range item {
			sumAssetNameLengths += assetName.length
		}
	}

	const (
		k0 = 6
		k1 = 12
		k2 = 1
	)

	bundleSize := k0 + roundupBytes2Words(numAssets*k1+sumAssetNameLengths+(k2*sumPolicyIdLengths))
	required := quot(MinADA, adaOnlyUTxOSize) * (utxoEntrySizeWithoutValue + bundleSize)
	if required < MinADA {
		return MinADA
	}
	return uint64(required)
}
