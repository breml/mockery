// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewGetInt creates a new instance of GetInt. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetInt(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetInt {
	mock := &GetInt{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// GetInt is an autogenerated mock type for the GetInt type
type GetInt struct {
	mock.Mock
}

// Get provides a mock function for the type GetInt
func (_mock *GetInt) Get() int {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 int
	if returnFunc, ok := ret.Get(0).(func() int); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(int)
	}
	return r0
}

type GetInt_expecter struct {
	mock *mock.Mock
}

func (_m *GetInt) EXPECT() *GetInt_expecter {
	return &GetInt_expecter{mock: &_m.Mock}
}
