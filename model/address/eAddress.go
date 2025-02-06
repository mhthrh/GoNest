package address

import (
	error2 "github.com/mhthrh/common-lib/model/error"
	"time"
)

func StreetNotFound(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "ADR100100",
		Type:          "Address Error",
		Message:       "street name cannot be empty",
		Details:       "street name cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func PostalCodeNotFound(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "ADR100101",
		Type:          "Address Error",
		Message:       "postal code cannot be empty",
		Details:       "postal code cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func StateNotFound(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "ADR100102",
		Type:          "Address Error",
		Message:       "state name cannot be empty",
		Details:       "state name cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func CityNotFound(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "ADR100103",
		Type:          "Address Error",
		Message:       "city name cannot be empty",
		Details:       "city name cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
func CountryNotFound(e *error2.XError) *error2.XError {
	return &error2.XError{
		Code:          "ADR100104",
		Type:          "Address Error",
		Message:       "address code cannot be empty",
		Details:       "address code cannot be empty",
		InternalError: e,
		Time:          time.Now().String(),
	}
}
