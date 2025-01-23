package errors

import "time"

func StreetNotFound() *XError {
	return &XError{
		Code:    "ADR100100",
		Type:    "Address Error",
		Message: "street name cannot be empty",
		Details: "street name cannot be empty",
		Time:    time.Now().String(),
	}
}
func PostalCodeNotFound() *XError {
	return &XError{
		Code:    "ADR100101",
		Type:    "Address Error",
		Message: "postal code cannot be empty",
		Details: "postal code cannot be empty",
		Time:    time.Now().String(),
	}
}
func StateNotFound() *XError {
	return &XError{
		Code:    "ADR100102",
		Type:    "Address Error",
		Message: "state name cannot be empty",
		Details: "state name cannot be empty",
	}
}
func CityNotFound() *XError {
	return &XError{
		Code:    "ADR100103",
		Type:    "Address Error",
		Message: "city name cannot be empty",
		Details: "",
		Time:    time.Now().String(),
	}
}
func CountryNotFound() *XError {
	return &XError{
		Code:    "ADR100104",
		Type:    "Address Error",
		Message: "country code cannot be empty",
		Details: "country code cannot be empty",
	}
}
