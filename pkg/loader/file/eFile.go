package file

import (
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func FileParameter(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100100",
		Type:          cError.Types(15),
		Message:       "config file parameter error",
		Details:       "config file parameter error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileInitializerError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100101",
		Type:          cError.Types(15),
		Message:       "cannot initialize file parameter",
		Details:       "cannot initialize file parameter",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
