package config

import (
	"github.com/mhthrh/common-lib/errors"
	"time"
)

func FileParameter(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "CONFIG100100",
		Type:          "config Error",
		Message:       "config file parameter error",
		Details:       "config file parameter error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileInitializerError(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "CONFIG100101",
		Type:          "config Error",
		Message:       "cannot initialize file parameter",
		Details:       "cannot initialize file parameter",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
