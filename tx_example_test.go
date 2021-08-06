package libada

import (
	"encoding/hex"
	"fmt"

	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/islishude/bip32"
)

func ExampleNewTx() {
	_, rawxprv, err := bech32.Decode("xprv1wq9th6qy7ej8w03cgas2lzwvzncvwaje8ek8lps6u6mwyhtpaewap0w2f9hhrpg6tvrffcvz5ecevvhfl6vnznswss96kqyncm2drut3kfps0pv0r4q3fuy6urzs52ywdpzx55muncgf283fr378hwzwp5p27ewe")
	if err != nil {
		panic(err)
	}

	rootprv, err := bip32.NewXPrv(rawxprv)
	if err != nil {
		panic(err)
	}
	rootpub, rootxpub := rootprv.PublicKey(), rootprv.XPub().Bytes()
	shellAddress := NewKeyedEnterpriseAddress(rootpub, Testnet)
	byronAddress, err := NewSimpleLegacyAddress(rootxpub)
	if err != nil {
		panic(err)
	}

	// ADA only transaction
	{
		tx := NewTx()
		tx.AddInputs(MustInput("2558aad25ec6b0e74009f36dc60d7fec6602ce43d603e80c9edde9dd54c78eb4", 0))
		tx.AddOutputs(NewOutput(shellAddress, 9000000)).SetFee(1000000).SetInvalidAfter(832163)
		tx.AddKeyWitness(NewKeysWitness(rootpub, rootprv.Sign(tx.Hash())))

		fmt.Println(hex.EncodeToString(tx.Bytes()))
	}

	// With Token transaction
	{
		tx := NewTx()
		tx.AddInputs(MustInput("9640acb862c4060ab61d0c77a18348bd22c76fd855dae12d09255b5cbebe44f3", 0))

		nativeAsset := MustParseAssetId("2ad8fbdc6f18d97ce61b113301f05ff6b2a241c4e0bf5d06127d30a24f4343")
		output0 := NewOutput(shellAddress, 0).AddAsset(nativeAsset, 10000).
			SetMinAda() // you can use 0 ADA at first and set min ada at last

		tx.AddOutputs(output0, NewOutput(shellAddress, 2e6)).
			SetFee(2000000).
			SetInvalidAfter(832163).
			SetInvalidBefore(822163)

		witness, err := NewBootstrapWitness(byronAddress, rootxpub, rootprv.Sign(tx.Hash()))
		if err != nil {
			panic(err)
		}
		tx.AddBootstrapWitness(witness)
		fmt.Println(tx.ID())
	}

	// Output:
	// 83a400818258202558aad25ec6b0e74009f36dc60d7fec6602ce43d603e80c9edde9dd54c78eb400018182581d60c5fb57853a4a5bbb13bcd494ee0facd4e1fca33720a640664df151341a00895440021a000f4240031a000cb2a3a10081825820c6e47dceb9235ad31fec01ef3e34b02839d22fb583cd28fac7ec81ebbe653a395840e94817956d902702ed6172db8edca0193b33aba511a72e1998cc18aa492be8cffaca96c8305ffbb2836c3c04c886e23302e222b99185b87b701a40a10c50660ff6
	// 08c1e412dc65a8220f650a396ac5ca4758d8df7ff65defb790af2437b2d84df9
}
