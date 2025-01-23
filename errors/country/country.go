package country

import (
	"fmt"
	"github.com/mhthrh/common-lib/errors"
	"time"
)

func FileUnreachable(e error) *errors.XError {
	return &errors.XError{
		Code:          "CNTRY100100",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: fmt.Sprintf("internal error: %v", e),
		Time:          time.Now().String(),
	}
}
func FileEmpty(e error) *errors.XError {
	return &errors.XError{
		Code:          "CNTRY100101",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: fmt.Sprintf("internal error: %v", e),
		Time:          time.Now().String(),
	}
}
