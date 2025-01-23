package test

import (
	"github.com/mhthrh/common-lib/errors"
)

type Test struct {
	Name     string
	Input    any
	OutPut   any
	HasError bool
	Err      *errors.XError
}
