package entity

import "time"

type Role string

const (
	ADMIN    Role = "admin"
	SALESMAN Role = "salesman"
)

type Salesman struct {
	ID          int64
	Email       string
	Password    string
	Name        string
	Phone       string
	Role        Role
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
