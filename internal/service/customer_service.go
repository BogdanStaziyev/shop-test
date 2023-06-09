package service

import (
	"fmt"
	"strings"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/entity"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/passwords"
)

type customerService struct {
	cr CustomerRepo
	pg passwords.Generator
}

func NewCustomerService(cr CustomerRepo, pg passwords.Generator) *customerService {
	return &customerService{
		cr: cr,
		pg: pg,
	}
}

func (c *customerService) Create(customer entity.Customer) (id int64, err error) {
	// Generates a password hash for storage in the database
	customer.Password, err = c.pg.GeneratePasswordHash(customer.Password)
	// Save Email in low register
	customer.Email = strings.ToLower(customer.Email)
	if err != nil {
		return id, fmt.Errorf("customerService Create customer, could not generate hash: %w", err)
	}
	id, err = c.cr.Create(customer)
	if err != nil {
		return id, fmt.Errorf("customerService Create customer, could not save customer: %w", err)
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

func (c *customerService) Delete(id int64) error {
	if err := c.cr.Delete(id); err != nil {
		return fmt.Errorf("customerService Delete customer: %w", err)
	}
	return nil
}

func (c *customerService) Update(id int64, customer entity.Customer) (err error) {
	// Generates a password hash for storage in the database
	customer.Password, err = c.pg.GeneratePasswordHash(customer.Password)
	if err != nil {
		return fmt.Errorf("customerService Update customer, could not generate hash: %w", err)
	}
	customer.ID = id
	// Update Email in low register
	customer.Email = strings.ToLower(customer.Email)
	err = c.cr.Update(customer)
	if err != nil {
		return fmt.Errorf("customerService Update customer: %w", err)
	}
	return nil
}
