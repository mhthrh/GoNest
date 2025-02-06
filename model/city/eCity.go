package city

import (
	"fmt"
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func FileUnreachable(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100100",
		Type:          cError.Types(11),
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileEmpty(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100101",
		Type:          cError.Types(11),
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotLoaded(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100102",
		Type:          cError.Types(11),
		Message:       "cities not loaded",
		Details:       "cities not loaded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotFound(e *cError.XError, city, country string) *cError.XError {
	return &cError.XError{
		Code:          "100102",
		Type:          cError.Types(11),
		Message:       "cities not found",
		Details:       fmt.Sprintf("cities %s not found in %s", city, country),
		InternalError: e,
		Time:          time.Now().String(),
	}
}
