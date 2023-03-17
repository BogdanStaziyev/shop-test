package service

import (
	"errors"
	"github.com/BogdanStaziyev/shop-test/internal/entity"
	"github.com/BogdanStaziyev/shop-test/internal/service/mocks"
	"github.com/BogdanStaziyev/shop-test/pkg/passwords"
	pMock "github.com/BogdanStaziyev/shop-test/pkg/passwords/passwordMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_customerService_Create(t *testing.T) {
	id := int64(1)
	customerExample := entity.Customer{
		Email:    "www@www.com",
		Password: "1234567890",
		Name:     "user user",
		Phone:    "1234567890",
	}

	tests := []struct {
		name        string
		customer    entity.Customer
		repo        func(customer entity.Customer) CustomerRepo
		passwordGen func(password string) passwords.Generator
		wantId      int64
		wantErr     bool
	}{
		{
			name:     "success created",
			customer: customerExample,
			repo: func(c entity.Customer) CustomerRepo {
				c.Password = "0000000000"
				mock := mocks.NewCustomerRepo(t)
				mock.On("Create", c).
					Return(id, nil).Times(1)
				return mock
			},
			passwordGen: func(password string) passwords.Generator {
				mock := pMock.NewGenerator(t)
				mock.On("GeneratePasswordHash", password).
					Return("0000000000", nil).Times(1)
				return mock
			},
			wantId:  id,
			wantErr: false,
		},
		{
			name:     "error password created",
			customer: customerExample,
			repo: func(c entity.Customer) CustomerRepo {
				mock := mocks.NewCustomerRepo(t)
				return mock
			},
			passwordGen: func(password string) passwords.Generator {
				mock := pMock.NewGenerator(t)
				mock.On("GeneratePasswordHash", password).
					Return("", errors.New("")).Times(1)
				return mock
			},
			wantId:  0,
			wantErr: true,
		},
		{
			name:     "success created",
			customer: customerExample,
			repo: func(c entity.Customer) CustomerRepo {
				c.Password = "0000000000"
				mock := mocks.NewCustomerRepo(t)
				mock.On("Create", c).
					Return(int64(0), errors.New("")).Times(1)
				return mock
			},
			passwordGen: func(password string) passwords.Generator {
				mock := pMock.NewGenerator(t)
				mock.On("GeneratePasswordHash", password).
					Return("0000000000", nil).Times(1)
				return mock
			},
			wantId:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := customerService{
				pg: tt.passwordGen(tt.customer.Password),
				cr: tt.repo(tt.customer),
			}
			resID, err := NewCustomerService(serviceMock.cr, serviceMock.pg).Create(tt.customer)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, resID, tt.wantId)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, resID, tt.wantId)
			}
		})
	}
}

func Test_customerService_Delete(t *testing.T) {
	tests := []struct {
		name    string
		id      int64
		repo    func(customerID int64) CustomerRepo
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			repo: func(customerID int64) CustomerRepo {
				mock := mocks.NewCustomerRepo(t)
				mock.On("Delete", customerID).
					Return(nil).Times(1)
				return mock
			},
			wantErr: false,
		},
		{
			name: "error from db",
			id:   1,
			repo: func(customerID int64) CustomerRepo {
				mock := mocks.NewCustomerRepo(t)
				mock.On("Delete", customerID).
					Return(errors.New("")).Times(1)
				return mock
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := customerService{
				pg: nil,
				cr: tt.repo(tt.id),
			}
			err := NewCustomerService(serviceMock.cr, serviceMock.pg).Delete(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_customerService_FindByID(t *testing.T) {
	customerExample := entity.Customer{
		Email:    "www@www.com",
		Password: "1234567890",
		Name:     "user user",
		Phone:    "1234567890",
	}

	tests := []struct {
		name    string
		id      int64
		repo    func(customerID int64) CustomerRepo
		result  entity.Customer
		wantErr bool
	}{
		{
			name: "success",
			id:   1,
			repo: func(customerID int64) CustomerRepo {
				mock := mocks.NewCustomerRepo(t)
				mock.On("GetByID", customerID).
					Return(customerExample, nil).Times(1)
				return mock
			},
			result:  customerExample,
			wantErr: false,
		},
		{
			name: "error from db",
			id:   1,
			repo: func(customerID int64) CustomerRepo {
				mock := mocks.NewCustomerRepo(t)
				mock.On("GetByID", customerID).
					Return(entity.Customer{}, errors.New("")).Times(1)
				return mock
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := customerService{
				pg: nil,
				cr: tt.repo(tt.id),
			}
			res, err := NewCustomerService(serviceMock.cr, serviceMock.pg).FindByID(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, res, tt.result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, res, tt.result)
			}
		})
	}
}

func Test_customerService_Update(t *testing.T) {
	customerExample := entity.Customer{
		Email:    "www@www.com",
		Password: "1234567890",
		Name:     "user user",
		Phone:    "1234567890",
	}

	tests := []struct {
		name        string
		customer    entity.Customer
		id          int64
		repo        func(customer entity.Customer) CustomerRepo
		passwordGen func(password string) passwords.Generator
		wantErr     bool
	}{
		{
			name:     "success created",
			customer: customerExample,
			repo: func(c entity.Customer) CustomerRepo {
				c.Password = "0000000000"
				mock := mocks.NewCustomerRepo(t)
				mock.On("Update", c).
					Return(nil).Times(1)
				return mock
			},
			passwordGen: func(password string) passwords.Generator {
				mock := pMock.NewGenerator(t)
				mock.On("GeneratePasswordHash", password).
					Return("0000000000", nil).Times(1)
				return mock
			},
			wantErr: false,
		},
		{
			name:     "error password created",
			customer: customerExample,
			repo: func(c entity.Customer) CustomerRepo {
				mock := mocks.NewCustomerRepo(t)
				return mock
			},
			passwordGen: func(password string) passwords.Generator {
				mock := pMock.NewGenerator(t)
				mock.On("GeneratePasswordHash", password).
					Return("", errors.New("")).Times(1)
				return mock
			},
			wantErr: true,
		},
		{
			name:     "success created",
			customer: customerExample,
			repo: func(c entity.Customer) CustomerRepo {
				c.Password = "0000000000"
				mock := mocks.NewCustomerRepo(t)
				mock.On("Update", c).
					Return(errors.New("")).Times(1)
				return mock
			},
			passwordGen: func(password string) passwords.Generator {
				mock := pMock.NewGenerator(t)
				mock.On("GeneratePasswordHash", password).
					Return("0000000000", nil).Times(1)
				return mock
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := customerService{
				pg: tt.passwordGen(tt.customer.Password),
				cr: tt.repo(tt.customer),
			}
			err := NewCustomerService(serviceMock.cr, serviceMock.pg).Update(tt.id, tt.customer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
