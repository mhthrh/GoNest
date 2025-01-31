package postgres

import (
	l "github.com/mhthrh/common-lib/config/loader"
	customModelError "github.com/mhthrh/common-lib/errors"
	"github.com/mhthrh/common-lib/pkg/pool"
	"testing"
	"time"
)

var (
	f    l.IConfig
	post pool.IConnection
	e    error
	c    *l.Config
)

func init() {
	f = l.New("common-lib/config/file", "config-test.json")
	c, e = f.Initialize()
	if e != nil {
		panic(e)
	}
	post = New(c.DataBase, 1, time.Second)
}

func TestConnection(t *testing.T) {
	err := make(chan customModelError.XError)
	post.Initialize(err)
	for {
		select {
		case ee := <-err:
			t.Log(ee)
		}
	}
	post.Get()
}
