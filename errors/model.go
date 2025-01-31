package errors

import (
	"fmt"
	"runtime"
	"time"
)

type XError struct {
	Code          string  `json:"code"`
	Type          string  `json:"type"`
	Message       string  `json:"message"`
	Details       string  `json:"detail"`
	InternalError *XError `json:"internalError"`
	Time          string  `json:"time"`
}

func RunTimeError(err error) *XError {
	_, file, line, _ := runtime.Caller(1)
	return &XError{
		Code:          "runtime",
		Type:          "runtime",
		Message:       err.Error(),
		Details:       fmt.Sprintf("file name: %s, line number: %d", file, line),
		InternalError: nil,
		Time:          time.Now().String(),
	}
}
