package country

import (
	"fmt"
	cError "github.com/mhthrh/common-lib/model/error"
	"time"
)

func FileUnreachable(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100100",
		Type:          cError.Types(12),
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FileEmpty(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100101",
		Type:          cError.Types(12),
		Message:       "internal error",
		Details:       "internal error",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotLoaded(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100102",
		Type:          cError.Types(12),
		Message:       "country not loaded",
		Details:       "country not loaded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func NotFound(e *cError.XError, country string) *cError.XError {
	return &cError.XError{
		Code:          "100103",
		Type:          cError.Types(12),
		Message:       "country not found",
		Details:       fmt.Sprintf("country with code/name: %s not found", country),
		InternalError: e,
		Time:          time.Now().String(),
	}
}
