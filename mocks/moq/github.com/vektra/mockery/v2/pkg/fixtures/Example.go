// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package test

import (
	"net/http"
	"sync"

	number_dir_http "github.com/vektra/mockery/v2/pkg/fixtures/12345678/http"
	my_http "github.com/vektra/mockery/v2/pkg/fixtures/http"
)

// ExampleMock is a mock implementation of test.Example.
//
//	func TestSomethingThatUsesExample(t *testing.T) {
//
//		// make and configure a mocked test.Example
//		mockedExample := &ExampleMock{
//			AFunc: func() http.Flusher {
//				panic("mock out the A method")
//			},
//			BFunc: func(fixtureshttp string) my_http.MyStruct {
//				panic("mock out the B method")
//			},
//			CFunc: func(fixtureshttp string) number_dir_http.MyStruct {
//				panic("mock out the C method")
//			},
//		}
//
//		// use mockedExample in code that requires test.Example
//		// and then make assertions.
//
//	}
type ExampleMock struct {
	// AFunc mocks the A method.
	AFunc func() http.Flusher

	// BFunc mocks the B method.
	BFunc func(fixtureshttp string) my_http.MyStruct

	// CFunc mocks the C method.
	CFunc func(fixtureshttp string) number_dir_http.MyStruct

	// calls tracks calls to the methods.
	calls struct {
		// A holds details about calls to the A method.
		A []struct {
		}
		// B holds details about calls to the B method.
		B []struct {
			// Fixtureshttp is the fixtureshttp argument value.
			Fixtureshttp string
		}
		// C holds details about calls to the C method.
		C []struct {
			// Fixtureshttp is the fixtureshttp argument value.
			Fixtureshttp string
		}
	}
	lockA sync.RWMutex
	lockB sync.RWMutex
	lockC sync.RWMutex
}

// A calls AFunc.
func (mock *ExampleMock) A() http.Flusher {
	if mock.AFunc == nil {
		panic("ExampleMock.AFunc: method is nil but Example.A was just called")
	}
	callInfo := struct {
	}{}
	mock.lockA.Lock()
	mock.calls.A = append(mock.calls.A, callInfo)
	mock.lockA.Unlock()
	return mock.AFunc()
}

// ACalls gets all the calls that were made to A.
// Check the length with:
//
//	len(mockedExample.ACalls())
func (mock *ExampleMock) ACalls() []struct {
} {
	var calls []struct {
	}
	mock.lockA.RLock()
	calls = mock.calls.A
	mock.lockA.RUnlock()
	return calls
}

// ResetACalls reset all the calls that were made to A.
func (mock *ExampleMock) ResetACalls() {
	mock.lockA.Lock()
	mock.calls.A = nil
	mock.lockA.Unlock()
}

// B calls BFunc.
func (mock *ExampleMock) B(fixtureshttp string) my_http.MyStruct {
	if mock.BFunc == nil {
		panic("ExampleMock.BFunc: method is nil but Example.B was just called")
	}
	callInfo := struct {
		Fixtureshttp string
	}{
		Fixtureshttp: fixtureshttp,
	}
	mock.lockB.Lock()
	mock.calls.B = append(mock.calls.B, callInfo)
	mock.lockB.Unlock()
	return mock.BFunc(fixtureshttp)
}

// BCalls gets all the calls that were made to B.
// Check the length with:
//
//	len(mockedExample.BCalls())
func (mock *ExampleMock) BCalls() []struct {
	Fixtureshttp string
} {
	var calls []struct {
		Fixtureshttp string
	}
	mock.lockB.RLock()
	calls = mock.calls.B
	mock.lockB.RUnlock()
	return calls
}

// ResetBCalls reset all the calls that were made to B.
func (mock *ExampleMock) ResetBCalls() {
	mock.lockB.Lock()
	mock.calls.B = nil
	mock.lockB.Unlock()
}

// C calls CFunc.
func (mock *ExampleMock) C(fixtureshttp string) number_dir_http.MyStruct {
	if mock.CFunc == nil {
		panic("ExampleMock.CFunc: method is nil but Example.C was just called")
	}
	callInfo := struct {
		Fixtureshttp string
	}{
		Fixtureshttp: fixtureshttp,
	}
	mock.lockC.Lock()
	mock.calls.C = append(mock.calls.C, callInfo)
	mock.lockC.Unlock()
	return mock.CFunc(fixtureshttp)
}

// CCalls gets all the calls that were made to C.
// Check the length with:
//
//	len(mockedExample.CCalls())
func (mock *ExampleMock) CCalls() []struct {
	Fixtureshttp string
} {
	var calls []struct {
		Fixtureshttp string
	}
	mock.lockC.RLock()
	calls = mock.calls.C
	mock.lockC.RUnlock()
	return calls
}

// ResetCCalls reset all the calls that were made to C.
func (mock *ExampleMock) ResetCCalls() {
	mock.lockC.Lock()
	mock.calls.C = nil
	mock.lockC.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *ExampleMock) ResetCalls() {
	mock.lockA.Lock()
	mock.calls.A = nil
	mock.lockA.Unlock()

	mock.lockB.Lock()
	mock.calls.B = nil
	mock.lockB.Unlock()

	mock.lockC.Lock()
	mock.calls.C = nil
	mock.lockC.Unlock()
}
