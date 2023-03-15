package entity

import (
	"time"

	// external package for work with money
	"github.com/shopspring/decimal"
)

// Product describes the products sold in the online store.
type Product struct {
	ID          int64
	Name        string
	Description string
	Price       decimal.Decimal
	SalesmanID  int64
	Orders      []Order
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
