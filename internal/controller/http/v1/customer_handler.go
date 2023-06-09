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

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/logger"
	"github.com/BogdanStaziyev/shop-test/pkg/validators"
)

type customerHandler struct {
	v  validators.Validator
	l  logger.Interface
	cs CustomerService
}

func newCustomerHandler(r chi.Router, cs CustomerService, v validators.Validator, l logger.Interface) {
	c := &customerHandler{
		v:  v,
		l:  l,
		cs: cs,
	}

	r.Route("/customers", func(customer chi.Router) {
		customer.Post(
			"/",
			c.createCustomer(),
		)
		customer.Get(
			"/{id}",
			c.getByID(),
		)
		customer.Delete(
			"/{id}",
			c.delete(),
		)
		customer.Put(
			"/{id}",
			c.updateCustomer(),
		)
		customer.Handle("/*", NotFoundJSON())
	})
}

// createCustomer handler function for creating a new customer.
// Validates the incoming request, and sends the validated entity to the service.
func (c customerHandler) createCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer requests.RequestCustomer
		if err := c.v.ValidateRequest(r, &customer); err != nil {
			c.l.Error("CustomerHandler createCustomer validation error", "err", err)
			responses.ErrorResponse(w, http.StatusBadRequest, "Could not validate customer data")
			return
		}

		id, err := c.cs.Create(customer.RequestTOEntity())
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				c.l.Error("CustomerHandler createCustomer", "err", err)
				responses.ErrorResponse(w, http.StatusConflict, "Could not save new customer, already exists")
				return

			} else {
				c.l.Error("CustomerHandler createCustomer", "err", err)
				responses.ErrorResponse(w, http.StatusInternalServerError, "Could not save new customer")
				return
			}
		}

		responses.Response(w, http.StatusCreated, fmt.Sprintf("New customer created successfully, id: %d", id))
	}
}

// getByID function is a request handler for retrieving customer data by ID.
// Performing request validation and sending the validated entity to the service.
func (c customerHandler) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			c.l.Error("ID should be a number getByID", "have: ", id)
			responses.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
			return
		}

		customer, err := c.cs.FindByID(id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				c.l.Error("CustomerHandler getByID", "err", err)
				responses.ErrorResponse(w, http.StatusNotFound, "Could not find, customer not exists")
				return
			} else {
				c.l.Error("CustomerHandler getByID", "err", err)
				responses.ErrorResponse(w, http.StatusInternalServerError, "Could not find customer")
				return
			}
		}
		responses.Response(w, http.StatusOK, customer.Response())
	}
}

// delete HTTP handler for deleting a customer with a specific ID.
func (c customerHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			c.l.Error("ID should be a number CustomerHandler delete", "have: ", id)
			responses.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
			return
		}

		err = c.cs.Delete(id)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				c.l.Error("CustomerHandler delete", "err", err)
				responses.ErrorResponse(w, http.StatusNotFound, "Could not delete, customer not exists")
				return
			} else {
				c.l.Error("CustomerHandler delete", "err", err)
				responses.ErrorResponse(w, http.StatusInternalServerError, "Could not delete customer")
				return
			}
		}
		responses.MessageResponse(w, http.StatusOK, "Customer deleted successfully")
	}
}

// updateCustomer function handles requests to update a customer's data by their ID.
// Validating the request and sending the updated entity to the service.
func (c customerHandler) updateCustomer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer requests.RequestCustomer
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			c.l.Error("ID should be a number CustomerHandler updateCustomer", "have: ", id)
			responses.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
			return
		}
		if err = c.v.ValidateRequest(r, &customer); err != nil {
			c.l.Error("CustomerHandler updateCustomer validation error", "err", err)
			responses.ErrorResponse(w, http.StatusBadRequest, "Could not validate customer data")
			return
		}

		err = c.cs.Update(id, customer.RequestTOEntity())
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				c.l.Error("CustomerHandler updateCustomer", "err", err)
				responses.ErrorResponse(w, http.StatusNotFound, "Could not update, customer not exists")
				return

			} else {
				c.l.Error("CustomerHandler updateCustomer", "err", err)
				responses.ErrorResponse(w, http.StatusInternalServerError, "Could not update customer")
				return
			}
		}

		responses.MessageResponse(w, http.StatusOK, "Customer updated successfully")
	}
}
