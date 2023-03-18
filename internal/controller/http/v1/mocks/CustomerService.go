// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/BogdanStaziyev/shop-test/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// CustomerService is an autogenerated mock type for the CustomerService type
type CustomerService struct {
	mock.Mock
}

// Create provides a mock function with given fields: customer
func (_m *CustomerService) Create(customer entity.Customer) (int64, error) {
	ret := _m.Called(customer)

	var r0 int64
	if rf, ok := ret.Get(0).(func(entity.Customer) int64); ok {
		r0 = rf(customer)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *CustomerService) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *CustomerService) FindByID(id int64) (entity.Customer, error) {
	ret := _m.Called(id)

	var r0 entity.Customer
	if rf, ok := ret.Get(0).(func(int64) entity.Customer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Customer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, customer
func (_m *CustomerService) Update(id int64, customer entity.Customer) error {
	ret := _m.Called(id, customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, entity.Customer) error); ok {
		r0 = rf(id, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCustomerService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerService creates a new instance of CustomerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerService(t mockConstructorTestingTNewCustomerService) *CustomerService {
	mock := &CustomerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}