package requests

type RegisterCustomer struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Name     string `json:"name" validate:"required,gte=3"`
	Phone    string `json:"phone" validate:"required,gte=10"`
}
