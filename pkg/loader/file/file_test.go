package file

import (
	"github.com/mhthrh/GoNest/model/test"
	"testing"
)

func TestInitialize(t *testing.T) {
	c := New("common-lib/config/file", "config-test.json")
	tests := []test.Test{
		{
			Name:     "TestInitial-0",
			Input:    []string{"", ""},
			OutPut:   nil,
			HasError: true,
			Err:      FileParameter(nil),
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
