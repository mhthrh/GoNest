package city

import (
	"github.com/google/uuid"
	cErrModel "github.com/mhthrh/common-lib/model/error"
)

type ICity interface {
	Load() *cErrModel.XError
	Cities() ([]City, *cErrModel.XError)
	GetByCountry(country string) ([]City, *cErrModel.XError)
	GetByCity(city string) ([]City, *cErrModel.XError)
	GetByCityAndCountry(city string, country string) (*City, *cErrModel.XError)
}
type City struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"countryCode"`
}
