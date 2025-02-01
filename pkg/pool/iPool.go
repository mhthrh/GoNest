package pool

import (
	"github.com/google/uuid"
	customModelError "github.com/mhthrh/common-lib/errors"
)

type IConnection interface {
	Maker(chan Request, chan Response)
	Manager(chan Connection)
	Refresh(chan struct{}) chan Response
	Release(chan uuid.UUID) chan Response
	ReleaseAll() *customModelError.XError
}

// Connection If the value is nil, this method retrieves a connection from the pool.If a valid connection is provided, it returns the connection to the pool.
type Connection struct {
	Id    uuid.UUID
	Cnn   interface{}
	InUse bool
}

// Request If count is 0, this method queries the current number of connections. Otherwise, it updates the connection count accordingly.
type Request struct {
	Count uint
	Type  CTypes
	Stop  bool
}
type Response struct {
	Total uint
	InUse uint
	Error *customModelError.XError
}
