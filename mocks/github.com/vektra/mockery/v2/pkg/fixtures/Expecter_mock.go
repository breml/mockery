// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Expecter is an autogenerated mock type for the Expecter type
type Expecter struct {
	mock.Mock
}

type Expecter_Expecter struct {
	mock *mock.Mock
}

func (_m *Expecter) EXPECT() *Expecter_Expecter {
	return &Expecter_Expecter{mock: &_m.Mock}
}

// ManyArgsReturns provides a mock function with given fields: str, i
func (_m *Expecter) ManyArgsReturns(str string, i int) ([]string, error) {
	ret := _m.Called(str, i)

	if len(ret) == 0 {
		panic("no return value specified for ManyArgsReturns")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]string, error)); ok {
		return rf(str, i)
	}
	if rf, ok := ret.Get(0).(func(string, int) []string); ok {
		r0 = rf(str, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(str, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Expecter_ManyArgsReturns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ManyArgsReturns'
type Expecter_ManyArgsReturns_Call struct {
	*mock.Call
}

// ManyArgsReturns is a helper method to define mock.On call
//   - str string
//   - i int
func (_e *Expecter_Expecter) ManyArgsReturns(str interface{}, i interface{}) *Expecter_ManyArgsReturns_Call {
	return &Expecter_ManyArgsReturns_Call{Call: _e.mock.On("ManyArgsReturns", str, i)}
}

func (_c *Expecter_ManyArgsReturns_Call) Run(run func(str string, i int)) *Expecter_ManyArgsReturns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *Expecter_ManyArgsReturns_Call) Return(strs []string, err error) *Expecter_ManyArgsReturns_Call {
	_c.Call.Return(strs, err)
	return _c
}

func (_c *Expecter_ManyArgsReturns_Call) RunAndReturn(run func(string, int) ([]string, error)) *Expecter_ManyArgsReturns_Call {
	_c.Call.Return(run)
	return _c
}

// NoArg provides a mock function with no fields
func (_m *Expecter) NoArg() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NoArg")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Expecter_NoArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NoArg'
type Expecter_NoArg_Call struct {
	*mock.Call
}

// NoArg is a helper method to define mock.On call
func (_e *Expecter_Expecter) NoArg() *Expecter_NoArg_Call {
	return &Expecter_NoArg_Call{Call: _e.mock.On("NoArg")}
}

func (_c *Expecter_NoArg_Call) Run(run func()) *Expecter_NoArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Expecter_NoArg_Call) Return(_a0 string) *Expecter_NoArg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Expecter_NoArg_Call) RunAndReturn(run func() string) *Expecter_NoArg_Call {
	_c.Call.Return(run)
	return _c
}

// NoReturn provides a mock function with given fields: str
func (_m *Expecter) NoReturn(str string) {
	_m.Called(str)
}

// Expecter_NoReturn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NoReturn'
type Expecter_NoReturn_Call struct {
	*mock.Call
}

// NoReturn is a helper method to define mock.On call
//   - str string
func (_e *Expecter_Expecter) NoReturn(str interface{}) *Expecter_NoReturn_Call {
	return &Expecter_NoReturn_Call{Call: _e.mock.On("NoReturn", str)}
}

func (_c *Expecter_NoReturn_Call) Run(run func(str string)) *Expecter_NoReturn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Expecter_NoReturn_Call) Return() *Expecter_NoReturn_Call {
	_c.Call.Return()
	return _c
}

func (_c *Expecter_NoReturn_Call) RunAndReturn(run func(string)) *Expecter_NoReturn_Call {
	_c.Run(run)
	return _c
}

// Variadic provides a mock function with given fields: ints
func (_m *Expecter) Variadic(ints ...int) error {
	_va := make([]interface{}, len(ints))
	for _i := range ints {
		_va[_i] = ints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Variadic")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...int) error); ok {
		r0 = rf(ints...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Expecter_Variadic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Variadic'
type Expecter_Variadic_Call struct {
	*mock.Call
}

// Variadic is a helper method to define mock.On call
//   - ints ...int
func (_e *Expecter_Expecter) Variadic(ints ...interface{}) *Expecter_Variadic_Call {
	return &Expecter_Variadic_Call{Call: _e.mock.On("Variadic",
		append([]interface{}{}, ints...)...)}
}

func (_c *Expecter_Variadic_Call) Run(run func(ints ...int)) *Expecter_Variadic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Expecter_Variadic_Call) Return(_a0 error) *Expecter_Variadic_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Expecter_Variadic_Call) RunAndReturn(run func(...int) error) *Expecter_Variadic_Call {
	_c.Call.Return(run)
	return _c
}

// VariadicMany provides a mock function with given fields: i, a, intfs
func (_m *Expecter) VariadicMany(i int, a string, intfs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, i, a)
	_ca = append(_ca, intfs...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VariadicMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, ...interface{}) error); ok {
		r0 = rf(i, a, intfs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Expecter_VariadicMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VariadicMany'
type Expecter_VariadicMany_Call struct {
	*mock.Call
}

// VariadicMany is a helper method to define mock.On call
//   - i int
//   - a string
//   - intfs ...interface{}
func (_e *Expecter_Expecter) VariadicMany(i interface{}, a interface{}, intfs ...interface{}) *Expecter_VariadicMany_Call {
	return &Expecter_VariadicMany_Call{Call: _e.mock.On("VariadicMany",
		append([]interface{}{i, a}, intfs...)...)}
}

func (_c *Expecter_VariadicMany_Call) Run(run func(i int, a string, intfs ...interface{})) *Expecter_VariadicMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(int), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Expecter_VariadicMany_Call) Return(_a0 error) *Expecter_VariadicMany_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Expecter_VariadicMany_Call) RunAndReturn(run func(int, string, ...interface{}) error) *Expecter_VariadicMany_Call {
	_c.Call.Return(run)
	return _c
}

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
