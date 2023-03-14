package entity

import "time"

type Customer struct {
	ID          int64
	Email       string
	Password    string
	Name        string
	Phone       string
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
