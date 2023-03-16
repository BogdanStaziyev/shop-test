package entity

import (
	"time"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/responses"
)

// Customer structure describes a customer in the online store
type Customer struct {
	ID          int64
	Email       string
	Password    string
	Name        string
	Phone       string
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

// Response return response customer structure from entity customer
func (c Customer) Response() responses.Customer {
	return responses.Customer{
		ID:    c.ID,
		Email: c.Email,
		Name:  c.Name,
		Phone: c.Phone,
	}
}
