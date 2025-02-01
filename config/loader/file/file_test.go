package file

import (
	customeError "github.com/mhthrh/common-lib/errors/config"
	"github.com/mhthrh/common-lib/model/test"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "TestNew-0",
			Input:    []string{"", ""},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.FileParameter(nil),
		}, {
			Name:     "TestNew-1",
			Input:    []string{"", "", ""},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.FileParameter(nil),
		}, {
			Name:     "TestNew-2",
			Input:    []string{"common-lib/config/file", "config-test.json"},
			OutPut:   nil,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		_, err := New(tst.Input.([]string)[0], tst.Input.([]string)[1])
		if tst.HasError {
			if err == nil {
				t.Errorf("TestNew-%s: expected error but got none", tst.Name)
				break
			}
			if err.Code != tst.Err.Code {
				t.Errorf("TestNew-%s: expected error: %v, got: %v", tst.Name, tst.Err, err)
			}
		}
		if err != nil && !tst.HasError {
			t.Errorf("TestNew-%s: expected nil, got: %v", tst.Name, err)
		}

	}
}

func TestInitialize(t *testing.T) {
	c, _ := New("common-lib/config/file", "config-test.json")
	tests := []test.Test{
		{
			Name:     "TestInitial-0",
			Input:    []string{"", ""},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.FileParameter(nil),
		}, {
			Name:     "TestInitial-1",
			Input:    []string{"common-lib/config/file", "config-test.json"},
			OutPut:   nil,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		cnfg, err := c.Initialize()
		if err != nil {
			t.Errorf("TestInitialize-%s: expected nil error, got: %v", tst.Name, err)
		}
		if cnfg.DataBase.Host != "localhost" {
			t.Errorf("TestInitialize-%s: expected nil, got %v", tst.Name, cnfg.DataBase.Host)
		}
	}

}
