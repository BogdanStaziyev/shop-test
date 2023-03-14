package entity

import (
	"time"

	// external package for work with money
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       decimal.Decimal
	// TODO change to salesman id
	Salesman    Salesman
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
