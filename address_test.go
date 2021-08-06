package libada

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    Address
		Kind    AddressKind
		Network Network
		wantErr bool
	}{
		{
			name:    "base mainnet",
			args:    args{"addr1qyy6nhfyks7wdu3dudslys37v252w2nwhv0fw2nfawemmn8k8ttq8f3gag0h89aepvx3xf69g0l9pf80tqv7cve0l33sdn8p3d"},
			Kind:    BaseAddressKind,
			Network: Mainnet,
			wantErr: false,
		},
		{
			name:    "base testnet",
			args:    args{"addr_test1qqy6nhfyks7wdu3dudslys37v252w2nwhv0fw2nfawemmn8k8ttq8f3gag0h89aepvx3xf69g0l9pf80tqv7cve0l33sw96paj"},
			Kind:    BaseAddressKind,
			Network: Testnet,
			wantErr: false,
		},
		{
			name:    "enterprise mainnet",
			args:    args{"addr1vyy6nhfyks7wdu3dudslys37v252w2nwhv0fw2nfawemmnqs6l44z"},
			Kind:    EnterpriseAddressKind,
			Network: Mainnet,
			wantErr: false,
		},
		{
			name:    "enterprise testnet",
			args:    args{"addr_test1vqy6nhfyks7wdu3dudslys37v252w2nwhv0fw2nfawemmnqtjtf68"},
			Kind:    EnterpriseAddressKind,
			Network: Testnet,
			wantErr: false,
		},
		{
			name:    "byron",
			args:    args{"DdzFFzCqrhsfZBimA3LVHNR4DHXvT4HBD21NUyKjKQGoaBGmzJJ4uNwDPh2zViGWuSFnobHyxJRUj1AFSUQASu2kLTcpg5PZEfPe4TfJ"},
			Kind:    LegacyAddressKind,
			Network: Mainnet,
			wantErr: false,
		},
		{
			name:    "byron simple",
			args:    args{"Ae2tdPwUPEZHtBmjZBF4YpMkK9tMSPTE2ADEZTPN97saNkhG78TvXdp3GDk"},
			Kind:    LegacyAddressKind,
			Network: Mainnet,
			wantErr: false,
		},
		{
			name:    "stake mainnet",
			args:    args{"stake1uyevw2xnsc0pvn9t9r9c7qryfqfeerchgrlm3ea2nefr9hqxdekzz"},
			Kind:    RewardAddressKind,
			Network: Mainnet,
			wantErr: false,
		},
		{
			name:    "stake testnet",
			args:    args{"stake_test1uqevw2xnsc0pvn9t9r9c7qryfqfeerchgrlm3ea2nefr9hqp8n5xl"},
			Kind:    RewardAddressKind,
			Network: Testnet,
			wantErr: false,
		},
		{
			name:    "pointer mainnet",
			args:    args{"addr1gyy6nhfyks7wdu3dudslys37v252w2nwhv0fw2nfawemmnyph3wczvf2dqflgt"},
			Kind:    PointerAddressKind,
			Network: Mainnet,
			wantErr: false,
		},
		{
			name:    "pointer testnet",
			args:    args{"addr_test1gqy6nhfyks7wdu3dudslys37v252w2nwhv0fw2nfawemmnqpqgps5mee0p"},
			Kind:    PointerAddressKind,
			Network: Testnet,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := DecodeAddress(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if tt.Network != addr.GetNetwork() {
					t.Errorf("Decode(): network got %v want %v", addr.GetNetwork(), tt.Network)
				}

				if tt.Kind != addr.Kind() {
					t.Errorf("Decode(): kind got %v want %v", addr.Kind(), tt.Kind)
				}
			}
		})
	}
}
