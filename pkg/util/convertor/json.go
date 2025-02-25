package convertor

import (
	"encoding/json"
	"errors"
	cError "github.com/mhthrh/GoNest/model/error"
)

type Json struct {
}

func (j Json) Marshal(v interface{}) (string, error) {
	switch v.(type) {
	case cError.XError:
		jsonBytes, err := json.Marshal(v)
		return string(jsonBytes), err
	}

	return "", errors.New("marshal Error")
}

func (j Json) Unmarshal(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}
