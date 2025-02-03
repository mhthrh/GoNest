package pool

import (
	"github.com/mhthrh/common-lib/errors"
	"time"
)

func DatabaseUnreachable(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100100",
		Type:          "dataBase Error",
		Message:       "Host is not reachable",
		Details:       "host is not reachable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func DbCnnNotExist(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100101",
		Type:          "dataBase Error",
		Message:       "connection not exist",
		Details:       "connection not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func DbConnectionFailed(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100102",
		Type:          "dataBase Error",
		Message:       "failed to connect to database",
		Details:       "failed to connect to database",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func ConnectionInUse(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100103",
		Type:          "dataBase Error",
		Message:       "connection already in use",
		Details:       "connection already in use",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func ReleaseAllError(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100104",
		Type:          "dataBase Error",
		Message:       "cannot release all connection in certain state",
		Details:       "cannot release all connection in certain state",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func InputParamsMismatch(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100105",
		Type:          "Pool Error",
		Message:       "input params mismatch",
		Details:       "input params mismatch, check new method parameter(S)",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func InputParamsTypeMismatch(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100106",
		Type:          "Pool Error",
		Message:       "input params type mismatch",
		Details:       "type mismatch, check new method parameter(S)",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func MaximumConnection(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100107",
		Type:          "Pool Error",
		Message:       "maximum connection exceeded",
		Details:       "maximum connection exceeded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func ConnectionTypeNotAcceptable(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100108",
		Type:          "Pool Error",
		Message:       "connection type not acceptable",
		Details:       "connection type not acceptable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func StopSignal(e *errors.XError) *errors.XError {
	return &errors.XError{
		Code:          "DB100109",
		Type:          "Pool Error",
		Message:       "stop signal detected",
		Details:       "stop signal detected",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
