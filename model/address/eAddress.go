package address

import (
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

func StreetNotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "ADR100100",
		Type:          cError.Types(15),
		Message:       "street name cannot be empty",
		Details:       "street name cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func PostalCodeNotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "ADR100101",
		Type:          cError.Types(15),
		Message:       "postal code cannot be empty",
		Details:       "postal code cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func StateNotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "ADR100102",
		Type:          cError.Types(15),
		Message:       "state name cannot be empty",
		Details:       "state name cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func CityNotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "ADR100103",
		Type:          cError.Types(15),
		Message:       "city name cannot be empty",
		Details:       "city name cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func CountryNotFound(e *cError.XError) *cError.XError {
	return &cError.XError{
		Code:          "ADR100104",
		Type:          cError.Types(15),
		Message:       "address code cannot be empty",
		Details:       "address code cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
