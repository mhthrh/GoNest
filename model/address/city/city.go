package city

import (
	"github.com/google/uuid"
	cErrModel "github.com/mhthrh/common-lib/errors"
)

type ICity interface {
	Load() *cErrModel.XError
	Cities() []City
	GetByCountry(country string) ([]City, *cErrModel.XError)
	GetByCity(city string) ([]City, *cErrModel.XError)
}
type City struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"countryCode"`
}
