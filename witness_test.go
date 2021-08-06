package libada

import (
	"encoding/hex"
	"testing"

	"github.com/Bitrue-exchange/libada-go/internal/bech32"
	"github.com/fxamacker/cbor/v2"
	"github.com/islishude/bip32"
)

func TestNewKeysWitness(t *testing.T) {
	hrp, vkey, err := bech32.Decode("ed25519_pk1dgaagyh470y66p899txcl3r0jaeaxu6yd7z2dxyk55qcycdml8gszkxze2")
	if err != nil || len(vkey) != 32 || hrp != "ed25519_pk" {
		t.Fatal("invalid vkey")
	}

	sign, err := hex.DecodeString("08a722892c7e4c56a927988504f63f9174faef632d5a7f6e6371d27c6cb662561bd9348c7478db5908fa42962143156f3c044eadda0aeecd0d9ea81958e7cb99")
	if err != nil || len(sign) != 64 {
		t.Fatal("invalid sign")
	}

	witness := &Witness{Keys: []*KeysWitness{NewKeysWitness(vkey, sign)}}

	data, err := cbor.Marshal(witness)
	if err != nil {
		t.Fatal(err)
	}

	want := "a100818258206a3bd412f5f3c9ad04e52acd8fc46f9773d373446f84a69896a5018261bbf9d1584008a722892c7e4c56a927988504f63f9174faef632d5a7f6e6371d27c6cb662561bd9348c7478db5908fa42962143156f3c044eadda0aeecd0d9ea81958e7cb99"
	if got := hex.EncodeToString(data); got != want {
		t.Errorf("got: %s want : %s", got, want)
	}
}

func TestNewBootstrapWitness(t *testing.T) {
	_, rawXprv, err := bech32.Decode("xprv1wq9th6qy7ej8w03cgas2lzwvzncvwaje8ek8lps6u6mwyhtpaewap0w2f9hhrpg6tvrffcvz5ecevvhfl6vnznswss96kqyncm2drut3kfps0pv0r4q3fuy6urzs52ywdpzx55muncgf283fr378hwzwp5p27ewe")
	if err != nil {
		t.Fatal(err)
	}

	xprv, err := bip32.NewXPrv(rawXprv)
	if err != nil {
		t.Fatal(err)
	}

	txid, err := hex.DecodeString("401915f86d55a8a722bb84c49c73400fabfa378c289cd4f513d294c1fe415aa0")
	if err != nil {
		t.Fatal(err)
	}

	byron, err := NewSimpleLegacyAddress(xprv.XPub().Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if got := byron.String(); got != "Ae2tdPwUPEZ4BWbYfmcKUvfG42w7K2y4jbmBTczAJ5gs36J44Kscez8uvqG" {
		t.Fatal("byron adress isn't correct", got)
	}

	witness, err := NewBootstrapWitness(byron, xprv.XPub().Bytes(), xprv.Sign(txid))
	if err != nil {
		t.Fatal(err)
	}

	{
		vkey := "c6e47dceb9235ad31fec01ef3e34b02839d22fb583cd28fac7ec81ebbe653a39"
		if got := hex.EncodeToString(witness.VKey); got != vkey {
			t.Errorf("vkey isn't equal: got %s want %s", got, vkey)
		}
	}

	{
		sig := "7fb7c44821841bc48adccdd287e29929dad35cc9e39bed957cd6214903dce56cbe3a2d119dcac95bb42c72ec32de86a45a52f95c6b6d931c6dcfbb89222ff10a"
		if gotsig := hex.EncodeToString(witness.Sign); gotsig != sig {
			t.Errorf("sig isn't equal: got %s want %s", gotsig, sig)
		}
	}

	{
		attributes := "a0"
		if gotsf := hex.EncodeToString(witness.Attribute); gotsf != attributes {
			t.Errorf("suffix isn't equal: got %s want %s", gotsf, attributes)
		}
	}

	{
		data, err := cbor.Marshal(witness)
		if err != nil {
			t.Fatal(err)
		}
		want := "845820c6e47dceb9235ad31fec01ef3e34b02839d22fb583cd28fac7ec81ebbe653a3958407fb7c44821841bc48adccdd287e29929dad35cc9e39bed957cd6214903dce56cbe3a2d119dcac95bb42c72ec32de86a45a52f95c6b6d931c6dcfbb89222ff10a582071b24307858f1d4114f09ae0c50a288e68446a537c9e10951e291c7c7bb84e0d41a0"
		if got := hex.EncodeToString(data); want != got {
			t.Errorf("witness isn't equal: got %s want %s", got, want)
		}
	}
}
