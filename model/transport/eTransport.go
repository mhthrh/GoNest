package transport

import (
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func NotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "404",
		Type:          cError.Types(17),
		Message:       "address not found",
		Details:       "page not found",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
