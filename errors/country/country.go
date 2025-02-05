package country

import (
	"github.com/mhthrh/common-lib/errors"
	"time"
)

func FileUnreachable(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "CNTRY100100",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileEmpty(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "CNTRY100101",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotLoaded(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "CNTRY100102",
		Type:          "Country Error",
		Message:       "country not loaded",
		Details:       "country not loaded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
