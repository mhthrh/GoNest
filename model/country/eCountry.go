package country

import (
	"fmt"
	error2 "github.com/mhthrh/common-lib/model/error"
	"time"
)

func FileUnreachable(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "CNTRY100100",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileEmpty(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "CNTRY100101",
		Type:          "Country Error",
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotLoaded(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "CNTRY100102",
		Type:          "Country Error",
		Message:       "country not loaded",
		Details:       "country not loaded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotFound(e *error2.XError, country string) *error2.XError {
	return &error2.XError{
		Code:          "CNTRY100102",
		Type:          "Country Error",
		Message:       "country not found",
		Details:       fmt.Sprintf("country with code/name: %s not found", country),
		InternalError: e,
		Time:          time.Now().String(),
	}
}
