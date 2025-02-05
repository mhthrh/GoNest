package country

import (
	cErrModel "github.com/mhthrh/common-lib/errors"
)

type ICountry interface {
	Load() []Country
	GetByName(string) Country
	GetByCode(string) (Country, *cErrModel.XError)
}
type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
