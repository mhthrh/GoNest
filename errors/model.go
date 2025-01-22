package errors

type XError struct {
	Code    string `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Details string `json:"detail"`
	Time    string `json:"time"`
}
