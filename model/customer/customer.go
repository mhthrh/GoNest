package customer

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	id        uuid.UUID
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
	LastLogin time.Time `json:"lastLogin"`
	Stat      Status    `json:"status"`
}
