package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	l "github.com/mhthrh/common-lib/config/loader"
	"github.com/mhthrh/common-lib/config/loader/file"
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/model/error"
	"github.com/mhthrh/common-lib/model/test"
	"github.com/mhthrh/common-lib/pkg/pool"
	"testing"
)

var (
	f     l.IConfig
	c, c1 *l.Config
)

func init() {
	f = file.New("common-lib/config/file", "config-test.json")
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
			Err:      pool.InputParamsMismatch(nil),
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
		Type:  pool.Types(0),
		Stop:  false,
	}, {
		Count: 10,
		Type:  pool.Types(1),
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
				if r.Error.Code != pool.ConnectionTypeNotAcceptable(nil).Code {
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
				if r.Error.Code != pool.StopSignal(nil).Code {
					t.Errorf("expected stop signal but got %v", r.Error)
				}
			}

		}
	}
}

func TestManager(t *testing.T) {
	req := make(chan pool.Request)
	res := make(chan pool.Response)
	var id uuid.UUID
	request := make(chan pool.ManageRequest)
	response := make(chan *pool.Connection)

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(req, res)
	req <- pool.Request{
		Count: 10,
		Type:  pool.Types(1),
		Stop:  false,
	}
	result := <-res
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}

	tests := []pool.ManageRequest{
		{
			Command: pool.Commands(0),
			ID:      uuid.UUID{},
		}, {
			Command: pool.Commands(1),
			ID:      uuid.UUID{},
		}, {
			Command: pool.Commands(2),
			ID:      uuid.New(),
		}, {
			Command: pool.Commands(2),
			ID:      uuid.UUID{},
		},
	}
	go p.Manager(request, response)
	for i, tst := range tests {
		if i == 3 {
			tst.ID = id
		}
		request <- pool.ManageRequest{
			Command: tst.Command,
			ID:      tst.ID,
		}
		r := <-response
		switch i {
		case 0:
			if r.Err == nil {
				t.Error(fmt.Errorf("expected error but got nil"))
			}
			if r.Err.Code != pool.CommandNotExist(nil).Code {
				t.Errorf("expected %v but got %v", pool.CommandNotExist(nil), r.Err)
			}
		case 1:
			if r.Err != nil {
				t.Error(fmt.Errorf("expected no error but got %v", r.Err))
			}
			if r.Cnn.(*sql.DB).Ping() != nil {
				t.Error(fmt.Errorf("expected open connection but got %v", r.Cnn.(*sql.DB).Ping()))
			}
			id = r.Id
		case 2:
			if r.Err.Code != pool.DbCnnNotExist(nil).Code {
				t.Error(fmt.Errorf("expected %v but got %v", pool.DbCnnNotExist(nil), r.Err))
			}
		case 3:
			if r.Err.Code != customModelError.Success().Code {
				t.Error(fmt.Errorf("expected %v but got %v", customModelError.Success(), r.Err))
			}

		}
	}
}

func TestRefresh(t *testing.T) {
	req := make(chan struct{})
	res := make(chan *customModelError.XError)
	request := make(chan pool.Request)
	response := make(chan pool.Response)

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(request, response)
	request <- pool.Request{
		Count: 10,
		Type:  pool.Types(1),
		Stop:  false,
	}
	result := <-response
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}

	go p.Refresh(req, res)
	req <- struct{}{}
	r := <-res
	if r.Code != customModelError.Success().Code {
		t.Error(fmt.Errorf("expected %v but got %v", customModelError.Success(), r.Code))
	}
}

func TestRelease(t *testing.T) {
	req := make(chan pool.ReleaseRequest)
	res := make(chan *customModelError.XError)

	manageRequest := make(chan pool.ManageRequest)
	manageResponse := make(chan *pool.Connection)

	request := make(chan pool.Request)
	response := make(chan pool.Response)

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(request, response)
	request <- pool.Request{
		Count: 10,
		Type:  pool.Types(1),
		Stop:  false,
	}
	result := <-response
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}

	go p.Manager(manageRequest, manageResponse)
	manageRequest <- pool.ManageRequest{
		Command: pool.Commands(1),
		ID:      uuid.UUID{},
	}
	mRes := <-manageResponse

	if mRes.Err != nil {
		t.Error(fmt.Errorf("expected no error but got %v", mRes.Err))
	}
	go p.Release(req, res)

	tests := []pool.ReleaseRequest{
		{
			ID:    uuid.New(),
			Force: false,
			Stop:  false,
		}, {
			ID:    mRes.Id,
			Force: false,
			Stop:  false,
		}, {
			ID:    mRes.Id,
			Force: false,
			Stop:  true,
		},
	}

	for i, tst := range tests {
		req <- tst
		r := <-res
		switch i {
		case 0:
			if r == nil {
				t.Error(fmt.Errorf("expected %v but got nil", pool.DbCnnNotExist(nil)))
			}
		case 1:
			if r.Code != customModelError.Success().Code {
				t.Error(fmt.Errorf("expected %v but got %v", customModelError.Success(), r.Code))
			}
		case 2:
			if r.Code != customModelError.Success().Code {
				t.Error(fmt.Errorf("expected %v but got %v", customModelError.Success(), r.Code))
			}
		}

	}
}
func TestReleaseAll(t *testing.T) {
	request := make(chan pool.Request)
	response := make(chan pool.Response)
	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(request, response)
	request <- pool.Request{
		Count: 10,
		Type:  pool.Types(1),
		Stop:  false,
	}
	result := <-response
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}
	err = p.ReleaseAll(true)
	if err != nil {
		t.Error(fmt.Errorf("expected no error but got %v", err))
	}
}
