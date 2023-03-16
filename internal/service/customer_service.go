package service

import (
	"fmt"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/entity"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/passwords"
)

type customerService struct {
	cr CustomerRepo
	pg passwords.Generator
}

var _ CustomerService = (*customerService)(nil)

func NewCustomerService(cr CustomerRepo, pg passwords.Generator) *customerService {
	return &customerService{
		cr: cr,
		pg: pg,
	}
}

func (c *customerService) Save(customer entity.Customer) (id int64, err error) {
	// Generates a password hash for storage in the database
	customer.Password, err = c.pg.GeneratePasswordHash(customer.Password)
	if err != nil {
		return id, fmt.Errorf("customerService Save customer, could not generate hash: %w", err)
	}
	id, err = c.cr.Create(customer)
	if err != nil {
		return id, fmt.Errorf("customerService Save customer, could not save customer: %w", err)
	}
	return id, nil
}

func (c *customerService) FindByID(id int64) (entity.Customer, error) {
	customer, err := c.cr.GetByID(id)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("customerService FindByID customer: %w", err)
	}
	return customer, nil
}
