package city

import (
	"fmt"
	"github.com/mhthrh/common-lib/errors"
	"time"
)

func FileUnreachable(e error) *errors.XError {
	return &errors.XError{
		Code:          "City100100",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: fmt.Sprintf("internal error: %v", e),
		Time:          time.Now().String(),
	}
}
func FileEmpty(e error) *errors.XError {
	return &errors.XError{
		Code:          "City100101",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: fmt.Sprintf("internal error: %v", e),
		Time:          time.Now().String(),
	}
}
