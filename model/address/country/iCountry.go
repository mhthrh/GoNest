package country

import (
	cErrModel "github.com/mhthrh/common-lib/errors"
)

type ICountry interface {
	Load() ([]Country, *cErrModel.XError)
	GetByName(string) (Country, *cErrModel.XError)
	GetByCode(string) (Country, *cErrModel.XError)
}
type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
