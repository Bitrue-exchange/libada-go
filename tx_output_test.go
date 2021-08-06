package libada

import (
	"crypto/rand"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/fxamacker/cbor/v2"
)

func TestNewOutput(t *testing.T) {
	var testdata = []struct {
		address string
		amount  uint64
		want    string
	}{
		{
			address: "addr_test1qqth544yyqh8ahg0899ms59emls89cs9l9ra0n9nlrwtgahppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kqdaq2fy",
			amount:  0x82,
			want:    "82583900177a56a4202e7edd0f394bb850b9dfe072e205f947d7ccb3f8dcb476e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c1882",
		},
		{
			address: "addr_test1qqth544yyqh8ahg0899ms59emls89cs9l9ra0n9nlrwtgahppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kqdaq2fy",
			amount:  100,
			want:    "82583900177a56a4202e7edd0f394bb850b9dfe072e205f947d7ccb3f8dcb476e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c1864",
		},
		{
			address: "addr_test1qpnl5q67ypgpkpg8x4h8uaf8370npkke4rlrsfne3vlwpv8ppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kq6npgqm",
			amount:  10,
			want:    "8258390067fa035e20501b0507356e7e75278f9f30dad9a8fe3826798b3ee0b0e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c0a",
		},
		{
			address: "Ae2tdPwUPEZEJsmHfRC3gutyFZuJSXLumCu1oovMSps3EVewbqoHDnoZm3d",
			amount:  100,
			want:    "82582b82d818582183581caf67c9c546e2a69870cc7c2bc5f6633253dfa36143ffddfda4523ee6a0001a78f496c81864",
		},
		{
			address: "DdzFFzCqrht7FAf8MpryP1p8sgkmFRUnDpifnnu4ZxpBjbCTSDwJVAaDsDqrC7SLYFx8fUrDcNNsD4AMiUgg2wGywTVpcfB1F3AHrGkv",
			amount:  100,
			want:    "82584c82d818584283581cd1ed7fa89a505f9c83de0ef1178b383258e71b2375b4e1279d5e0e0fa101581e581cc655d01a842282dd46918c9198d9e98efad24af87b9b0205aed69fd4001a237a87d71864",
		},
		{
			address: "addr1q860w3lz6zd9tn0uqcvw782qmcs6lnztz2asy0555sr2kvh57ar795y62hxlcpscauw5ph3p4lxyky4mqglfffqx4veq6gs3hg",
			amount:  100,
			want:    "82583901f4f747e2d09a55cdfc0618ef1d40de21afcc4b12bb023e94a406ab32f4f747e2d09a55cdfc0618ef1d40de21afcc4b12bb023e94a406ab321864",
		},
	}

	for i, item := range testdata {
		output1 := MustOutput(item.address, item.amount)
		rawbytes, err := cbor.Marshal(output1)
		if err != nil {
			t.Fatal(err)
		}

		if got := hex.EncodeToString(rawbytes); got != item.want {
			t.Errorf("[%d] want %s got %s", i, item.want, got)
			continue
		}

		var output2 Output
		if err := cbor.Unmarshal(rawbytes, &output2); err != nil {
			t.Errorf("[%d] can not decode raw bytes: %s", i, err)
			continue
		}

		if !reflect.DeepEqual(output1, &output2) {
			t.Errorf("[%d] consistent output", i)
		}
	}
}

