// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewUnsafeInterface creates a new instance of UnsafeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeInterface {
	mock := &UnsafeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// UnsafeInterface is an autogenerated mock type for the UnsafeInterface type
type UnsafeInterface struct {
	mock.Mock
}

// Do provides a mock function for the type UnsafeInterface
func (_mock *UnsafeInterface) Do(ptr *Pointer) {
	_mock.Called(ptr)
	return
}

type UnsafeInterface_expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeInterface) EXPECT() *UnsafeInterface_expecter {
	return &UnsafeInterface_expecter{mock: &_m.Mock}
}
