package city

import "github.com/google/uuid"

type City struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CountryCode string    `json:"countryCode"`
}