func TestNewOutputWithAsset(t *testing.T) {
	var testdata = []struct {
		address     string
		amount      uint64
		assetId     string
		assetAmount uint64
		want        string
	}{
		{
			address:     "addr_test1vqpjd93t42ju4majh9tcz69z2fvmaeyxzxvpr3x95g9mw4sxmvk7w",
			amount:      100,
			assetId:     "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354",
			assetAmount: 10,
			want:        "82581d600326962baaa5caefb2b9578168a25259bee486119811c4c5a20bb756821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a",
		},
		{
			address:     "addr_test1vqpjd93t42ju4majh9tcz69z2fvmaeyxzxvpr3x95g9mw4sxmvk7w",
			amount:      100,
			assetId:     "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9",
			assetAmount: 10,
			want:        "82581d600326962baaa5caefb2b9578168a25259bee486119811c4c5a20bb756821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1400a",
		},
		{
			address:     "addr1q860w3lz6zd9tn0uqcvw782qmcs6lnztz2asy0555sr2kvh57ar795y62hxlcpscauw5ph3p4lxyky4mqglfffqx4veq6gs3hg",
			amount:      100,
			assetId:     "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354",
			assetAmount: 10,
			want:        "82583901f4f747e2d09a55cdfc0618ef1d40de21afcc4b12bb023e94a406ab32f4f747e2d09a55cdfc0618ef1d40de21afcc4b12bb023e94a406ab32821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a",
		},
		{
			address:     "Ae2tdPwUPEZEJsmHfRC3gutyFZuJSXLumCu1oovMSps3EVewbqoHDnoZm3d",
			amount:      100,
			assetId:     "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354",
			assetAmount: 10,
			want:        "82582b82d818582183581caf67c9c546e2a69870cc7c2bc5f6633253dfa36143ffddfda4523ee6a0001a78f496c8821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a",
		},
		{
			address:     "DdzFFzCqrht7FAf8MpryP1p8sgkmFRUnDpifnnu4ZxpBjbCTSDwJVAaDsDqrC7SLYFx8fUrDcNNsD4AMiUgg2wGywTVpcfB1F3AHrGkv",
			amount:      100,
			assetId:     "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354",
			assetAmount: 10,
			want:        "82584c82d818584283581cd1ed7fa89a505f9c83de0ef1178b383258e71b2375b4e1279d5e0e0fa101581e581cc655d01a842282dd46918c9198d9e98efad24af87b9b0205aed69fd4001a237a87d7821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a",
		},
	}

	for i, item := range testdata {
		output1 := MustOutputWithAssets(item.address, item.amount, item.assetId, item.assetAmount)
		rawbytes, err := cbor.Marshal(output1)
		if err != nil {
			t.Errorf("[%d] can not encode output: %s", i, err)
			continue
		}

		if getbytes := hex.EncodeToString(rawbytes); getbytes != item.want {
			t.Errorf("[%d]\nwant bytes %s\ngot bytes %s", i, item.want, getbytes)
			continue
		}

		var output2 Output
		if err := cbor.Unmarshal(rawbytes, &output2); err != nil {
			t.Errorf("[%d] can not decode raw bytes: %s", i, err)
		}

		if !reflect.DeepEqual(output1, &output2) {
			t.Errorf("[%d] consistent output", i)
		}
	}
}

func TestNewOutput_slice(t *testing.T) {
	outputs := []*Output{
		MustOutput("addr_test1qqth544yyqh8ahg0899ms59emls89cs9l9ra0n9nlrwtgahppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kqdaq2fy", 100),
		MustOutput("addr_test1qpnl5q67ypgpkpg8x4h8uaf8370npkke4rlrsfne3vlwpv8ppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kq6npgqm", 10),
		MustOutputWithAssets("addr_test1vqpjd93t42ju4majh9tcz69z2fvmaeyxzxvpr3x95g9mw4sxmvk7w", 100, "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354", 10),
	}

	data, err := cbor.Marshal(outputs)
	if err != nil {
		t.Fatal(err)
	}

	want := "8382583900177a56a4202e7edd0f394bb850b9dfe072e205f947d7ccb3f8dcb476e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c18648258390067fa035e20501b0507356e7e75278f9f30dad9a8fe3826798b3ee0b0e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c0a82581d600326962baaa5caefb2b9578168a25259bee486119811c4c5a20bb756821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a"
	if got := hex.EncodeToString(data); got != want {
		t.Fatalf("want %s got %s", want, got)
	}
}

