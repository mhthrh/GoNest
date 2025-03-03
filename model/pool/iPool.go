package pool

import (
	"context"
	"github.com/google/uuid"
	customModelError "github.com/mhthrh/GoNest/model/error"
)

type IConnection interface {
	Maker(context.Context, <-chan Request, chan<- Response)
	Manager(context.Context, <-chan ManageRequest, chan<- *Connection)
	Refresh(context.Context, chan struct{}, chan<- RefreshResponse)
	Release(context.Context, chan ReleaseRequest, chan *customModelError.XError)
	ReleaseAll(bool) *customModelError.XError
}

// Connection If the value is nil, this method retrieves a connection from the pool.If a valid connection is provided, it returns the connection to the pool.
type Connection struct {
	Id    uuid.UUID
	Cnn   interface{}
	InUse bool
	Err   *customModelError.XError
}

// Request If count is 0, this method queries the current number of connections. Otherwise, it updates the connection count accordingly.
type Request struct {
	Count uint
	Type  Types
}
type Response struct {
	Total uint
	InUse uint
	Error *customModelError.XError
}
type ReleaseRequest struct {
	ID    uuid.UUID
	Force bool
}
type ManageRequest struct {
	Command Commands
	ID      uuid.UUID
}

type RefreshResponse struct {
	KilledCount uint
	TotalCount  int
}
