// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// CacheWatcherMock is an autogenerated mock type for the Watcher type
type CacheWatcherMock struct {
	mock.Mock
}

type CacheWatcherMock_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheWatcherMock) EXPECT() *CacheWatcherMock_Expecter {
	return &CacheWatcherMock_Expecter{mock: &_m.Mock}
}

// Watch provides a mock function with given fields: options
func (_m *CacheWatcherMock) Watch(options v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(options)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(v1.ListOptions) (watch.Interface, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(v1.ListOptions) watch.Interface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(v1.ListOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheWatcherMock_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type CacheWatcherMock_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - options v1.ListOptions
func (_e *CacheWatcherMock_Expecter) Watch(options interface{}) *CacheWatcherMock_Watch_Call {
	return &CacheWatcherMock_Watch_Call{Call: _e.mock.On("Watch", options)}
}

func (_c *CacheWatcherMock_Watch_Call) Run(run func(options v1.ListOptions)) *CacheWatcherMock_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(v1.ListOptions))
	})
	return _c
}

func (_c *CacheWatcherMock_Watch_Call) Return(_a0 watch.Interface, _a1 error) *CacheWatcherMock_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheWatcherMock_Watch_Call) RunAndReturn(run func(v1.ListOptions) (watch.Interface, error)) *CacheWatcherMock_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCacheWatcherMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheWatcherMock creates a new instance of CacheWatcherMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheWatcherMock(t mockConstructorTestingTNewCacheWatcherMock) *CacheWatcherMock {
	mock := &CacheWatcherMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}