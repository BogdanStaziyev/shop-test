package requests

import (

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/entity"
)

// RequestCustomer structure that should come in JSON format.
// All fields are mandatory for validation
type RequestCustomer struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Name     string `json:"name" validate:"required,gte=3"`
	Phone    string `json:"phone" validate:"required,gte=10"`
}

// RequestTOEntity converts the requests structure into an entity
func (r RequestCustomer) RequestTOEntity() entity.Customer {
	return entity.Customer{
		Email:    r.Email,
		Password: r.Password,
		Name:     r.Name,
		Phone:    r.Phone,
	}
}
