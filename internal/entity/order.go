package entity

import "time"

// Order structure describes an order in the online store
type Order struct {
	ID          int64
	CustomerID  int64
	Product     []Product
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
