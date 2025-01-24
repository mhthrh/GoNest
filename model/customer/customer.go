package customer

import (
	"github.com/google/uuid"
	"github.com/mhthrh/common-lib/model/address"
	"time"
)

type Customer struct {
	ID         uuid.UUID       `json:"id"`
	CustomerID string          `json:"customerID"`
	IdType     IdTypes         `json:"idType"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Email      string          `json:"email"`
	Mobile     string          `json:"mobile"`
	Address    address.Address `json:"address"`
	FirstName  string          `json:"firstName"`
	MiddleName string          `json:"middleName"`
	LastName   string          `json:"lastName"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	Stat       Status          `json:"status"`
	Picture    string          `json:"picture"`
	Document   string          `json:"document"`
}
