package v1

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/requests"
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/v1/mocks"
	"github.com/BogdanStaziyev/shop-test/internal/entity"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/logger"
	"github.com/BogdanStaziyev/shop-test/pkg/validators"
	validatorMock "github.com/BogdanStaziyev/shop-test/pkg/validators/validatorMock"

	// Test
	"github.com/stretchr/testify/require"
)

func Test_customerHandler_createCustomer(t *testing.T) {
	rek, _ := http.NewRequest("POST", "/customers/", nil)

	test := map[string]struct {
		httpRequest          *http.Request
		entityCustomer       entity.Customer
		requestCustomer      requests.RequestCustomer
		serviceBuilder       func(c entity.Customer) CustomerService
		validatorBuilder     func(r *http.Request, req *requests.RequestCustomer) validators.Validator
		expectedResponseBody string
		expectedStatusCode   int
	}{
		"success": {
			httpRequest:     rek,
			entityCustomer:  entity.Customer{},
			requestCustomer: requests.RequestCustomer{},
			serviceBuilder: func(c entity.Customer) CustomerService {
				mock := mocks.NewCustomerService(t)
				mock.On("Create", c).Return(int64(1), nil).Times(1)
				return mock
			},
			validatorBuilder: func(r *http.Request, c *requests.RequestCustomer) validators.Validator {
				mock := validatorMock.NewValidator(t)
				mock.On("ValidateRequest", r, c).Return(nil).Times(1)
				return mock
			},
			expectedResponseBody: `{"code":201,"response":"New customer created successfully, id: 1"}`,
			expectedStatusCode:   http.StatusCreated,
		},
		"error validator": {
			httpRequest:     rek,
			entityCustomer:  entity.Customer{},
			requestCustomer: requests.RequestCustomer{},
			serviceBuilder: func(c entity.Customer) CustomerService {
				mock := mocks.NewCustomerService(t)
				return mock
			},
			validatorBuilder: func(r *http.Request, c *requests.RequestCustomer) validators.Validator {
				mock := validatorMock.NewValidator(t)
				mock.On("ValidateRequest", r, c).Return(errors.New("")).Times(1)
				return mock
			},
			expectedResponseBody: `{"code":400, "error":"Could not validate customer data"}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		"error customer already exists": {
			httpRequest:     rek,
			entityCustomer:  entity.Customer{},
			requestCustomer: requests.RequestCustomer{},
			serviceBuilder: func(c entity.Customer) CustomerService {
				mock := mocks.NewCustomerService(t)
				mock.On("Create", c).Return(int64(0), errors.New("duplicate key value violates unique constraint")).Times(1)
				return mock
			},
			validatorBuilder: func(r *http.Request, c *requests.RequestCustomer) validators.Validator {
				mock := validatorMock.NewValidator(t)
				mock.On("ValidateRequest", r, c).Return(nil).Times(1)
				return mock
			},
			expectedResponseBody: `{"code":409, "error":"Could not save new customer, already exists"}`,
			expectedStatusCode:   http.StatusConflict,
		},
		"error internalServer": {
			httpRequest:     rek,
			entityCustomer:  entity.Customer{},
			requestCustomer: requests.RequestCustomer{},
			serviceBuilder: func(c entity.Customer) CustomerService {
				mock := mocks.NewCustomerService(t)
				mock.On("Create", c).Return(int64(0), errors.New("")).Times(1)
				return mock
			},
			validatorBuilder: func(r *http.Request, c *requests.RequestCustomer) validators.Validator {
				mock := validatorMock.NewValidator(t)
				mock.On("ValidateRequest", r, c).Return(nil).Times(1)
				return mock
			},
			expectedResponseBody: `{"code":500, "error":"Could not save new customer"}`,
			expectedStatusCode:   http.StatusInternalServerError,
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			ch := customerHandler{
				v:  tt.validatorBuilder(tt.httpRequest, &tt.requestCustomer),
				l:  logger.New(""),
				cs: tt.serviceBuilder(tt.entityCustomer),
			}
			response := httptest.NewRecorder()
			newCustomerHandler(chi.NewRouter(), ch.cs, ch.v, ch.l)
			ch.createCustomer().ServeHTTP(response, tt.httpRequest)
			require.Equal(t, tt.expectedStatusCode, response.Code)
			require.JSONEq(t, tt.expectedResponseBody, response.Body.String())
		})
	}
}

type Request struct {
	Method    string
	Url       string
	PathParam *PathParam
}

type PathParam struct {
	Name  string
	Value string
}

func Test_customerHandler_getByID(t *testing.T) {
	test := map[string]struct {
		serviceBuilder       func(id string) CustomerService
		idExample            string
		expectedResponseBody string
		expectedStatusCode   int
	}{
		"success": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				idInt64, _ := strconv.ParseInt(id, 10, 64)
				mock.On("FindByID", idInt64).Return(entity.Customer{}, nil).Times(1)
				return mock
			},
			idExample:            "1",
			expectedResponseBody: `{"code":200, "response": {"email":"", "id":0, "name":"", "phone":""}}`,
			expectedStatusCode:   http.StatusOK,
		},
		"error empty id": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				return mock
			},
			idExample:            "a",
			expectedResponseBody: `{"code":400, "error":"invalid request body"}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		"error customer not exist": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				idInt64, _ := strconv.ParseInt(id, 10, 64)
				mock.On("FindByID", idInt64).Return(entity.Customer{}, errors.New("no rows in result set")).Times(1)
				return mock
			},
			idExample:            "5",
			expectedResponseBody: `{"code":404, "error": "Could not find, customer not exists"}`,
			expectedStatusCode:   http.StatusNotFound,
		},
		"error internal server error": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				idInt64, _ := strconv.ParseInt(id, 10, 64)
				mock.On("FindByID", idInt64).Return(entity.Customer{}, errors.New("")).Times(1)
				return mock
			},
			idExample:            "5",
			expectedResponseBody: `{"code":500, "error": "Could not find customer"}`,
			expectedStatusCode:   http.StatusInternalServerError,
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", fmt.Sprintf("/customers/%s", tt.idExample), nil)
			response := httptest.NewRecorder()

			ch := customerHandler{
				v:  validators.NewValidator(),
				l:  logger.New(""),
				cs: tt.serviceBuilder(tt.idExample),
			}
			router := chi.NewRouter()
			router.Get("/customers/{id}", ch.getByID())
			router.ServeHTTP(response, req)
			require.Equal(t, tt.expectedStatusCode, response.Code)
			require.JSONEq(t, tt.expectedResponseBody, response.Body.String())
		})
	}
}

