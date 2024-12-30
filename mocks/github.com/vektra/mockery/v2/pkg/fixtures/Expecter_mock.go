// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewExpecter creates a new instance of Expecter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExpecter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Expecter {
	mock := &Expecter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Expecter is an autogenerated mock type for the Expecter type
type Expecter struct {
	mock.Mock
}

// ManyArgsReturns provides a mock function for the type Expecter
func (_mock *Expecter) ManyArgsReturns(str string, i int) ([]string, error) {
	ret := _mock.Called(str, i)

	if len(ret) == 0 {
		panic("no return value specified for ManyArgsReturns")
	}

	var r0 []string
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(string, int) ([]string, error)); ok {
		return returnFunc(str, i)
	}
	if returnFunc, ok := ret.Get(0).(func(string, int) []string); ok {
		r0 = returnFunc(str, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = returnFunc(str, i)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// NoArg provides a mock function for the type Expecter
func (_mock *Expecter) NoArg() string {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for NoArg")
	}

	var r0 string
	if returnFunc, ok := ret.Get(0).(func() string); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(string)
	}
	return r0
}

// NoReturn provides a mock function for the type Expecter
func (_mock *Expecter) NoReturn(str string) {
	_mock.Called(str)
	return
}

// Variadic provides a mock function for the type Expecter
func (_mock *Expecter) Variadic(ints ...int) error {
	// int
	_va := make([]interface{}, len(ints))
	for _i := range ints {
		_va[_i] = ints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _mock.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Variadic")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(...int) error); ok {
		r0 = returnFunc(ints...)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// VariadicMany provides a mock function for the type Expecter
func (_mock *Expecter) VariadicMany(i int, a string, intfs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, i, a)
	_ca = append(_ca, intfs...)
	ret := _mock.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VariadicMany")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(int, string, ...interface{}) error); ok {
		r0 = returnFunc(i, a, intfs...)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

type Expecter_expecter struct {
	mock *mock.Mock
}

func (_m *Expecter) EXPECT() *Expecter_expecter {
	return &Expecter_expecter{mock: &_m.Mock}
}
