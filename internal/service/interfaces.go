package service

import "github.com/BogdanStaziyev/shop-test/internal/entity"

// Databases structure that includes all databases
type Databases struct {
	Customer CustomerRepo
}

// Services structure that includes all services
type Services struct {
	Customer CustomerService
}

type CustomerService interface {
	// Save receives the entity.Customer from the handler.
	// Then hashes the password and passes the structure with the hash instead of the password to the repository.
	// After processing, simply returns the id and error from db
	Save(customer entity.Customer) (int64, error)
	// FindByID receives the customer structure from the handler.
	// Passes it to the repository and after receiving the entity.Customer and error, returns they.
	FindByID(id int64) (entity.Customer, error)
}

type CustomerRepo interface {
	// Create saved the entity.Customer to the database and returns the id and error
	Create(customer entity.Customer) (int64, error)
	// GetByID returns the entity.Customer and error
	GetByID(id int64) (entity.Customer, error)
}
