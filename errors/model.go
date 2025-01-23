package errors

type XError struct {
	Code          string  `json:"code"`
	Type          string  `json:"type"`
	Message       string  `json:"message"`
	Details       string  `json:"detail"`
	InternalError *XError `json:"internalError"`
	Time          string  `json:"time"`
}

func ConverError(err error) *XError {
	return &XError{}
}
