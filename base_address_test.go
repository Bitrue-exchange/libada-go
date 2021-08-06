package libada

import (
	"encoding/hex"
	"testing"
)

func TestBase_Bytes(t *testing.T) {
	spendPubkey, err := hex.DecodeString("73fea80d424276ad0978d4fe5310e8bc2d485f5f6bb3bf87612989f112ad5a7d")
	if err != nil {
		t.Fatal(err)
	}

	stakePubkey, err := hex.DecodeString("2c041c9c6a676ac54d25e2fdce44c56581e316ae43adc4c7bf17f23214d8d892")
	if err != nil {
		t.Fatal(err)
	}

	mainnet := &BaseAddress{
		Network: Mainnet,
		Payment: NewKeysStakeCred(Blake2b224(spendPubkey)),
		Stake:   NewKeysStakeCred(Blake2b224(stakePubkey)),
	}

	wantMainnet := "addr1qx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer3jcu5d8ps7zex2k2xt3uqxgjqnnj83ws8lhrn648jjxtwqfjkjv7"
	if got := mainnet.String(); got != wantMainnet {
		t.Fatalf("mainnet want %s got %s", wantMainnet, got)
	}

	testnet := &BaseAddress{
		Network: Testnet,
		Payment: mainnet.Payment,
		Stake:   mainnet.Stake,
	}
	wantTestnet := "addr_test1qz2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer3jcu5d8ps7zex2k2xt3uqxgjqnnj83ws8lhrn648jjxtwq2ytjqp"
	if got := testnet.String(); got != wantTestnet {
		t.Fatalf("testnet want %s got %s", wantTestnet, got)
	}
}
