package libada

import (
	"reflect"
	"testing"
)

func TestAssetIdByRaw(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Asset
		wantErr bool
	}{
		{
			"case 1",
			"77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9315354",
			&Asset{PolicyId: [28]byte{119, 231, 164, 104, 136, 134, 70, 117, 116, 4, 92, 106, 1, 38, 209, 64, 100, 75, 151, 225, 43, 100, 77, 71, 62, 113, 211, 233}, Name: NewAssetName([]byte("1ST"))},
			false,
		},
		{
			"case 2",
			"77e7a4688886467574045c6a0126d140644b97e12b644d473e71d3e9324e44",
			&Asset{PolicyId: [28]byte{119, 231, 164, 104, 136, 134, 70, 117, 116, 4, 92, 106, 1, 38, 209, 64, 100, 75, 151, 225, 43, 100, 77, 71, 62, 113, 211, 233}, Name: NewAssetName([]byte("2ND"))},
			false,
		},
		{
			"case 3",
			"6b8d07d69639e9413dd637a1a815a7323c69c86abbafb66dbfdb1aa7",
			&Asset{PolicyId: [28]byte{107, 141, 7, 214, 150, 57, 233, 65, 61, 214, 55, 161, 168, 21, 167, 50, 60, 105, 200, 106, 187, 175, 182, 109, 191, 219, 26, 167}, Name: NewAssetName(nil)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAssetId(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssetIdByRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssetIdByRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssetName_String(t *testing.T) {
	type fields struct {
		data   [32]byte
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case 1", fields{[32]byte{72, 69, 76, 76, 79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 5}, "HELLO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AssetName{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("AssetName.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
