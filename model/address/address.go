package address

import (
	"github.com/mhthrh/common-lib/model/address/city"
	"github.com/mhthrh/common-lib/model/address/country"
)

type Address struct {
	Street     string          `json:"street"`
	City       city.City       `json:"city"`
	State      string          `json:"state"`
	PostalCode string          `json:"postalCode"`
	Country    country.Country `json:"address"`
}
