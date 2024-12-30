// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	"io"

	mock "github.com/stretchr/testify/mock"
)

// NewRequesterVariadicOneArgument creates a new instance of RequesterVariadicOneArgument. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterVariadicOneArgument(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterVariadicOneArgument {
	mock := &RequesterVariadicOneArgument{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// RequesterVariadicOneArgument is an autogenerated mock type for the RequesterVariadic type
type RequesterVariadicOneArgument struct {
	mock.Mock
}

// Get provides a mock function for the type RequesterVariadicOneArgument
func (_mock *RequesterVariadicOneArgument) Get(values ...string) bool {
	var tmpRet mock.Arguments
	if len(values) > 0 {
		tmpRet = _mock.Called(values)
	} else {
		tmpRet = _mock.Called()
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 bool
	if returnFunc, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = returnFunc(values...)
	} else {
		r0 = ret.Get(0).(bool)
	}
	return r0
}

// MultiWriteToFile provides a mock function for the type RequesterVariadicOneArgument
func (_mock *RequesterVariadicOneArgument) MultiWriteToFile(filename string, w ...io.Writer) string {
	var tmpRet mock.Arguments
	if len(w) > 0 {
		tmpRet = _mock.Called(filename, w)
	} else {
		tmpRet = _mock.Called(filename)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for MultiWriteToFile")
	}

	var r0 string
	if returnFunc, ok := ret.Get(0).(func(string, ...io.Writer) string); ok {
		r0 = returnFunc(filename, w...)
	} else {
		r0 = ret.Get(0).(string)
	}
	return r0
}

// OneInterface provides a mock function for the type RequesterVariadicOneArgument
func (_mock *RequesterVariadicOneArgument) OneInterface(a ...interface{}) bool {
	var tmpRet mock.Arguments
	if len(a) > 0 {
		tmpRet = _mock.Called(a)
	} else {
		tmpRet = _mock.Called()
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for OneInterface")
	}

	var r0 bool
	if returnFunc, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = returnFunc(a...)
	} else {
		r0 = ret.Get(0).(bool)
	}
	return r0
}

// Sprintf provides a mock function for the type RequesterVariadicOneArgument
func (_mock *RequesterVariadicOneArgument) Sprintf(format string, a ...interface{}) string {
	var tmpRet mock.Arguments
	if len(a) > 0 {
		tmpRet = _mock.Called(format, a)
	} else {
		tmpRet = _mock.Called(format)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Sprintf")
	}

	var r0 string
	if returnFunc, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = returnFunc(format, a...)
	} else {
		r0 = ret.Get(0).(string)
	}
	return r0
}

type RequesterVariadicOneArgument_expecter struct {
	mock *mock.Mock
}

func (_m *RequesterVariadicOneArgument) EXPECT() *RequesterVariadicOneArgument_expecter {
	return &RequesterVariadicOneArgument_expecter{mock: &_m.Mock}
}
