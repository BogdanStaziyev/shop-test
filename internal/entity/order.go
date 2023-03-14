package entity

import "time"

type Order struct {
	ID          int64
	Customer    Customer
	Product     []Product
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
