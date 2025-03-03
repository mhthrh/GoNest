package grpc_test

import (
	"context"
	l "github.com/mhthrh/GoNest/model/loader"
	"github.com/mhthrh/GoNest/model/pool"
	"github.com/mhthrh/GoNest/pkg/loader/file"
	g "github.com/mhthrh/GoNest/pkg/pool/grpc"
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
func TestGrpc_Maker(t *testing.T) {
	ctx := context.Background()
	req := make(chan pool.Request)
	res := make(chan pool.Response)
	gRPC, err := g.NewGrpc(c.Grpcs)
	if err != nil {
		t.Error(err)
	}
	go gRPC.Maker(ctx, req, res)
	req <- pool.Request{
		Count: 10,
		Type:  pool.Types(9),
	}

	select {
	case r := <-res:
		if r.Error != nil {
			t.Error(r.Error)
		}
	}
	ctx.Done()
}
