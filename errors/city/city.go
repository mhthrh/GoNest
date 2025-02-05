package city

import (
	"fmt"
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
func NotLoaded(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "City100102",
		Type:          "City Error",
		Message:       "cities not loaded",
		Details:       "cities not loaded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotFound(e *errors.XError, city, country string) *errors.XError {
	return &errors.XError{
		Code:          "City100102",
		Type:          "Country Error",
		Message:       "cities not found",
		Details:       fmt.Sprintf("cities %s not found in %s", city, country),
		InternalError: e,
		Time:          time.Now().String(),
	}
}
