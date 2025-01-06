// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.
package internalversion

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// DisruptionInterfaceMock is an autogenerated mock type for the DisruptionInterface type
type DisruptionInterfaceMock struct {
	mock.Mock
}

type DisruptionInterfaceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DisruptionInterfaceMock) EXPECT() *DisruptionInterfaceMock_Expecter {
	return &DisruptionInterfaceMock_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, disruption, opts
func (_m *DisruptionInterfaceMock) Create(ctx context.Context, disruption *v1beta1.Disruption, opts v1.CreateOptions) (*v1beta1.Disruption, error) {
	ret := _m.Called(ctx, disruption, opts)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *v1beta1.Disruption
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Disruption, v1.CreateOptions) (*v1beta1.Disruption, error)); ok {
		return rf(ctx, disruption, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Disruption, v1.CreateOptions) *v1beta1.Disruption); ok {
		r0 = rf(ctx, disruption, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Disruption)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.Disruption, v1.CreateOptions) error); ok {
		r1 = rf(ctx, disruption, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionInterfaceMock_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type DisruptionInterfaceMock_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - disruption *v1beta1.Disruption
//   - opts v1.CreateOptions
func (_e *DisruptionInterfaceMock_Expecter) Create(ctx interface{}, disruption interface{}, opts interface{}) *DisruptionInterfaceMock_Create_Call {
	return &DisruptionInterfaceMock_Create_Call{Call: _e.mock.On("Create", ctx, disruption, opts)}
}

func (_c *DisruptionInterfaceMock_Create_Call) Run(run func(ctx context.Context, disruption *v1beta1.Disruption, opts v1.CreateOptions)) *DisruptionInterfaceMock_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.Disruption), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *DisruptionInterfaceMock_Create_Call) Return(_a0 *v1beta1.Disruption, _a1 error) *DisruptionInterfaceMock_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionInterfaceMock_Create_Call) RunAndReturn(run func(context.Context, *v1beta1.Disruption, v1.CreateOptions) (*v1beta1.Disruption, error)) *DisruptionInterfaceMock_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *DisruptionInterfaceMock) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisruptionInterfaceMock_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DisruptionInterfaceMock_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *DisruptionInterfaceMock_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *DisruptionInterfaceMock_Delete_Call {
	return &DisruptionInterfaceMock_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *DisruptionInterfaceMock_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *DisruptionInterfaceMock_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *DisruptionInterfaceMock_Delete_Call) Return(_a0 error) *DisruptionInterfaceMock_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DisruptionInterfaceMock_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *DisruptionInterfaceMock_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *DisruptionInterfaceMock) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Disruption, error) {
	ret := _m.Called(ctx, name, opts)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *v1beta1.Disruption
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*v1beta1.Disruption, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *v1beta1.Disruption); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Disruption)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionInterfaceMock_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type DisruptionInterfaceMock_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *DisruptionInterfaceMock_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *DisruptionInterfaceMock_Get_Call {
	return &DisruptionInterfaceMock_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *DisruptionInterfaceMock_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *DisruptionInterfaceMock_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *DisruptionInterfaceMock_Get_Call) Return(_a0 *v1beta1.Disruption, _a1 error) *DisruptionInterfaceMock_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionInterfaceMock_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*v1beta1.Disruption, error)) *DisruptionInterfaceMock_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *DisruptionInterfaceMock) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DisruptionList, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *v1beta1.DisruptionList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*v1beta1.DisruptionList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *v1beta1.DisruptionList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.DisruptionList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionInterfaceMock_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type DisruptionInterfaceMock_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *DisruptionInterfaceMock_Expecter) List(ctx interface{}, opts interface{}) *DisruptionInterfaceMock_List_Call {
	return &DisruptionInterfaceMock_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *DisruptionInterfaceMock_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *DisruptionInterfaceMock_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *DisruptionInterfaceMock_List_Call) Return(_a0 *v1beta1.DisruptionList, _a1 error) *DisruptionInterfaceMock_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionInterfaceMock_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*v1beta1.DisruptionList, error)) *DisruptionInterfaceMock_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, disruption, opts
func (_m *DisruptionInterfaceMock) Update(ctx context.Context, disruption *v1beta1.Disruption, opts v1.UpdateOptions) (*v1beta1.Disruption, error) {
	ret := _m.Called(ctx, disruption, opts)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *v1beta1.Disruption
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Disruption, v1.UpdateOptions) (*v1beta1.Disruption, error)); ok {
		return rf(ctx, disruption, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Disruption, v1.UpdateOptions) *v1beta1.Disruption); ok {
		r0 = rf(ctx, disruption, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Disruption)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.Disruption, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, disruption, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionInterfaceMock_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type DisruptionInterfaceMock_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - disruption *v1beta1.Disruption
//   - opts v1.UpdateOptions
func (_e *DisruptionInterfaceMock_Expecter) Update(ctx interface{}, disruption interface{}, opts interface{}) *DisruptionInterfaceMock_Update_Call {
	return &DisruptionInterfaceMock_Update_Call{Call: _e.mock.On("Update", ctx, disruption, opts)}
}

func (_c *DisruptionInterfaceMock_Update_Call) Run(run func(ctx context.Context, disruption *v1beta1.Disruption, opts v1.UpdateOptions)) *DisruptionInterfaceMock_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.Disruption), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *DisruptionInterfaceMock_Update_Call) Return(_a0 *v1beta1.Disruption, _a1 error) *DisruptionInterfaceMock_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionInterfaceMock_Update_Call) RunAndReturn(run func(context.Context, *v1beta1.Disruption, v1.UpdateOptions) (*v1beta1.Disruption, error)) *DisruptionInterfaceMock_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *DisruptionInterfaceMock) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for Watch")
	}

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisruptionInterfaceMock_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type DisruptionInterfaceMock_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *DisruptionInterfaceMock_Expecter) Watch(ctx interface{}, opts interface{}) *DisruptionInterfaceMock_Watch_Call {
	return &DisruptionInterfaceMock_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *DisruptionInterfaceMock_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *DisruptionInterfaceMock_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *DisruptionInterfaceMock_Watch_Call) Return(_a0 watch.Interface, _a1 error) *DisruptionInterfaceMock_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DisruptionInterfaceMock_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *DisruptionInterfaceMock_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewDisruptionInterfaceMock creates a new instance of DisruptionInterfaceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisruptionInterfaceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisruptionInterfaceMock {
	mock := &DisruptionInterfaceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
