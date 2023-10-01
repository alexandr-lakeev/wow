// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ReadWriter is an autogenerated mock type for the ReadWriter type
type ReadWriter struct {
	mock.Mock
}

type ReadWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *ReadWriter) EXPECT() *ReadWriter_Expecter {
	return &ReadWriter_Expecter{mock: &_m.Mock}
}

// Read provides a mock function with given fields: p
func (_m *ReadWriter) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadWriter_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type ReadWriter_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - p []byte
func (_e *ReadWriter_Expecter) Read(p interface{}) *ReadWriter_Read_Call {
	return &ReadWriter_Read_Call{Call: _e.mock.On("Read", p)}
}

func (_c *ReadWriter_Read_Call) Run(run func(p []byte)) *ReadWriter_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *ReadWriter_Read_Call) Return(n int, err error) *ReadWriter_Read_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *ReadWriter_Read_Call) RunAndReturn(run func([]byte) (int, error)) *ReadWriter_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: p
func (_m *ReadWriter) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type ReadWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - p []byte
func (_e *ReadWriter_Expecter) Write(p interface{}) *ReadWriter_Write_Call {
	return &ReadWriter_Write_Call{Call: _e.mock.On("Write", p)}
}

func (_c *ReadWriter_Write_Call) Run(run func(p []byte)) *ReadWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *ReadWriter_Write_Call) Return(n int, err error) *ReadWriter_Write_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *ReadWriter_Write_Call) RunAndReturn(run func([]byte) (int, error)) *ReadWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewReadWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewReadWriter creates a new instance of ReadWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReadWriter(t mockConstructorTestingTNewReadWriter) *ReadWriter {
	mock := &ReadWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
