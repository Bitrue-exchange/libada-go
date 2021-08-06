package libada

import (
	"encoding/hex"
	"testing"
)

func TestEnterprise_String(t *testing.T) {
	type fields struct {
		Network Network
		Payment StakeCredential
	}

	MustDecodeHex := func(raw string) []byte {
		b, err := hex.DecodeString(raw)
		if err != nil {
			t.Fatal(err)
		}
		return b
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "testnet address",
			fields: fields{
				Network: Testnet,
				Payment: StakeCredential{
					Kind: KeyStakeCredentialType,
					Data: MustDecodeHex("7ffcf441f7bae81f15e61fda81833eef8ee779a1f83aaab2371598a5"),
				},
			},
			want: "addr_test1vplleazp77aws8c4uc0a4qvr8mhcaeme58ur424jxu2e3fg6cgfgj",
		},
		{
			name: "mainnet address",
			fields: fields{
				Network: Mainnet,
				Payment: StakeCredential{
					Kind: KeyStakeCredentialType,
					Data: MustDecodeHex("7ffcf441f7bae81f15e61fda81833eef8ee779a1f83aaab2371598a5"),
				},
			},
			want: "addr1v9lleazp77aws8c4uc0a4qvr8mhcaeme58ur424jxu2e3fgpsu48h",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EnterpriseAddress{
				Network: tt.fields.Network,
				Payment: tt.fields.Payment,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Enterprise.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEnterprise(t *testing.T) {
	pubkey := []byte{62, 10, 246, 49, 19, 79, 17, 153, 156, 252, 94, 21, 12, 6, 65, 246, 17, 58, 139, 207, 101, 15, 20, 43, 229, 253, 173, 13, 9, 102, 211, 130}
	type args struct {
		key     []byte
		kind    StakeCredentialType
		network Network
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testnet",
			args: args{key: pubkey, kind: KeyStakeCredentialType, network: Testnet},
			want: "addr_test1vqpjd93t42ju4majh9tcz69z2fvmaeyxzxvpr3x95g9mw4sxmvk7w",
		},
		{
			name: "mainnet",
			args: args{key: pubkey, kind: KeyStakeCredentialType, network: Mainnet},
			want: "addr1vypjd93t42ju4majh9tcz69z2fvmaeyxzxvpr3x95g9mw4sanc23t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnterpriseAddress(tt.args.key, tt.args.kind, tt.args.network).String(); got != tt.want {
				t.Errorf("NewEnterprise() = %s, want %s", got, tt.want)
			}
		})
	}
}
