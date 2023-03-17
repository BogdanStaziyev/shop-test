package v1

import (
	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/entity"
)

// Services structure that includes all services
type Services struct {
	Customer CustomerService
}

//go:generate mockery --dir . --name CustomerService --output ./mocks
type CustomerService interface {
	// Create receives the entity.Customer from the handler.
	// Then hashes the password and passes the structure with the hash instead of the password to the repository.
	// After processing, simply returns the id and error from db
	Create(customer entity.Customer) (int64, error)
	// FindByID receives the customer structure from the handler.
	// Passes it to the repository and after receiving the entity.Customer and error, returns they.
	FindByID(id int64) (entity.Customer, error)
	// Delete it simply passes the id to the repository without performing any actions.
	Delete(id int64) error
	// Update receives the new entity.Customer from the handler.
	// Then hashes the password and assigns the received id to the structure
	//Passes the structure with the hash instead of the password to the repository.
	Update(id int64, customer entity.Customer) error
}
