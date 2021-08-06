package libada

import (
	"encoding/hex"
	"testing"
)

func TestBlake2b224(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"c6e47dceb9235ad31fec01ef3e34b02839d22fb583cd28fac7ec81ebbe653a39"}, "c5fb57853a4a5bbb13bcd494ee0facd4e1fca33720a640664df15134"},
		{"2", args{"a893e820e9fa5452c8293512b47ab41562ff7a791530be0d371f8322c4b905f5"}, "515f7a4b41c731bff4414cb6fcd273d50ed75ed02d307a31e3647a9f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg, err := hex.DecodeString(tt.args.raw)
			if err != nil {
				t.Error(err)
				return
			}
			if got := hex.EncodeToString(Blake2b224(arg)); got != tt.want {
				t.Errorf("Hash224() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha3AndBlake2b224(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"1", "b1eff5cc8aa913b221196608061f1b0d", "12fef368d7bf55ea025bb11a8bdefe7ea4dff54d4dee0881291ee660"},
		{"2", "2b315947ef4306fa00c68ba466a2007d", "e9069dec75762a667c37a4eb7386347a2460c67ee29e4939c6cc7d8c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg, err := hex.DecodeString(tt.args)
			if err != nil {
				t.Error(tt.name, "decode hex failed")
				return
			}

			if got := hex.EncodeToString(Sha3AndBlake2b224(arg)); got != tt.want {
				t.Errorf("Sha3AndBlake2b224() = %v, want %v", got, tt.want)
			}
		})
	}
}
