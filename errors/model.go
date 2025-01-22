package errors

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Details string `json:"detail"`
	Time    string `json:"time"`
}
