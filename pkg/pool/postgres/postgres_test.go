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
	type Input struct {
		req chan pool.Request
		res chan pool.Response
	}
	input := Input{
		req: make(chan pool.Request),
		res: make(chan pool.Response),
	}

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	tests := []pool.Request{{
		Count: 10,
		Type:  pool.CTypes(0),
		Stop:  false,
	}, {
		Count: 10,
		Type:  pool.CTypes(1),
		Stop:  false,
	}, {
		Count: 0,
		Type:  0,
		Stop:  true,
	},
	}
	go p.Maker(input.req, input.res)

	for i, tst := range tests {
		input.req <- tst
		select {
		case r := <-input.res:
			switch i {
			case 0:
				if r.Error == nil {
					t.Error(fmt.Errorf("expected no error but got %v", r.Error))
				}
				if r.Error.Code != customeError.ConnectionTypeNotAcceptable(nil).Code {
					t.Errorf("expected stop signal but got %v", r.Error)
				}
			case 1:
				if r.Error != nil {
					t.Error(fmt.Errorf("expected no error but got %v", r.Error))
				}
				if r.Total != tst.Count {
					t.Error(fmt.Errorf("expected %v got %v", tst.Count, r.Total))
				}
			case 2:
				if r.Error == nil {
					t.Error(fmt.Errorf("expercted %v got no error", r.Error))
				}
				if r.Error.Code != customeError.StopSignal(nil).Code {
					t.Errorf("expected stop signal but got %v", r.Error)
				}
			}

		}
	}
}