func Test_customerHandler_delete(t *testing.T) {
	test := map[string]struct {
		serviceBuilder       func(id string) CustomerService
		idExample            string
		expectedResponseBody string
		expectedStatusCode   int
	}{
		"success": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				idInt64, _ := strconv.ParseInt(id, 10, 64)
				mock.On("Delete", idInt64).Return(nil).Times(1)
				return mock
			},
			idExample:            "1",
			expectedResponseBody: `{"code":200, "message": "Customer deleted successfully"}`,
			expectedStatusCode:   http.StatusOK,
		},
		"error empty id": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				return mock
			},
			idExample:            "a",
			expectedResponseBody: `{"code":400, "error":"invalid request body"}`,
			expectedStatusCode:   http.StatusBadRequest,
		},
		"error customer not exist": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				idInt64, _ := strconv.ParseInt(id, 10, 64)
				mock.On("Delete", idInt64).Return(errors.New("no rows in result set")).Times(1)
				return mock
			},
			idExample:            "5",
			expectedResponseBody: `{"code":404, "error": "Could not delete, customer not exists"}`,
			expectedStatusCode:   http.StatusNotFound,
		},
		"error internal server error": {
			serviceBuilder: func(id string) CustomerService {
				mock := mocks.NewCustomerService(t)
				idInt64, _ := strconv.ParseInt(id, 10, 64)
				mock.On("Delete", idInt64).Return(errors.New("")).Times(1)
				return mock
			},
			idExample:            "5",
			expectedResponseBody: `{"code":500, "error": "Could not delete customer"}`,
			expectedStatusCode:   http.StatusInternalServerError,
		},
	}

	for name, tt := range test {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/customers/%s", tt.idExample), nil)
			response := httptest.NewRecorder()

			ch := customerHandler{
				v:  validators.NewValidator(),
				l:  logger.New(""),
				cs: tt.serviceBuilder(tt.idExample),
			}
			router := chi.NewRouter()
			router.Delete("/customers/{id}", ch.delete())
			router.ServeHTTP(response, req)
			require.Equal(t, tt.expectedStatusCode, response.Code)
			require.JSONEq(t, tt.expectedResponseBody, response.Body.String())
		})
	}
}
