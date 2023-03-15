package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates the request to a data structure depending on the tags
func (v *Validator) Validate(r *http.Request, i interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		return err
	}
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
