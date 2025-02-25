package transport

import (
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func NotFoundError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "404",
		Type:          cError.Types(17),
		Message:       "Address not found",
		Details:       "Page not found",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func UnauthorizedError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "401",
		Type:          cError.Types(17),
		Message:       "Unauthorized",
		Details:       "Authentication required",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func ForbiddenError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "403",
		Type:          cError.Types(17),
		Message:       "Forbidden",
		Details:       "Access denied",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func InternalServerError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "500",
		Type:          cError.Types(17),
		Message:       "Internal Server Error",
		Details:       "Something went wrong",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func BadRequestError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "400",
		Type:          cError.Types(17),
		Message:       "Bad Request",
		Details:       "Invalid request parameters",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func ConflictError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "409",
		Type:          cError.Types(17),
		Message:       "Conflict",
		Details:       "Resource conflict detected",
		InternalError: e,
		Time:          time.Now().String(),
	}
}

func ServiceUnavailableError(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "503",
		Type:          cError.Types(17),
		Message:       "Service Unavailable",
		Details:       "Service is temporarily unavailable",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
