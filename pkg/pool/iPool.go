package pool

import (
	"github.com/google/uuid"
	customModelError "github.com/mhthrh/common-lib/errors"
)

type IConnection interface {
	Initialize(chan customModelError.XError)
	Get() *Connection
	Put(key uuid.UUID) *customModelError.XError
	Release(uuid.UUID) *customModelError.XError
	ReleaseAll() *customModelError.XError
}

type Connection struct {
	CId   uuid.UUID
	Typ   CTypes
	Cnn   interface{}
	InUse bool
	Err   *error
}
