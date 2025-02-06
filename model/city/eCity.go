package city

import (
	"fmt"
	error2 "github.com/mhthrh/common-lib/model/error"
	"time"
)

func FileUnreachable(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "City100100",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileEmpty(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "City100101",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotLoaded(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "City100102",
		Type:          "City Error",
		Message:       "cities not loaded",
		Details:       "cities not loaded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotFound(e *error2.XError, city, country string) *error2.XError {
	return &error2.XError{
		Code:          "City100102",
		Type:          "Country Error",
		Message:       "cities not found",
		Details:       fmt.Sprintf("cities %s not found in %s", city, country),
		InternalError: e,
		Time:          time.Now().String(),
	}
}
