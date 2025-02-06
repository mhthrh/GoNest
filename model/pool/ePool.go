package pool

import (
	error2 "github.com/mhthrh/common-lib/model/error"
	"time"
)

func DatabaseUnreachable(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100100",
		Type:          "dataBase Error",
		Message:       "Host is not reachable",
		Details:       "host is not reachable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func DbCnnNotExist(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100101",
		Type:          "dataBase Error",
		Message:       "connection not exist",
		Details:       "connection not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func DbConnectionFailed(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100102",
		Type:          "dataBase Error",
		Message:       "failed to connect to database",
		Details:       "failed to connect to database",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func ConnectionInUse(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100103",
		Type:          "dataBase Error",
		Message:       "connection already in use",
		Details:       "connection already in use",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func ReleaseAllError(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100104",
		Type:          "dataBase Error",
		Message:       "cannot release all connection in certain state",
		Details:       "cannot release all connection in certain state",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func InputParamsMismatch(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100105",
		Type:          "Pool Error",
		Message:       "input params mismatch",
		Details:       "input params mismatch, check new method parameter(S)",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func InputParamsTypeMismatch(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100106",
		Type:          "Pool Error",
		Message:       "input params type mismatch",
		Details:       "type mismatch, check new method parameter(S)",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func MaximumConnection(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100107",
		Type:          "Pool Error",
		Message:       "maximum connection exceeded",
		Details:       "maximum connection exceeded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func ConnectionTypeNotAcceptable(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100108",
		Type:          "Pool Error",
		Message:       "connection type not acceptable",
		Details:       "connection type not acceptable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func StopSignal(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB100109",
		Type:          "Pool Error",
		Message:       "stop signal detected",
		Details:       "stop signal detected",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FreeConnectionNotExist(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB1001010",
		Type:          "Pool Error",
		Message:       "free connection not exist",
		Details:       "free connection not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func CommandNotExist(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "DB1001011",
		Type:          "Pool Error",
		Message:       "command not exist",
		Details:       "command not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
