package postgres_test

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	cMError "github.com/mhthrh/GoNest/model/error"
	l "github.com/mhthrh/GoNest/model/loader"
	cPool "github.com/mhthrh/GoNest/model/pool"
	"github.com/mhthrh/GoNest/model/test"
	"github.com/mhthrh/GoNest/pkg/loader/file"
	. "github.com/mhthrh/GoNest/pkg/pool/postgres"
	"testing"
)

var (
	f     l.IConfig
	c, c1 *l.Config
)

func init() {
	f = file.New("GoNest/config/file", "config-test.json")
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
			Err:      cPool.InputParamsMismatch(nil),
		}, {
			Name:     "test-2",
			Input:    c.DataBase,
			OutPut:   nil,
			HasError: false,
			Err:      nil,
		},
	}

	for _, tst := range tests {
		_, e := New(tst.Input.(l.DB))
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
		req chan cPool.Request
		res chan cPool.Response
	}
	input := Input{
		req: make(chan cPool.Request),
		res: make(chan cPool.Response),
	}

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	tests := []cPool.Request{{
		Count: 10,
		Type:  cPool.Types(0),
		Stop:  false,
	}, {
		Count: 10,
		Type:  cPool.Types(1),
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
				if r.Error.Code != cPool.ConnectionTypeNotAcceptable(nil).Code {
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
				if r.Error.Code != cPool.StopSignal(nil).Code {
					t.Errorf("expected stop signal but got %v", r.Error)
				}
			}

		}
	}
}

func TestManager(t *testing.T) {
	req := make(chan cPool.Request)
	res := make(chan cPool.Response)
	var id uuid.UUID
	request := make(chan cPool.ManageRequest)
	response := make(chan *cPool.Connection)

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(req, res)
	req <- cPool.Request{
		Count: 10,
		Type:  cPool.Types(1),
		Stop:  false,
	}
	result := <-res
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}

	tests := []cPool.ManageRequest{
		{
			Command: cPool.Commands(0),
			ID:      uuid.UUID{},
		}, {
			Command: cPool.Commands(1),
			ID:      uuid.UUID{},
		}, {
			Command: cPool.Commands(2),
			ID:      uuid.New(),
		}, {
			Command: cPool.Commands(2),
			ID:      uuid.UUID{},
		},
	}
	go p.Manager(request, response)
	for i, tst := range tests {
		if i == 3 {
			tst.ID = id
		}
		request <- cPool.ManageRequest{
			Command: tst.Command,
			ID:      tst.ID,
		}
		r := <-response
		switch i {
		case 0:
			if r.Err == nil {
				t.Error(fmt.Errorf("expected error but got nil"))
			}
			if r.Err.Code != cPool.CommandNotExist(nil).Code {
				t.Errorf("expected %v but got %v", cPool.CommandNotExist(nil), r.Err)
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
			if r.Err.Code != cPool.DbCnnNotExist(nil).Code {
				t.Error(fmt.Errorf("expected %v but got %v", cPool.DbCnnNotExist(nil), r.Err))
			}
		case 3:
			if r.Err.Code != cMError.Success().Code {
				t.Error(fmt.Errorf("expected %v but got %v", cMError.Success(), r.Err))
			}

		}
	}
}

func TestRefresh(t *testing.T) {
	req := make(chan struct{})
	res := make(chan cPool.RefreshResponse)
	request := make(chan cPool.Request)
	response := make(chan cPool.Response)

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(request, response)
	request <- cPool.Request{
		Count: 10,
		Type:  cPool.Types(1),
		Stop:  false,
	}
	result := <-response
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}

	go p.Refresh(req, res)
	req <- struct{}{}
	r := <-res
	if r.KilledCount != 0 {
		t.Error(fmt.Errorf("expected %v but got %v", cMError.Success(), 0))
	}
}

func TestRelease(t *testing.T) {
	req := make(chan cPool.ReleaseRequest)
	res := make(chan *cMError.XError)

	manageRequest := make(chan cPool.ManageRequest)
	manageResponse := make(chan *cPool.Connection)

	request := make(chan cPool.Request)
	response := make(chan cPool.Response)

	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(request, response)
	request <- cPool.Request{
		Count: 10,
		Type:  cPool.Types(1),
		Stop:  false,
	}
	result := <-response
	if result.Error != nil {
		t.Error(fmt.Errorf("expected no error but got %v", result.Error))
	}

	go p.Manager(manageRequest, manageResponse)
	manageRequest <- cPool.ManageRequest{
		Command: cPool.Commands(1),
		ID:      uuid.UUID{},
	}
	mRes := <-manageResponse

	if mRes.Err != nil {
		t.Error(fmt.Errorf("expected no error but got %v", mRes.Err))
	}
	go p.Release(req, res)

	tests := []cPool.ReleaseRequest{
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
				t.Error(fmt.Errorf("expected %v but got nil", cPool.DbCnnNotExist(nil)))
			}
		case 1:
			if r.Code != cMError.Success().Code {
				t.Error(fmt.Errorf("expected %v but got %v", cMError.Success(), r.Code))
			}
		case 2:
			if r.Code != cMError.Success().Code {
				t.Error(fmt.Errorf("expected %v but got %v", cMError.Success(), r.Code))
			}
		}

	}
}
func TestReleaseAll(t *testing.T) {
	request := make(chan cPool.Request)
	response := make(chan cPool.Response)
	p, err := New(c.DataBase)
	if err != nil {
		t.Error(err)
	}
	go p.Maker(request, response)
	request <- cPool.Request{
		Count: 10,
		Type:  cPool.Types(1),
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
