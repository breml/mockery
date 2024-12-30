// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewRequesterArgSameAsPkg creates a new instance of RequesterArgSameAsPkg. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterArgSameAsPkg(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterArgSameAsPkg {
	mock := &RequesterArgSameAsPkg{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// RequesterArgSameAsPkg is an autogenerated mock type for the RequesterArgSameAsPkg type
type RequesterArgSameAsPkg struct {
	mock.Mock
}

// Get provides a mock function for the type RequesterArgSameAsPkg
func (_mock *RequesterArgSameAsPkg) Get(test string) {
	_mock.Called(test)
	return
}

type RequesterArgSameAsPkg_expecter struct {
	mock *mock.Mock
}

func (_m *RequesterArgSameAsPkg) EXPECT() *RequesterArgSameAsPkg_expecter {
	return &RequesterArgSameAsPkg_expecter{mock: &_m.Mock}
}
