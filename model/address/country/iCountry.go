package country

import (
	cErrModel "github.com/mhthrh/common-lib/errors"
)

type ICountry interface {
	Load() *cErrModel.XError
	Countries() ([]Country, *cErrModel.XError)
	GetByName(country string) (*Country, *cErrModel.XError)
	GetByCode(city string) (*Country, *cErrModel.XError)
}
type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
