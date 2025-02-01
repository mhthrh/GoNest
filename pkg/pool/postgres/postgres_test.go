package postgres

import (
	"fmt"
	l "github.com/mhthrh/common-lib/config/loader"
	"github.com/mhthrh/common-lib/config/loader/file"
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/errors"
	customeError "github.com/mhthrh/common-lib/errors/pool"
	"github.com/mhthrh/common-lib/model/test"
	"github.com/mhthrh/common-lib/pkg/pool"
	"testing"
)

var (
	f     l.IConfig
	e     *customModelError.XError
	c, c1 *l.Config
)

func init() {
	f, e = file.New("common-lib/config/file", "config-test.json")
	c, _ = f.Initialize()
	c1, _ = f.Initialize()
	c1.DataBase.Host = ""
}
func TestNew(t *testing.T) {

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

func TestMaker(t *testing.T) {
	inp := struct {
		req chan pool.Request
		res chan pool.Response
	}
	{
		req: make(chan pool.Request),
		res: make(chan pool.Response),
	}

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(inp.req, inp.res)
	tests := []test.Test{
		{
			Name:     "test-1",
			Input:    inp,
			OutPut:   nil,
			HasError: true,
			Err:      customeError.ConnectionTypeNotAcceptable(nil),
		},
	}
	for _, tst := range tests {
		tst.Input.(inp)
		//tst.OutPut.(chan pool.Response) <- res
		fmt.Println(<-inp.res)
		if tst.HasError {
			if e == nil {
				t.Error(fmt.Errorf("expected error but got nil"))
			}
			if e.Code != tst.Err.Code {
				t.Error(fmt.Errorf("expected error code %v but got %v", tst.Err.Code, e.Code))
			}
		}
	}

}
