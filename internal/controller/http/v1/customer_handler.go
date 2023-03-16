package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/requests"
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/responses"
	"github.com/BogdanStaziyev/shop-test/internal/service"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/logger"
	"github.com/BogdanStaziyev/shop-test/pkg/validators"
)

type customerHandler struct {
	v  *validators.Validator
	l  logger.Interface
	cs service.CustomerService
}

func newCustomerHandler(r chi.Router, cs service.CustomerService, v *validators.Validator, l logger.Interface) {
	c := &customerHandler{
		v:  v,
		l:  l,
		cs: cs,
	}

	r.Route("/customer", func(customer chi.Router) {
		customer.Post(
			"/save",
			c.createCustomer(),
		)
		customer.Get(
			"/{id}",
			c.getByID(),
		)
		customer.Handle("/*", NotFoundJSON())
	})
}

// saveCustomer Admin can create a new customers
func (c customerHandler) createCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer requests.RegisterCustomer
		if err := c.v.Validate(r, &customer); err != nil {
			c.l.Error("CustomerHandler validation error", "err", err)
			responses.ErrorResponse(w, http.StatusBadRequest, "Could not validate customer data")
			return
		}

		id, err := c.cs.Save(customer.RequestTOEntity())
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				c.l.Error("CustomerHandler SaveCustomer", "err", err)
				responses.ErrorResponse(w, http.StatusNotFound, "Could not save new customer user already exists")
				return

			} else {
				c.l.Error("CustomerHandler SaveCustomer", "err", err)
				responses.ErrorResponse(w, http.StatusInternalServerError, "Could not save new customer")
				return
			}
		}

		responses.Response(w, http.StatusCreated, fmt.Sprintf("New customer created successfully, id: %d", id))
	}
}

// getByID only Admin can get customers
func (c customerHandler) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			c.l.Error("ID should be a number", "have: ", id)
			responses.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
			return
		}

		customer, err := c.cs.FindByID(id)
		if err != nil {
			c.l.Error("CustomerHandler GetByID", "err", err)
			responses.ErrorResponse(w, http.StatusInternalServerError, "Could not find customer")
			return
		}

		responses.Response(w, http.StatusOK, customer.Response())
	}
}
