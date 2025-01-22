package errors

import "time"

func StreetEmpty() *XError {
	return &XError{
		Code:    "ADR100100",
		Type:    "Address Error",
		Message: "street name cannot be empty",
		Details: "street name cannot be empty",
		Time:    time.Now().String(),
	}
}
