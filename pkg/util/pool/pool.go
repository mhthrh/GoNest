package pool

import (
	"github.com/google/uuid"
)

type IPool interface {
	New(chan struct{}) chan Pool
	Release(chan Pool)
}

type Pool struct {
	CId uuid.UUID
	Typ CTypes
	Cnn *interface{}
	Err error
}
