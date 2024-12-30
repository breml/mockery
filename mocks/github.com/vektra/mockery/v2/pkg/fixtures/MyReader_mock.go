// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewMyReader creates a new instance of MyReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMyReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MyReader {
	mock := &MyReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MyReader is an autogenerated mock type for the MyReader type
type MyReader struct {
	mock.Mock
}

// Read provides a mock function for the type MyReader
func (_mock *MyReader) Read(p []byte) (int, error) {
	ret := _mock.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 int
	var r1 error
	if returnFunc, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return returnFunc(p)
	}
	if returnFunc, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = returnFunc(p)
	} else {
		r0 = ret.Get(0).(int)
	}
	if returnFunc, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = returnFunc(p)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

type MyReader_expecter struct {
	mock *mock.Mock
}

func (_m *MyReader) EXPECT() *MyReader_expecter {
	return &MyReader_expecter{mock: &_m.Mock}
}
