package pool

import (
	"github.com/google/uuid"
	customModelError "github.com/mhthrh/common-lib/errors"
)

type IPool interface {
	Get() (*Connection, *customModelError.XError)
	Put(key uuid.UUID) *customModelError.XError
	Initialize()
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