func TestOutput_MinAdaRequired(t *testing.T) {
	var policyId [28]byte

	var random32Bytes = func() []byte {
		var a [32]byte
		if _, err := rand.Read(a[:]); err != nil {
			t.Fatal(err)
		}
		return a[:]
	}

	type fields struct {
		Address Address
		Amount  uint64
		Asset   map[[28]byte]map[AssetName]uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name:   "no_token_minimum",
			fields: fields{Address: MustDecodeAddress("addr_test1vrkz8ndlgg4dsnakzgt3l3yrjxwrs7f9w4x852eydfhp6ac9hv97j"), Amount: 100},
			want:   1e6,
		},
		{
			name: "one_policy_one_smallest_name",
			fields: fields{
				Address: MustDecodeAddress("addr_test1vrkz8ndlgg4dsnakzgt3l3yrjxwrs7f9w4x852eydfhp6ac9hv97j"),
				Amount:  1407406,
				Asset: map[[28]byte]map[AssetName]uint64{
					policyId: {NewAssetName(nil): 1},
				},
			},
			want: 1407406,
		},
		{
			name: "one_policy_one_small_name",
			fields: fields{
				Address: MustDecodeAddress("addr_test1vrkz8ndlgg4dsnakzgt3l3yrjxwrs7f9w4x852eydfhp6ac9hv97j"),
				Amount:  1444443,
				Asset: map[[28]byte]map[AssetName]uint64{
					policyId: {NewAssetName([]byte{1}): 1},
				},
			},
			want: 1444443,
		},
		{
			name: "one_policy_one_largest_name",
			fields: fields{
				Address: MustDecodeAddress("addr_test1vrkz8ndlgg4dsnakzgt3l3yrjxwrs7f9w4x852eydfhp6ac9hv97j"),
				Amount:  1555554,
				Asset: map[[28]byte]map[AssetName]uint64{
					policyId: {NewAssetName(make([]byte, 32)): 1},
				},
			},
			want: 1555554,
		},
		{
			name: "one_policy_three_small_names",
			fields: fields{
				Address: MustDecodeAddress("addr_test1vrkz8ndlgg4dsnakzgt3l3yrjxwrs7f9w4x852eydfhp6ac9hv97j"),
				Amount:  1555554,
				Asset: map[[28]byte]map[AssetName]uint64{
					policyId: {NewAssetName([]byte{1}): 1, NewAssetName([]byte{2}): 1, NewAssetName([]byte{3}): 1},
				},
			},
			want: 1555554,
		},
		{
			name: "one_policy_three_largest_names",
			fields: fields{
				Address: MustDecodeAddress("addr_test1vrkz8ndlgg4dsnakzgt3l3yrjxwrs7f9w4x852eydfhp6ac9hv97j"),
				Amount:  1962961,
				Asset: map[[28]byte]map[AssetName]uint64{
					policyId: {NewAssetName(random32Bytes()): 1, NewAssetName(random32Bytes()): 1, NewAssetName(random32Bytes()): 1},
				},
			},
			want: 1962961,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Output{
				Address: tt.fields.Address,
				Amount:  tt.fields.Amount,
				Assets:  tt.fields.Asset,
			}
			if got := o.MinAdaRequired(); got != tt.want {
				t.Errorf("Output.MinAdaRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutput_MustAddAsset(t *testing.T) {
	type testasset struct {
		id     string
		amount uint64
	}

	var testdata = []struct {
		name    string
		address string
		amount  uint64
		assets  []testasset
		want    string
		want2   string
	}{
		{
			name:    "one policy with two assets",
			address: "addr1v8rjcktafmp8xnr7szlhg9mxxeg64vycuzd3y2mes6ahf7sjyfkx6",
			amount:  100,
			assets: []testasset{
				{id: "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354", amount: 10},
				{id: "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9324e44", amount: 100},
			},
			want:  "82581d61c72c597d4ec2734c7e80bf7417663651aab098e09b122b7986bb74fa821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a2433153540a43324e441864",
			want2: "82581d61c72c597d4ec2734c7e80bf7417663651aab098e09b122b7986bb74fa821864a1581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a243324e441864433153540a",
		},
		{
			name:    "two policy with two assets",
			address: "addr1v8rjcktafmp8xnr7szlhg9mxxeg64vycuzd3y2mes6ahf7sjyfkx6",
			amount:  100,
			assets: []testasset{
				{id: "76802823e9009022a46e63d0842b3e8f46f164cf39997ed65c98d3c5315354", amount: 10},
				{id: "77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354", amount: 10},
			},
			want:  "82581d61c72c597d4ec2734c7e80bf7417663651aab098e09b122b7986bb74fa821864a2581c76802823e9009022a46e63d0842b3e8f46f164cf39997ed65c98d3c5a1433153540a581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a",
			want2: "82581d61c72c597d4ec2734c7e80bf7417663651aab098e09b122b7986bb74fa821864a2581c77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9a1433153540a581c76802823e9009022a46e63d0842b3e8f46f164cf39997ed65c98d3c5a1433153540a",
		},
	}

	for _, item := range testdata {
		output := MustOutput(item.address, item.amount)
		for _, testAssetItem := range item.assets {
			output.MustAddAsset(testAssetItem.id, testAssetItem.amount)
		}
		data, err := cbor.Marshal(output)
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(data); got != item.want && got != item.want2 {
			t.Errorf("[%s]\nwant %s\ngot %s", item.name, item.want, got)
		}
	}
}
