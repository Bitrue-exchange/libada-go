package libada

import (
	"encoding/hex"
	"testing"
)

func TestByron_String_simple(t *testing.T) {
	xpub, err := hex.DecodeString("f164779dc38920e994c98e407cf5d4557b5217d9eef58f24b9a96d2d681d32a5a541bdc9ba897c1156a2e0ee7c49a42f1322aa06f50047d7c3fc61c6c975ecb7")
	if err != nil {
		t.Fatal(err)
	}

	mainnetAddress := "Ae2tdPwUPEZ6MEMHsfoiwtAygYmsyixB3icdHw6Ex87WVarLNP4SKqkgBU9"
	{
		byron, err := NewSimpleLegacyAddress(xpub)
		if err != nil {
			t.Fatal("decode 1", err)
		}
		if got := byron.String(); got != mainnetAddress {
			t.Fatalf("want %s got %s", mainnetAddress, got)
		}
		if byron.GetNetwork() != Mainnet {
			t.Fatal("want Mainnet")
		}

		if byron.Attrs.Network != nil {
			t.Fatal("mainnet address network should null")
		}
	}

	{
		byron, err := DecodeAddress(mainnetAddress)
		if err != nil {
			t.Fatal("decode 2", err)
		}

		if byron.Kind() != LegacyAddressKind {
			t.Fatal("should be byron type address")
		}

		if byron.GetNetwork() != Mainnet {
			t.Fatal("want Mainnet 2")
		}

		if byron.(*LegacyAddress).Attrs.Network != nil {
			t.Fatal("mainnet address network should null 2")
		}
	}

	testnetAddress := "2cWKMJemoBaipzQe9BArYdo2iPUfJQdZAjm4iCzDA1AfNxJSTgm9FZQTmFCYhKkeYrede"
	{
		byron, err := DecodeAddress(testnetAddress)
		if err != nil {
			t.Fatal("decode 3", err)
		}

		if byron.Kind() != LegacyAddressKind {
			t.Fatal("should be byron type address 2")
		}

		if *byron.(*LegacyAddress).Attrs.Network != Testnet.ProtocolMagic() {
			t.Fatal("testnet address network should be not null")
		}
	}
}

func TestByron_String_complex(t *testing.T) {
	{
		mainnet := "DdzFFzCqrhtBzWpZB533kKmquDpWmHDQmBUAPqXdbKQTc9gycHGx68CfUPCkpMhE9YFfVNAuFFQgKk2T9YgN3aP2rbRrUFwYx9j9twyF"
		byron, err := DecodeAddress(mainnet)
		if err != nil {
			t.Fatal(err)
		}

		if byron.Kind() != LegacyAddressKind {
			t.Fatal("should be byron type address")
		}

		if byron.GetNetwork() != Mainnet {
			t.Fatal("want Mainnet")
		}

		if got := byron.String(); got != mainnet {
			t.Fatalf("want %s got %s", mainnet, got)
		}

		if byron.(*LegacyAddress).Attrs.Network != nil {
			t.Fatal("mainnet address network should null")
		}
	}
	{
		testnet := "37btjrVyb4KGM5rFFreGtZAs4PFB2Drb37uXRHebh8rCeVWFkW8De8XAbYqvfQrAqVthfJp9Qy2YzbzNhWSiUGY3D7yJkRkChyMveKCWT8qUTNEu6e"
		byron, err := DecodeAddress(testnet)
		if err != nil {
			t.Fatal(err)
		}

		if byron.Kind() != LegacyAddressKind {
			t.Fatal("should be byron type address")
		}

		if byron.GetNetwork() != Testnet {
			t.Fatal("want Testnet")
		}

		if got := byron.String(); got != testnet {
			t.Fatalf("want %s got %s", testnet, got)
		}

		if *byron.(*LegacyAddress).Attrs.Network != Testnet.ProtocolMagic() {
			t.Fatal("testnet address network should be not null")
		}
	}
}
