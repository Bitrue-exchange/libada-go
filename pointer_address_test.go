package libada

import (
	"testing"
)

func TestVariableNatDecode(t *testing.T) {
	testcase := []uint64{0, 127, 128, 255, 256275757658493284}

	for index, item := range testcase {
		encoding := VariableNatEncode(item)
		v, _, _ := VariableNatDecode(encoding)
		if v != item {
			t.Fatalf("[%d] failed", index)
		}
	}
}
