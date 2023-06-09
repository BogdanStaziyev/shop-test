// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// ValidateRequest provides a mock function with given fields: r, i
func (_m *Validator) ValidateRequest(r *http.Request, i interface{}) error {
	ret := _m.Called(r, i)

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Request, interface{}) error); ok {
		r0 = rf(r, i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidator creates a new instance of Validator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidator(t mockConstructorTestingTNewValidator) *Validator {
	mock := &Validator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
