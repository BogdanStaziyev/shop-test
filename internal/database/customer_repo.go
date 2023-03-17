package database

import (
	"context"
	"fmt"
	"log"
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
		return id, fmt.Errorf("customer repository Create error scan: %w", err)
	}
	return id, nil
}

func (c *customerRepo) GetByID(id int64) (customer entity.Customer, err error) {
	sql := `SELECT id, phone, password, name, email FROM customers WHERE id=$1 AND deleted_date IS NULL`
	err = c.Pool.QueryRow(context.Background(), sql, id).Scan(&customer.ID, &customer.Phone, &customer.Password, &customer.Name, &customer.Email)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("customer repository GetByID error scan: %w", err)
	}
	return customer, nil
}

func (c *customerRepo) Delete(id int64) error {
	sql := `UPDATE customers SET deleted_date = COALESCE(deleted_date, now()) WHERE id = $1 AND deleted_date IS NULL`
	result, err := c.Pool.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("customer repository Delete error: %w", err)
	}
	rows := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("customer repository Delete error: no rows in result set")
	}
	return nil
}

func (c *customerRepo) Update(customer entity.Customer) error {
	sql := `UPDATE customers SET name = $1, email = $2, phone = $3, password = $4, updated_date = $5 WHERE id = $6 AND deleted_date IS NULL`
	result, err := c.Pool.Exec(context.Background(), sql, customer.Name, customer.Email, &customer.Phone, customer.Password, time.Now(), customer.ID)
	if err != nil {
		return fmt.Errorf("customer repository Update error: %w", err)
	}
	rows := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("customer repository Update error: no rows in result set")
	}
	log.Println(rows)
	return nil
}
