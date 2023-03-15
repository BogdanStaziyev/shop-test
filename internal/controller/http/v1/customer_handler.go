package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/requests"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/logger"
	"github.com/BogdanStaziyev/shop-test/pkg/validators"
)

type customerHandler struct {
	v *validators.Validator
	l logger.Interface
}

func newCustomerHandler(r chi.Router, v *validators.Validator, l logger.Interface) {
	c := &customerHandler{
		v: v,
		l: l,
	}

	r.Route("/customer", func(r chi.Router) {
		r.Post(
			"/register",
			c.Register(),
		)
	})
}

func (c *customerHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer requests.RegisterCustomer
		if err := c.v.Validate(r, &customer); err != nil {
			c.l.Error("CustomerHandler validation error", "err", err)
		}
	}
}
