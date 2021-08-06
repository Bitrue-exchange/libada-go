package libada

import (
	"encoding/hex"
	"testing"

	"github.com/fxamacker/cbor/v2"
)

func TestInput_Bytes(t *testing.T) {
	testdata := []struct {
		txid  string
		index uint32
		want  string
	}{
		{
			txid:  "9c37cb5294abce709bfa57bdcab039d75e212a503bc48d5b45e6ff4eb6272759",
			index: 1,
			want:  "8258209c37cb5294abce709bfa57bdcab039d75e212a503bc48d5b45e6ff4eb627275901",
		},
		{
			txid:  "77df5bd720b6fb8699762fea21bdbc8193a61b770a59443266383ba01a6b8900",
			index: 1,
			want:  "82582077df5bd720b6fb8699762fea21bdbc8193a61b770a59443266383ba01a6b890001",
		},
	}
	for index, item := range testdata {
		data, err := cbor.Marshal(MustInput(item.txid, item.index))
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(data); got != item.want {
			t.Errorf("[%d] got %s want %s", index, got, item.want)
		}
	}
}

func TestInput_Bytes_slice(t *testing.T) {
	input_0 := MustInput("9c37cb5294abce709bfa57bdcab039d75e212a503bc48d5b45e6ff4eb6272759", 1)
	input_1 := MustInput("77df5bd720b6fb8699762fea21bdbc8193a61b770a59443266383ba01a6b8900", 1)
	data, err := cbor.Marshal([]*Input{input_0, input_1})
	if err != nil {
		t.Fatal(err)
	}
	want := "828258209c37cb5294abce709bfa57bdcab039d75e212a503bc48d5b45e6ff4eb62727590182582077df5bd720b6fb8699762fea21bdbc8193a61b770a59443266383ba01a6b890001"
	if got := hex.EncodeToString(data); got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}
