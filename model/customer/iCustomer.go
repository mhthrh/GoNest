package customer

import (
	"github.com/google/uuid"
	"github.com/mhthrh/GoNest/model/address"
	cError "github.com/mhthrh/GoNest/model/error"
	"time"
)

type ICustomer interface {
	RegisterCustomer(address address.Address, customer Customer) *cError.XError
	GetCustomerById(id string) (*Customer, bool)
	GetCustomerByName(name string) (*Customer, bool)
	GetCustomerByEmail(email string) (*Customer, bool)
	GetCustomerByPhone(phone string) (*Customer, bool)
	ChangeStatus(id string, status Status) *cError.XError
	EditCustomer(address address.Address, customer Customer) *cError.XError
}
type Customer struct {
	ID         uuid.UUID       `json:"id"`
	CustomerID string          `json:"customerID"`
	IdType     Types           `json:"idType"`
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
