package city

import (
	"github.com/mhthrh/common-lib/errors"
	"time"
)

func FileUnreachable(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "City100100",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileEmpty(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "City100101",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
