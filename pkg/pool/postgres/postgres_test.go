package postgres

import (
	"fmt"
	l "github.com/mhthrh/common-lib/config/loader"
	"github.com/mhthrh/common-lib/config/loader/file"
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/errors"
	customeError "github.com/mhthrh/common-lib/errors/pool"
	"github.com/mhthrh/common-lib/model/test"
	"testing"
)

var (
	f l.IConfig
	e *customModelError.XError
	c *l.Config
)

func TestNew(t *testing.T) {
	f, e = file.New("common-lib/config/file", "config-test.json")
	c, e = f.Initialize()
	c1, e := f.Initialize()
	if e != nil {
		t.Error("Failed to initialize config")
	}
	c1.DataBase.Host = ""
	tests := []test.Test{
		{
			Name:     "test-1",
			Input:    c1.DataBase,
			OutPut:   nil,
			HasError: true,
			Err:      customeError.InputParamsMismatch(nil),
		}, {
			Name:     "test-2",
			Input:    c.DataBase,
			OutPut:   nil,
			HasError: false,
			Err:      nil,
		},
	}

	for _, tst := range tests {
		_, e := New(tst.Input.(model.DB))
		if tst.HasError {
			if e == nil {
				t.Error(fmt.Errorf("expected error but got nil"))
				break
			}
			if e.Code != tst.Err.Code {
				t.Error(fmt.Errorf("expected error code %v but got %v", tst.Err.Code, e.Code))
			}

		}
		if e != nil && !tst.HasError {
			t.Error(fmt.Errorf("expected nil but got %v", e))
		}
	}

}
