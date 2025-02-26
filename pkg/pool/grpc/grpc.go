package grpc

import (
	cError "github.com/mhthrh/GoNest/model/error"
	cPool "github.com/mhthrh/GoNest/model/pool"
)

type Grpc struct {
}

func NewGrpc() cPool.IConnection {
	return &Grpc{}
}
func (g Grpc) Maker(requests <-chan cPool.Request, responses chan<- cPool.Response) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) Manager(requests <-chan cPool.ManageRequest, connections chan<- *cPool.Connection) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) Refresh(c chan struct{}, responses chan<- cPool.RefreshResponse) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) Release(requests chan cPool.ReleaseRequest, errors chan *cError.XError) {
	//TODO implement me
	panic("implement me")
}

func (g Grpc) ReleaseAll(b bool) *cError.XError {
	//TODO implement me
	panic("implement me")
}
