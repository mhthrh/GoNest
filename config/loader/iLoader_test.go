package loader

import (
	"github.com/mhthrh/common-lib/config/loader/file"
	"github.com/mhthrh/common-lib/model/test"
	"reflect"
	"testing"
)

func TestFileConfig_Initialize(t *testing.T) {
	tests := []test.Test{
		{
			Name: "Test-1",
			Input: file.FileConfig{
				path: "",
				Name: "",
			},
			OutPut:   Config{},
			HasError: false,
			Err:      nil,
		},
	}
	for _, tt := range tests {

		f := file.FileConfig{
			Path: tt.Input.(file.FileConfig).Path,
			name: tt.Input.(file.FileConfig).name,
		}
		got, err := f.Initialize()
		if (err != nil) != tt.HasError {
			t.Errorf("Initialize() error = %v, wantErr %v", err, tt.HasError)
			return
		}
		if !reflect.DeepEqual(got, tt.OutPut) {
			t.Errorf("Initialize() got = %v, want %v", got, tt.OutPut)
		}
	}

}

func TestFileConfig_PrintConfig(t *testing.T) {

}

func TestVaultConfig_Initialize(t *testing.T) {

}

func TestVaultConfig_PrintConfig(t *testing.T) {

}
