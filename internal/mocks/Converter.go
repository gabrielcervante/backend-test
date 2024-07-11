// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Converter is an autogenerated mock type for the Converter type
type Converter struct {
	mock.Mock
}

// ConvertCSVToSQL provides a mock function with given fields: _a0, _a1
func (_m *Converter) ConvertCSVToSQL(_a0 context.Context, _a1 string) {
	_m.Called(_a0, _a1)
}

// NewConverter creates a new instance of Converter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConverter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Converter {
	mock := &Converter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}