// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.
package api

import mock "github.com/stretchr/testify/mock"

// DisruptionKindMock is an autogenerated mock type for the DisruptionKind type
type DisruptionKindMock struct {
	mock.Mock
}

type DisruptionKindMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DisruptionKindMock) EXPECT() *DisruptionKindMock_Expecter {
	return &DisruptionKindMock_Expecter{mock: &_m.Mock}
}

// GenerateArgs provides a mock function with given fields:
func (_m *DisruptionKindMock) GenerateArgs() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateArgs")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// DisruptionKindMock_GenerateArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateArgs'
type DisruptionKindMock_GenerateArgs_Call struct {
	*mock.Call
}

// GenerateArgs is a helper method to define mock.On call
func (_e *DisruptionKindMock_Expecter) GenerateArgs() *DisruptionKindMock_GenerateArgs_Call {
	return &DisruptionKindMock_GenerateArgs_Call{Call: _e.mock.On("GenerateArgs")}
}

func (_c *DisruptionKindMock_GenerateArgs_Call) Run(run func()) *DisruptionKindMock_GenerateArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DisruptionKindMock_GenerateArgs_Call) Return(_a0 []string) *DisruptionKindMock_GenerateArgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DisruptionKindMock_GenerateArgs_Call) RunAndReturn(run func() []string) *DisruptionKindMock_GenerateArgs_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields:
func (_m *DisruptionKindMock) Validate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisruptionKindMock_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type DisruptionKindMock_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
func (_e *DisruptionKindMock_Expecter) Validate() *DisruptionKindMock_Validate_Call {
	return &DisruptionKindMock_Validate_Call{Call: _e.mock.On("Validate")}
}

func (_c *DisruptionKindMock_Validate_Call) Run(run func()) *DisruptionKindMock_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DisruptionKindMock_Validate_Call) Return(_a0 error) *DisruptionKindMock_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DisruptionKindMock_Validate_Call) RunAndReturn(run func() error) *DisruptionKindMock_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewDisruptionKindMock creates a new instance of DisruptionKindMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisruptionKindMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisruptionKindMock {
	mock := &DisruptionKindMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
