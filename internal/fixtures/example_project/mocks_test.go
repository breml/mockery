// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// TEST MOCKERY BOILERPLATE

package example_project

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/vektra/mockery/v3/internal/fixtures/example_project/foo"
)

// NewMockRoot creates a new instance of MockRoot. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRoot(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRoot {
	mock := &MockRoot{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockRoot is an autogenerated mock type for the Root type
type MockRoot struct {
	mock.Mock
}

type MockRoot_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRoot) EXPECT() *MockRoot_Expecter {
	return &MockRoot_Expecter{mock: &_m.Mock}
}

// ReturnsFoo provides a mock function for the type MockRoot
func (_mock *MockRoot) ReturnsFoo() (foo.Foo, error) {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnsFoo")
	}

	var r0 foo.Foo
	var r1 error
	if returnFunc, ok := ret.Get(0).(func() (foo.Foo, error)); ok {
		return returnFunc()
	}
	if returnFunc, ok := ret.Get(0).(func() foo.Foo); ok {
		r0 = returnFunc()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(foo.Foo)
		}
	}
	if returnFunc, ok := ret.Get(1).(func() error); ok {
		r1 = returnFunc()
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockRoot_ReturnsFoo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnsFoo'
type MockRoot_ReturnsFoo_Call struct {
	*mock.Call
}

// ReturnsFoo is a helper method to define mock.On call
func (_e *MockRoot_Expecter) ReturnsFoo() *MockRoot_ReturnsFoo_Call {
	return &MockRoot_ReturnsFoo_Call{Call: _e.mock.On("ReturnsFoo")}
}

func (_c *MockRoot_ReturnsFoo_Call) Run(run func()) *MockRoot_ReturnsFoo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRoot_ReturnsFoo_Call) Return(foo1 foo.Foo, err error) *MockRoot_ReturnsFoo_Call {
	_c.Call.Return(foo1, err)
	return _c
}

func (_c *MockRoot_ReturnsFoo_Call) RunAndReturn(run func() (foo.Foo, error)) *MockRoot_ReturnsFoo_Call {
	_c.Call.Return(run)
	return _c
}

// TakesBaz provides a mock function for the type MockRoot
func (_mock *MockRoot) TakesBaz(baz *foo.Baz) {
	_mock.Called(baz)
	return
}

// MockRoot_TakesBaz_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TakesBaz'
type MockRoot_TakesBaz_Call struct {
	*mock.Call
}

// TakesBaz is a helper method to define mock.On call
//   - baz
func (_e *MockRoot_Expecter) TakesBaz(baz interface{}) *MockRoot_TakesBaz_Call {
	return &MockRoot_TakesBaz_Call{Call: _e.mock.On("TakesBaz", baz)}
}

func (_c *MockRoot_TakesBaz_Call) Run(run func(baz *foo.Baz)) *MockRoot_TakesBaz_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*foo.Baz))
	})
	return _c
}

func (_c *MockRoot_TakesBaz_Call) Return() *MockRoot_TakesBaz_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRoot_TakesBaz_Call) RunAndReturn(run func(baz *foo.Baz)) *MockRoot_TakesBaz_Call {
	_c.Run(run)
	return _c
}

// NewMockStringer creates a new instance of MockStringer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStringer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStringer {
	mock := &MockStringer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockStringer is an autogenerated mock type for the Stringer type
type MockStringer struct {
	mock.Mock
}

type MockStringer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStringer) EXPECT() *MockStringer_Expecter {
	return &MockStringer_Expecter{mock: &_m.Mock}
}

// String provides a mock function for the type MockStringer
func (_mock *MockStringer) String() string {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if returnFunc, ok := ret.Get(0).(func() string); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(string)
	}
	return r0
}

// MockStringer_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockStringer_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockStringer_Expecter) String() *MockStringer_String_Call {
	return &MockStringer_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockStringer_String_Call) Run(run func()) *MockStringer_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStringer_String_Call) Return(s string) *MockStringer_String_Call {
	_c.Call.Return(s)
	return _c
}

func (_c *MockStringer_String_Call) RunAndReturn(run func() string) *MockStringer_String_Call {
	_c.Call.Return(run)
	return _c
}
