// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.
package command

import mock "github.com/stretchr/testify/mock"

// BackgroundCmdMock is an autogenerated mock type for the BackgroundCmd type
type BackgroundCmdMock struct {
	mock.Mock
}

type BackgroundCmdMock_Expecter struct {
	mock *mock.Mock
}

func (_m *BackgroundCmdMock) EXPECT() *BackgroundCmdMock_Expecter {
	return &BackgroundCmdMock_Expecter{mock: &_m.Mock}
}

// DryRun provides a mock function with given fields:
func (_m *BackgroundCmdMock) DryRun() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DryRun")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BackgroundCmdMock_DryRun_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DryRun'
type BackgroundCmdMock_DryRun_Call struct {
	*mock.Call
}

// DryRun is a helper method to define mock.On call
func (_e *BackgroundCmdMock_Expecter) DryRun() *BackgroundCmdMock_DryRun_Call {
	return &BackgroundCmdMock_DryRun_Call{Call: _e.mock.On("DryRun")}
}

func (_c *BackgroundCmdMock_DryRun_Call) Run(run func()) *BackgroundCmdMock_DryRun_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BackgroundCmdMock_DryRun_Call) Return(_a0 bool) *BackgroundCmdMock_DryRun_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BackgroundCmdMock_DryRun_Call) RunAndReturn(run func() bool) *BackgroundCmdMock_DryRun_Call {
	_c.Call.Return(run)
	return _c
}

// KeepAlive provides a mock function with given fields:
func (_m *BackgroundCmdMock) KeepAlive() {
	_m.Called()
}

// BackgroundCmdMock_KeepAlive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeepAlive'
type BackgroundCmdMock_KeepAlive_Call struct {
	*mock.Call
}

// KeepAlive is a helper method to define mock.On call
func (_e *BackgroundCmdMock_Expecter) KeepAlive() *BackgroundCmdMock_KeepAlive_Call {
	return &BackgroundCmdMock_KeepAlive_Call{Call: _e.mock.On("KeepAlive")}
}

func (_c *BackgroundCmdMock_KeepAlive_Call) Run(run func()) *BackgroundCmdMock_KeepAlive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BackgroundCmdMock_KeepAlive_Call) Return() *BackgroundCmdMock_KeepAlive_Call {
	_c.Call.Return()
	return _c
}

func (_c *BackgroundCmdMock_KeepAlive_Call) RunAndReturn(run func()) *BackgroundCmdMock_KeepAlive_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *BackgroundCmdMock) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BackgroundCmdMock_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type BackgroundCmdMock_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *BackgroundCmdMock_Expecter) Start() *BackgroundCmdMock_Start_Call {
	return &BackgroundCmdMock_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *BackgroundCmdMock_Start_Call) Run(run func()) *BackgroundCmdMock_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BackgroundCmdMock_Start_Call) Return(_a0 error) *BackgroundCmdMock_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BackgroundCmdMock_Start_Call) RunAndReturn(run func() error) *BackgroundCmdMock_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *BackgroundCmdMock) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BackgroundCmdMock_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type BackgroundCmdMock_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *BackgroundCmdMock_Expecter) Stop() *BackgroundCmdMock_Stop_Call {
	return &BackgroundCmdMock_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *BackgroundCmdMock_Stop_Call) Run(run func()) *BackgroundCmdMock_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BackgroundCmdMock_Stop_Call) Return(_a0 error) *BackgroundCmdMock_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BackgroundCmdMock_Stop_Call) RunAndReturn(run func() error) *BackgroundCmdMock_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewBackgroundCmdMock creates a new instance of BackgroundCmdMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBackgroundCmdMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *BackgroundCmdMock {
	mock := &BackgroundCmdMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
