package entity

import "time"

type Role string

const (
	ADMIN    Role = "admin"
	SALESMAN Role = "salesman"
)

// Salesman structure describes the sellers and admins in the online store.
type Salesman struct {
	ID          int64
	Email       string
	Password    string
	Name        string
	Phone       string
	Products    []Product
	Role        Role
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
