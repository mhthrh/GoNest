package customer

import (
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func NotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "100100",
		Type:          cError.Types(10),
		Message:       "maximum connection exceeded",
		Details:       "maximum connection exceeded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
