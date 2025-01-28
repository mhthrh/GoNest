package config

import (
	"github.com/mhthrh/common-lib/model/test"
	"testing"
)

func TestFile_Initialize(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "test-1",
			Input:    []string{"config.json", "/common-lib/config/file"},
			OutPut:   nil,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		f := NewFile(tst.Input.([]string)[0], tst.Input.([]string)[1])
		if err := f.Initialize(); err != nil {
			t.Error(err)
		}

	}
}

func TestFile_PrintConfig(t *testing.T) {

}

func TestVault_Initialize(t *testing.T) {

}

func TestVault_PrintConfig(t *testing.T) {

}
