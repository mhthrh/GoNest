package file

import (
	error2 "github.com/mhthrh/common-lib/model/error"
	"time"
)

func FileParameter(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "CONFIG100100",
		Type:          "config Error",
		Message:       "config file parameter error",
		Details:       "config file parameter error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileInitializerError(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "CONFIG100101",
		Type:          "config Error",
		Message:       "cannot initialize file parameter",
		Details:       "cannot initialize file parameter",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
