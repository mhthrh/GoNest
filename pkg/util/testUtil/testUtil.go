package testUtil

import (
	"runtime"
)

func IsTest() bool {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return fn != nil && fn.Name() == "testing.tRunner"
}
