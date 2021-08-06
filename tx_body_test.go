package libada

import (
	"encoding/hex"
	"testing"
)

func TestBody_Bytes(t *testing.T) {
	body := &TransactionBody{
		Inputs: []*Input{
			MustInput("9c37cb5294abce709bfa57bdcab039d75e212a503bc48d5b45e6ff4eb6272759", 1),
			MustInput("77df5bd720b6fb8699762fea21bdbc8193a61b770a59443266383ba01a6b8900", 1),
		},
		Outputs: []*Output{
			MustOutput("addr_test1qqth544yyqh8ahg0899ms59emls89cs9l9ra0n9nlrwtgahppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kqdaq2fy", 100),
			MustOutput("addr_test1qpnl5q67ypgpkpg8x4h8uaf8370npkke4rlrsfne3vlwpv8ppgsq2ykgpqpgewlkwkyhsqn29k8dxp7xthncvfwht9kq6npgqm", 10),
		},
		Fee: 100,
	}

	want := "a300828258209c37cb5294abce709bfa57bdcab039d75e212a503bc48d5b45e6ff4eb62727590182582077df5bd720b6fb8699762fea21bdbc8193a61b770a59443266383ba01a6b890001018282583900177a56a4202e7edd0f394bb850b9dfe072e205f947d7ccb3f8dcb476e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c18648258390067fa035e20501b0507356e7e75278f9f30dad9a8fe3826798b3ee0b0e10a200512c808028cbbf6758978026a2d8ed307c65de78625d7596c0a021864"
	if got := hex.EncodeToString(body.Bytes()); got != want {
		t.Errorf("\ngot: %s\nwant: %s", got, want)
	}
}
