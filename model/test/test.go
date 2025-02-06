package test

import (
	error2 "github.com/mhthrh/GoNest/model/error"
)

type Test struct {
	Name     string
	Input    any
	OutPut   any
	HasError bool
	Err      *error2.XError
}
