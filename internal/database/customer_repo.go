package database

import (
	"context"
	"fmt"
	"time"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/entity"
	"github.com/BogdanStaziyev/shop-test/internal/service"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/postgres"
)

type customerRepo struct {
	*postgres.Postgres
}

var _ service.CustomerRepo = (*customerRepo)(nil)

func NewCustomerRepo(cr *postgres.Postgres) *customerRepo {
	return &customerRepo{cr}
}

func (c *customerRepo) Create(cus entity.Customer) (id int64, err error) {
	sql := `INSERT INTO customers (phone, password, name, email, created_date, updated_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = c.Pool.QueryRow(context.Background(), sql, cus.Phone, cus.Password, cus.Name, cus.Email, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return
	}
	return id, nil
}

func (c *customerRepo) GetByID(id int64) (customer entity.Customer, err error) {
	sql := `SELECT id, phone, password, name, email FROM customers WHERE id=$1 AND deleted_date IS NULL`
	err = c.Pool.QueryRow(context.TODO(), sql, id).Scan(&customer.ID, &customer.Phone, &customer.Password, &customer.Name, &customer.Email)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("customer repository error scan: %w", err)
	}
	return customer, nil
}
