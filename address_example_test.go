package libada

import "fmt"

func ExampleDecodeAddress() {
	RawAddressList := []string{
		"addr1v9wa6entm75duchtu50mu6u6hkagdgqzaevt0cwryaw3pnca870vt",
		"Ae2tdPwUPEZCdHLQP8Nke2AogJ7VgGiHbJDV3N1mWyQxy5AhwGgihL44t8w",
		"addr1q87ywcqelrrm9zrn4f9v5te2ss2w6f0j4ca39pwhsgxnupuec8u08zuyfu64xecytwcuc8nm6xkn0xj2sqx7m7g07fqqf96uvu",
		"DdzFFzCqrht7FAf8MpryP1p8sgkmFRUnDpifnnu4ZxpBjbCTSDwJVAaDsDqrC7SLYFx8fUrDcNNsD4AMiUgg2wGywTVpcfB1F3AHrGkv",
	}

	for _, address := range RawAddressList {
		got := MustDecodeAddress(address).String()
		fmt.Println(got == address)
	}

	// Output:
	// true
	// true
	// true
	// true
}
