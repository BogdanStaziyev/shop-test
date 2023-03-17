package service

import (
	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/entity"
)

// Databases structure that includes all databases
type Databases struct {
	Customer CustomerRepo
}

//go:generate mockery --dir . --name CustomerRepo --output ./mocks
type CustomerRepo interface {
	// Create saved the entity.Customer to the database and returns the id and error
	Create(customer entity.Customer) (int64, error)
	// GetByID returns the entity.Customer and error
	GetByID(id int64) (entity.Customer, error)
	// Delete searches in the database if a customer exists by id.
	// Sets the current time in the "deleted_date" field.
	Delete(id int64) error
	// Update searches in the database if a customer exists by id.
	// Sets the current time in the "updated_date" field.
	Update(customer entity.Customer) error
}
