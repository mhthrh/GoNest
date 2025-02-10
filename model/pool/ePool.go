package pool

import (
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func DatabaseUnreachable(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100100",
		Type:          cError.Types(9),
		Message:       "Host is not reachable",
		Details:       "host is not reachable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func DbCnnNotExist(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100101",
		Type:          cError.Types(9),
		Message:       "connection not exist",
		Details:       "connection not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func DbConnectionFailed(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100102",
		Type:          cError.Types(9),
		Message:       "failed to connect to database",
		Details:       "failed to connect to database",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func ConnectionInUse(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100103",
		Type:          cError.Types(9),
		Message:       "connection already in use",
		Details:       "connection already in use",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func ReleaseAllError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100104",
		Type:          cError.Types(9),
		Message:       "cannot release all connection in certain state",
		Details:       "cannot release all connection in certain state",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func InputParamsMismatch(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100105",
		Type:          cError.Types(9),
		Message:       "input params mismatch",
		Details:       "input params mismatch, check new method parameter(S)",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func InputParamsTypeMismatch(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100106",
		Type:          cError.Types(9),
		Message:       "input params type mismatch",
		Details:       "type mismatch, check new method parameter(S)",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func MaximumConnection(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100107",
		Type:          cError.Types(9),
		Message:       "maximum connection exceeded",
		Details:       "maximum connection exceeded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func ConnectionTypeNotAcceptable(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100108",
		Type:          cError.Types(9),
		Message:       "connection type not acceptable",
		Details:       "connection type not acceptable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func StopSignal(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB100109",
		Type:          cError.Types(9),
		Message:       "stop signal detected",
		Details:       "stop signal detected",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func FreeConnectionNotExist(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB1001010",
		Type:          cError.Types(9),
		Message:       "free connection not exist",
		Details:       "free connection not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func CommandNotExist(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB1001011",
		Type:          cError.Types(9),
		Message:       "command not exist",
		Details:       "command not exist",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func TimeOut(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB1001012",
		Type:          cError.Types(9),
		Message:       "timeout exceeded",
		Details:       "timeout exceeded",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func TerminateByMain(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "DB1001013",
		Type:          cError.Types(9),
		Message:       "terminate by main",
		Details:       "terminate by main",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
