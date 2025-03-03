package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	cError "github.com/mhthrh/GoNest/model/error"
	"github.com/mhthrh/GoNest/model/loader"
	cPool "github.com/mhthrh/GoNest/model/pool"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var (
	connections map[string][]*cPool.Connection
	ins         cPool.IConnection
	once        sync.Once
)

func init() {
	connections = make(map[string][]*cPool.Connection)
}

type Grpc struct {
	grpcs []loader.Grpc
	m     *sync.Mutex
}

func NewGrpc(g []loader.Grpc) (cPool.IConnection, *cError.XError) {
	if len(g) == 0 {
		return nil, cPool.InputParamsMismatch(nil)
	}
	for _, v := range g {
		connections[v.Name] = make([]*cPool.Connection, v.PoolSize)
	}
	once.Do(func() {
		ins = Grpc{
			grpcs: g,
			m:     &sync.Mutex{},
		}
	})
	return ins, nil
}
func (g Grpc) Maker(ctx context.Context, requests <-chan cPool.Request, responses chan<- cPool.Response) {
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-requests:
			if r.Type != cPool.Types(9) {
				responses <- cPool.Response{
					Total: 0,
					InUse: 0,
					Error: cPool.ConnectionTypeNotAcceptable(nil),
				}
				continue
			}
			for _, v := range g.grpcs {
				for range v.PoolSize - len(connections) {
					c, err := newConnection(v)
					if err != nil {
						responses <- cPool.Response{
							Total: uint(len(connections)),
							InUse: uint(0),
							Error: err,
						}
						break
					}
					add(v.Name, c)
				}
			}
			//responses <- cPool.Response{
			//	Total: uint(len(connections)),
			//	InUse: uint(len(lo.PickBy(connections, func(key string, value cPool.Connection) bool {
			//		return value.InUse == true
			//	}))),
			//	Error: nil,
			//}
			continue
		}
	}
}

func (g Grpc) Manager(ctx context.Context, requests <-chan cPool.ManageRequest, connections chan<- *cPool.Connection) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) Refresh(ctx context.Context, c chan struct{}, responses chan<- cPool.RefreshResponse) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) Release(ctx context.Context, requests chan cPool.ReleaseRequest, errors chan *cError.XError) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) ReleaseAll(b bool) *cError.XError {
	//TODO implement me
	panic("implement me")
}

func newConnection(d loader.Grpc) (c *cPool.Connection, e *cError.XError) {
	cnn, err := grpc.NewClient(fmt.Sprintf("%s:%d", d.Ip, d.Port), grpc.WithInsecure())
	if err != nil {
		return nil, cPool.DbConnectionFailed(cError.RunTimeError(err))
	}

	key := uuid.New()
	c = &cPool.Connection{
		Id:    key,
		Cnn:   cnn,
		InUse: false,
	}
	return c, nil
}
func add(name string, cNns *cPool.Connection) {
	if cNns == nil {
		log.Println("object is empty")
		return
	}
	v, ok := connections[name]
	if !ok {
		log.Printf("canot find connection name %s\n", name)
		return
	}
	for i := 0; i < len(v); i++ {
		if v[i] == nil {
			v[i] = cNns
			return
		}
	}
}
