package postgres

import (
	l "github.com/mhthrh/common-lib/config/loader"
	"github.com/mhthrh/common-lib/pkg/util/pool"
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
	f = l.New("", "")
	c, e = f.Initialize()
	if e != nil {
		panic(e)
	}
	post = New(c.DataBase, 1, time.Second)
}

func TestConnection(t *testing.T) {
	post.Initialize()
	post.Get()
}
