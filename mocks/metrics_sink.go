// Code generated by mockery. DO NOT EDIT.

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.
package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	types "github.com/DataDog/chaos-controller/types"
)

// SinkMock is an autogenerated mock type for the Sink type
type SinkMock struct {
	mock.Mock
}

type SinkMock_Expecter struct {
	mock *mock.Mock
}

func (_m *SinkMock) EXPECT() *SinkMock_Expecter {
	return &SinkMock_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *SinkMock) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type SinkMock_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *SinkMock_Expecter) Close() *SinkMock_Close_Call {
	return &SinkMock_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *SinkMock_Close_Call) Run(run func()) *SinkMock_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SinkMock_Close_Call) Return(_a0 error) *SinkMock_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_Close_Call) RunAndReturn(run func() error) *SinkMock_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetSinkName provides a mock function with given fields:
func (_m *SinkMock) GetSinkName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SinkMock_GetSinkName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSinkName'
type SinkMock_GetSinkName_Call struct {
	*mock.Call
}

// GetSinkName is a helper method to define mock.On call
func (_e *SinkMock_Expecter) GetSinkName() *SinkMock_GetSinkName_Call {
	return &SinkMock_GetSinkName_Call{Call: _e.mock.On("GetSinkName")}
}

func (_c *SinkMock_GetSinkName_Call) Run(run func()) *SinkMock_GetSinkName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SinkMock_GetSinkName_Call) Return(_a0 string) *SinkMock_GetSinkName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_GetSinkName_Call) RunAndReturn(run func() string) *SinkMock_GetSinkName_Call {
	_c.Call.Return(run)
	return _c
}

// MetricCleaned provides a mock function with given fields: succeed, kind, tags
func (_m *SinkMock) MetricCleaned(succeed bool, kind string, tags []string) error {
	ret := _m.Called(succeed, kind, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, string, []string) error); ok {
		r0 = rf(succeed, kind, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricCleaned_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricCleaned'
type SinkMock_MetricCleaned_Call struct {
	*mock.Call
}

// MetricCleaned is a helper method to define mock.On call
//   - succeed bool
//   - kind string
//   - tags []string
func (_e *SinkMock_Expecter) MetricCleaned(succeed interface{}, kind interface{}, tags interface{}) *SinkMock_MetricCleaned_Call {
	return &SinkMock_MetricCleaned_Call{Call: _e.mock.On("MetricCleaned", succeed, kind, tags)}
}

func (_c *SinkMock_MetricCleaned_Call) Run(run func(succeed bool, kind string, tags []string)) *SinkMock_MetricCleaned_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricCleaned_Call) Return(_a0 error) *SinkMock_MetricCleaned_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricCleaned_Call) RunAndReturn(run func(bool, string, []string) error) *SinkMock_MetricCleaned_Call {
	_c.Call.Return(run)
	return _c
}

// MetricCleanedForReinjection provides a mock function with given fields: succeed, kind, tags
func (_m *SinkMock) MetricCleanedForReinjection(succeed bool, kind string, tags []string) error {
	ret := _m.Called(succeed, kind, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, string, []string) error); ok {
		r0 = rf(succeed, kind, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricCleanedForReinjection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricCleanedForReinjection'
type SinkMock_MetricCleanedForReinjection_Call struct {
	*mock.Call
}

// MetricCleanedForReinjection is a helper method to define mock.On call
//   - succeed bool
//   - kind string
//   - tags []string
func (_e *SinkMock_Expecter) MetricCleanedForReinjection(succeed interface{}, kind interface{}, tags interface{}) *SinkMock_MetricCleanedForReinjection_Call {
	return &SinkMock_MetricCleanedForReinjection_Call{Call: _e.mock.On("MetricCleanedForReinjection", succeed, kind, tags)}
}

func (_c *SinkMock_MetricCleanedForReinjection_Call) Run(run func(succeed bool, kind string, tags []string)) *SinkMock_MetricCleanedForReinjection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricCleanedForReinjection_Call) Return(_a0 error) *SinkMock_MetricCleanedForReinjection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricCleanedForReinjection_Call) RunAndReturn(run func(bool, string, []string) error) *SinkMock_MetricCleanedForReinjection_Call {
	_c.Call.Return(run)
	return _c
}

// MetricCleanupDuration provides a mock function with given fields: duration, tags
func (_m *SinkMock) MetricCleanupDuration(duration time.Duration, tags []string) error {
	ret := _m.Called(duration, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, []string) error); ok {
		r0 = rf(duration, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricCleanupDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricCleanupDuration'
type SinkMock_MetricCleanupDuration_Call struct {
	*mock.Call
}

// MetricCleanupDuration is a helper method to define mock.On call
//   - duration time.Duration
//   - tags []string
func (_e *SinkMock_Expecter) MetricCleanupDuration(duration interface{}, tags interface{}) *SinkMock_MetricCleanupDuration_Call {
	return &SinkMock_MetricCleanupDuration_Call{Call: _e.mock.On("MetricCleanupDuration", duration, tags)}
}

func (_c *SinkMock_MetricCleanupDuration_Call) Run(run func(duration time.Duration, tags []string)) *SinkMock_MetricCleanupDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricCleanupDuration_Call) Return(_a0 error) *SinkMock_MetricCleanupDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricCleanupDuration_Call) RunAndReturn(run func(time.Duration, []string) error) *SinkMock_MetricCleanupDuration_Call {
	_c.Call.Return(run)
	return _c
}

// MetricDisruptionCompletedDuration provides a mock function with given fields: duration, tags
func (_m *SinkMock) MetricDisruptionCompletedDuration(duration time.Duration, tags []string) error {
	ret := _m.Called(duration, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, []string) error); ok {
		r0 = rf(duration, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricDisruptionCompletedDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricDisruptionCompletedDuration'
type SinkMock_MetricDisruptionCompletedDuration_Call struct {
	*mock.Call
}

// MetricDisruptionCompletedDuration is a helper method to define mock.On call
//   - duration time.Duration
//   - tags []string
func (_e *SinkMock_Expecter) MetricDisruptionCompletedDuration(duration interface{}, tags interface{}) *SinkMock_MetricDisruptionCompletedDuration_Call {
	return &SinkMock_MetricDisruptionCompletedDuration_Call{Call: _e.mock.On("MetricDisruptionCompletedDuration", duration, tags)}
}

func (_c *SinkMock_MetricDisruptionCompletedDuration_Call) Run(run func(duration time.Duration, tags []string)) *SinkMock_MetricDisruptionCompletedDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricDisruptionCompletedDuration_Call) Return(_a0 error) *SinkMock_MetricDisruptionCompletedDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricDisruptionCompletedDuration_Call) RunAndReturn(run func(time.Duration, []string) error) *SinkMock_MetricDisruptionCompletedDuration_Call {
	_c.Call.Return(run)
	return _c
}

// MetricDisruptionOngoingDuration provides a mock function with given fields: duration, tags
func (_m *SinkMock) MetricDisruptionOngoingDuration(duration time.Duration, tags []string) error {
	ret := _m.Called(duration, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, []string) error); ok {
		r0 = rf(duration, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricDisruptionOngoingDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricDisruptionOngoingDuration'
type SinkMock_MetricDisruptionOngoingDuration_Call struct {
	*mock.Call
}

// MetricDisruptionOngoingDuration is a helper method to define mock.On call
//   - duration time.Duration
//   - tags []string
func (_e *SinkMock_Expecter) MetricDisruptionOngoingDuration(duration interface{}, tags interface{}) *SinkMock_MetricDisruptionOngoingDuration_Call {
	return &SinkMock_MetricDisruptionOngoingDuration_Call{Call: _e.mock.On("MetricDisruptionOngoingDuration", duration, tags)}
}

func (_c *SinkMock_MetricDisruptionOngoingDuration_Call) Run(run func(duration time.Duration, tags []string)) *SinkMock_MetricDisruptionOngoingDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricDisruptionOngoingDuration_Call) Return(_a0 error) *SinkMock_MetricDisruptionOngoingDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricDisruptionOngoingDuration_Call) RunAndReturn(run func(time.Duration, []string) error) *SinkMock_MetricDisruptionOngoingDuration_Call {
	_c.Call.Return(run)
	return _c
}

// MetricDisruptionsCount provides a mock function with given fields: kind, tags
func (_m *SinkMock) MetricDisruptionsCount(kind types.DisruptionKindName, tags []string) error {
	ret := _m.Called(kind, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.DisruptionKindName, []string) error); ok {
		r0 = rf(kind, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricDisruptionsCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricDisruptionsCount'
type SinkMock_MetricDisruptionsCount_Call struct {
	*mock.Call
}

// MetricDisruptionsCount is a helper method to define mock.On call
//   - kind types.DisruptionKindName
//   - tags []string
func (_e *SinkMock_Expecter) MetricDisruptionsCount(kind interface{}, tags interface{}) *SinkMock_MetricDisruptionsCount_Call {
	return &SinkMock_MetricDisruptionsCount_Call{Call: _e.mock.On("MetricDisruptionsCount", kind, tags)}
}

func (_c *SinkMock_MetricDisruptionsCount_Call) Run(run func(kind types.DisruptionKindName, tags []string)) *SinkMock_MetricDisruptionsCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.DisruptionKindName), args[1].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricDisruptionsCount_Call) Return(_a0 error) *SinkMock_MetricDisruptionsCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricDisruptionsCount_Call) RunAndReturn(run func(types.DisruptionKindName, []string) error) *SinkMock_MetricDisruptionsCount_Call {
	_c.Call.Return(run)
	return _c
}

// MetricDisruptionsGauge provides a mock function with given fields: gauge
func (_m *SinkMock) MetricDisruptionsGauge(gauge float64) error {
	ret := _m.Called(gauge)

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(gauge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricDisruptionsGauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricDisruptionsGauge'
type SinkMock_MetricDisruptionsGauge_Call struct {
	*mock.Call
}

// MetricDisruptionsGauge is a helper method to define mock.On call
//   - gauge float64
func (_e *SinkMock_Expecter) MetricDisruptionsGauge(gauge interface{}) *SinkMock_MetricDisruptionsGauge_Call {
	return &SinkMock_MetricDisruptionsGauge_Call{Call: _e.mock.On("MetricDisruptionsGauge", gauge)}
}

func (_c *SinkMock_MetricDisruptionsGauge_Call) Run(run func(gauge float64)) *SinkMock_MetricDisruptionsGauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *SinkMock_MetricDisruptionsGauge_Call) Return(_a0 error) *SinkMock_MetricDisruptionsGauge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricDisruptionsGauge_Call) RunAndReturn(run func(float64) error) *SinkMock_MetricDisruptionsGauge_Call {
	_c.Call.Return(run)
	return _c
}

// MetricInformed provides a mock function with given fields: tags
func (_m *SinkMock) MetricInformed(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricInformed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricInformed'
type SinkMock_MetricInformed_Call struct {
	*mock.Call
}

// MetricInformed is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricInformed(tags interface{}) *SinkMock_MetricInformed_Call {
	return &SinkMock_MetricInformed_Call{Call: _e.mock.On("MetricInformed", tags)}
}

func (_c *SinkMock_MetricInformed_Call) Run(run func(tags []string)) *SinkMock_MetricInformed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricInformed_Call) Return(_a0 error) *SinkMock_MetricInformed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricInformed_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricInformed_Call {
	_c.Call.Return(run)
	return _c
}

// MetricInjectDuration provides a mock function with given fields: duration, tags
func (_m *SinkMock) MetricInjectDuration(duration time.Duration, tags []string) error {
	ret := _m.Called(duration, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, []string) error); ok {
		r0 = rf(duration, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricInjectDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricInjectDuration'
type SinkMock_MetricInjectDuration_Call struct {
	*mock.Call
}

// MetricInjectDuration is a helper method to define mock.On call
//   - duration time.Duration
//   - tags []string
func (_e *SinkMock_Expecter) MetricInjectDuration(duration interface{}, tags interface{}) *SinkMock_MetricInjectDuration_Call {
	return &SinkMock_MetricInjectDuration_Call{Call: _e.mock.On("MetricInjectDuration", duration, tags)}
}

func (_c *SinkMock_MetricInjectDuration_Call) Run(run func(duration time.Duration, tags []string)) *SinkMock_MetricInjectDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricInjectDuration_Call) Return(_a0 error) *SinkMock_MetricInjectDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricInjectDuration_Call) RunAndReturn(run func(time.Duration, []string) error) *SinkMock_MetricInjectDuration_Call {
	_c.Call.Return(run)
	return _c
}

// MetricInjected provides a mock function with given fields: succeed, kind, tags
func (_m *SinkMock) MetricInjected(succeed bool, kind string, tags []string) error {
	ret := _m.Called(succeed, kind, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, string, []string) error); ok {
		r0 = rf(succeed, kind, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricInjected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricInjected'
type SinkMock_MetricInjected_Call struct {
	*mock.Call
}

// MetricInjected is a helper method to define mock.On call
//   - succeed bool
//   - kind string
//   - tags []string
func (_e *SinkMock_Expecter) MetricInjected(succeed interface{}, kind interface{}, tags interface{}) *SinkMock_MetricInjected_Call {
	return &SinkMock_MetricInjected_Call{Call: _e.mock.On("MetricInjected", succeed, kind, tags)}
}

func (_c *SinkMock_MetricInjected_Call) Run(run func(succeed bool, kind string, tags []string)) *SinkMock_MetricInjected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricInjected_Call) Return(_a0 error) *SinkMock_MetricInjected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricInjected_Call) RunAndReturn(run func(bool, string, []string) error) *SinkMock_MetricInjected_Call {
	_c.Call.Return(run)
	return _c
}

// MetricOrphanFound provides a mock function with given fields: tags
func (_m *SinkMock) MetricOrphanFound(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricOrphanFound_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricOrphanFound'
type SinkMock_MetricOrphanFound_Call struct {
	*mock.Call
}

// MetricOrphanFound is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricOrphanFound(tags interface{}) *SinkMock_MetricOrphanFound_Call {
	return &SinkMock_MetricOrphanFound_Call{Call: _e.mock.On("MetricOrphanFound", tags)}
}

func (_c *SinkMock_MetricOrphanFound_Call) Run(run func(tags []string)) *SinkMock_MetricOrphanFound_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricOrphanFound_Call) Return(_a0 error) *SinkMock_MetricOrphanFound_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricOrphanFound_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricOrphanFound_Call {
	_c.Call.Return(run)
	return _c
}

// MetricPodsCreated provides a mock function with given fields: target, instanceName, namespace, succeed
func (_m *SinkMock) MetricPodsCreated(target string, instanceName string, namespace string, succeed bool) error {
	ret := _m.Called(target, instanceName, namespace, succeed)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, bool) error); ok {
		r0 = rf(target, instanceName, namespace, succeed)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricPodsCreated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricPodsCreated'
type SinkMock_MetricPodsCreated_Call struct {
	*mock.Call
}

// MetricPodsCreated is a helper method to define mock.On call
//   - target string
//   - instanceName string
//   - namespace string
//   - succeed bool
func (_e *SinkMock_Expecter) MetricPodsCreated(target interface{}, instanceName interface{}, namespace interface{}, succeed interface{}) *SinkMock_MetricPodsCreated_Call {
	return &SinkMock_MetricPodsCreated_Call{Call: _e.mock.On("MetricPodsCreated", target, instanceName, namespace, succeed)}
}

func (_c *SinkMock_MetricPodsCreated_Call) Run(run func(target string, instanceName string, namespace string, succeed bool)) *SinkMock_MetricPodsCreated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(bool))
	})
	return _c
}

func (_c *SinkMock_MetricPodsCreated_Call) Return(_a0 error) *SinkMock_MetricPodsCreated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricPodsCreated_Call) RunAndReturn(run func(string, string, string, bool) error) *SinkMock_MetricPodsCreated_Call {
	_c.Call.Return(run)
	return _c
}

// MetricPodsGauge provides a mock function with given fields: gauge
func (_m *SinkMock) MetricPodsGauge(gauge float64) error {
	ret := _m.Called(gauge)

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(gauge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricPodsGauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricPodsGauge'
type SinkMock_MetricPodsGauge_Call struct {
	*mock.Call
}

// MetricPodsGauge is a helper method to define mock.On call
//   - gauge float64
func (_e *SinkMock_Expecter) MetricPodsGauge(gauge interface{}) *SinkMock_MetricPodsGauge_Call {
	return &SinkMock_MetricPodsGauge_Call{Call: _e.mock.On("MetricPodsGauge", gauge)}
}

func (_c *SinkMock_MetricPodsGauge_Call) Run(run func(gauge float64)) *SinkMock_MetricPodsGauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *SinkMock_MetricPodsGauge_Call) Return(_a0 error) *SinkMock_MetricPodsGauge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricPodsGauge_Call) RunAndReturn(run func(float64) error) *SinkMock_MetricPodsGauge_Call {
	_c.Call.Return(run)
	return _c
}

// MetricReconcile provides a mock function with given fields:
func (_m *SinkMock) MetricReconcile() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricReconcile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricReconcile'
type SinkMock_MetricReconcile_Call struct {
	*mock.Call
}

// MetricReconcile is a helper method to define mock.On call
func (_e *SinkMock_Expecter) MetricReconcile() *SinkMock_MetricReconcile_Call {
	return &SinkMock_MetricReconcile_Call{Call: _e.mock.On("MetricReconcile")}
}

func (_c *SinkMock_MetricReconcile_Call) Run(run func()) *SinkMock_MetricReconcile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SinkMock_MetricReconcile_Call) Return(_a0 error) *SinkMock_MetricReconcile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricReconcile_Call) RunAndReturn(run func() error) *SinkMock_MetricReconcile_Call {
	_c.Call.Return(run)
	return _c
}

// MetricReconcileDuration provides a mock function with given fields: duration, tags
func (_m *SinkMock) MetricReconcileDuration(duration time.Duration, tags []string) error {
	ret := _m.Called(duration, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, []string) error); ok {
		r0 = rf(duration, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricReconcileDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricReconcileDuration'
type SinkMock_MetricReconcileDuration_Call struct {
	*mock.Call
}

// MetricReconcileDuration is a helper method to define mock.On call
//   - duration time.Duration
//   - tags []string
func (_e *SinkMock_Expecter) MetricReconcileDuration(duration interface{}, tags interface{}) *SinkMock_MetricReconcileDuration_Call {
	return &SinkMock_MetricReconcileDuration_Call{Call: _e.mock.On("MetricReconcileDuration", duration, tags)}
}

func (_c *SinkMock_MetricReconcileDuration_Call) Run(run func(duration time.Duration, tags []string)) *SinkMock_MetricReconcileDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricReconcileDuration_Call) Return(_a0 error) *SinkMock_MetricReconcileDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricReconcileDuration_Call) RunAndReturn(run func(time.Duration, []string) error) *SinkMock_MetricReconcileDuration_Call {
	_c.Call.Return(run)
	return _c
}

// MetricReinjected provides a mock function with given fields: succeed, kind, tags
func (_m *SinkMock) MetricReinjected(succeed bool, kind string, tags []string) error {
	ret := _m.Called(succeed, kind, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, string, []string) error); ok {
		r0 = rf(succeed, kind, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricReinjected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricReinjected'
type SinkMock_MetricReinjected_Call struct {
	*mock.Call
}

// MetricReinjected is a helper method to define mock.On call
//   - succeed bool
//   - kind string
//   - tags []string
func (_e *SinkMock_Expecter) MetricReinjected(succeed interface{}, kind interface{}, tags interface{}) *SinkMock_MetricReinjected_Call {
	return &SinkMock_MetricReinjected_Call{Call: _e.mock.On("MetricReinjected", succeed, kind, tags)}
}

func (_c *SinkMock_MetricReinjected_Call) Run(run func(succeed bool, kind string, tags []string)) *SinkMock_MetricReinjected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricReinjected_Call) Return(_a0 error) *SinkMock_MetricReinjected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricReinjected_Call) RunAndReturn(run func(bool, string, []string) error) *SinkMock_MetricReinjected_Call {
	_c.Call.Return(run)
	return _c
}

// MetricRestart provides a mock function with given fields:
func (_m *SinkMock) MetricRestart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricRestart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricRestart'
type SinkMock_MetricRestart_Call struct {
	*mock.Call
}

// MetricRestart is a helper method to define mock.On call
func (_e *SinkMock_Expecter) MetricRestart() *SinkMock_MetricRestart_Call {
	return &SinkMock_MetricRestart_Call{Call: _e.mock.On("MetricRestart")}
}

func (_c *SinkMock_MetricRestart_Call) Run(run func()) *SinkMock_MetricRestart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SinkMock_MetricRestart_Call) Return(_a0 error) *SinkMock_MetricRestart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricRestart_Call) RunAndReturn(run func() error) *SinkMock_MetricRestart_Call {
	_c.Call.Return(run)
	return _c
}

// MetricSelectorCacheGauge provides a mock function with given fields: gauge
func (_m *SinkMock) MetricSelectorCacheGauge(gauge float64) error {
	ret := _m.Called(gauge)

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(gauge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricSelectorCacheGauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricSelectorCacheGauge'
type SinkMock_MetricSelectorCacheGauge_Call struct {
	*mock.Call
}

// MetricSelectorCacheGauge is a helper method to define mock.On call
//   - gauge float64
func (_e *SinkMock_Expecter) MetricSelectorCacheGauge(gauge interface{}) *SinkMock_MetricSelectorCacheGauge_Call {
	return &SinkMock_MetricSelectorCacheGauge_Call{Call: _e.mock.On("MetricSelectorCacheGauge", gauge)}
}

func (_c *SinkMock_MetricSelectorCacheGauge_Call) Run(run func(gauge float64)) *SinkMock_MetricSelectorCacheGauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *SinkMock_MetricSelectorCacheGauge_Call) Return(_a0 error) *SinkMock_MetricSelectorCacheGauge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricSelectorCacheGauge_Call) RunAndReturn(run func(float64) error) *SinkMock_MetricSelectorCacheGauge_Call {
	_c.Call.Return(run)
	return _c
}

// MetricSelectorCacheTriggered provides a mock function with given fields: tags
func (_m *SinkMock) MetricSelectorCacheTriggered(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricSelectorCacheTriggered_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricSelectorCacheTriggered'
type SinkMock_MetricSelectorCacheTriggered_Call struct {
	*mock.Call
}

// MetricSelectorCacheTriggered is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricSelectorCacheTriggered(tags interface{}) *SinkMock_MetricSelectorCacheTriggered_Call {
	return &SinkMock_MetricSelectorCacheTriggered_Call{Call: _e.mock.On("MetricSelectorCacheTriggered", tags)}
}

func (_c *SinkMock_MetricSelectorCacheTriggered_Call) Run(run func(tags []string)) *SinkMock_MetricSelectorCacheTriggered_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricSelectorCacheTriggered_Call) Return(_a0 error) *SinkMock_MetricSelectorCacheTriggered_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricSelectorCacheTriggered_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricSelectorCacheTriggered_Call {
	_c.Call.Return(run)
	return _c
}

// MetricStuckOnRemoval provides a mock function with given fields: tags
func (_m *SinkMock) MetricStuckOnRemoval(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricStuckOnRemoval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricStuckOnRemoval'
type SinkMock_MetricStuckOnRemoval_Call struct {
	*mock.Call
}

// MetricStuckOnRemoval is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricStuckOnRemoval(tags interface{}) *SinkMock_MetricStuckOnRemoval_Call {
	return &SinkMock_MetricStuckOnRemoval_Call{Call: _e.mock.On("MetricStuckOnRemoval", tags)}
}

func (_c *SinkMock_MetricStuckOnRemoval_Call) Run(run func(tags []string)) *SinkMock_MetricStuckOnRemoval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricStuckOnRemoval_Call) Return(_a0 error) *SinkMock_MetricStuckOnRemoval_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricStuckOnRemoval_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricStuckOnRemoval_Call {
	_c.Call.Return(run)
	return _c
}

// MetricStuckOnRemovalGauge provides a mock function with given fields: gauge
func (_m *SinkMock) MetricStuckOnRemovalGauge(gauge float64) error {
	ret := _m.Called(gauge)

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(gauge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricStuckOnRemovalGauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricStuckOnRemovalGauge'
type SinkMock_MetricStuckOnRemovalGauge_Call struct {
	*mock.Call
}

// MetricStuckOnRemovalGauge is a helper method to define mock.On call
//   - gauge float64
func (_e *SinkMock_Expecter) MetricStuckOnRemovalGauge(gauge interface{}) *SinkMock_MetricStuckOnRemovalGauge_Call {
	return &SinkMock_MetricStuckOnRemovalGauge_Call{Call: _e.mock.On("MetricStuckOnRemovalGauge", gauge)}
}

func (_c *SinkMock_MetricStuckOnRemovalGauge_Call) Run(run func(gauge float64)) *SinkMock_MetricStuckOnRemovalGauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *SinkMock_MetricStuckOnRemovalGauge_Call) Return(_a0 error) *SinkMock_MetricStuckOnRemovalGauge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricStuckOnRemovalGauge_Call) RunAndReturn(run func(float64) error) *SinkMock_MetricStuckOnRemovalGauge_Call {
	_c.Call.Return(run)
	return _c
}

// MetricValidationCreated provides a mock function with given fields: tags
func (_m *SinkMock) MetricValidationCreated(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricValidationCreated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricValidationCreated'
type SinkMock_MetricValidationCreated_Call struct {
	*mock.Call
}

// MetricValidationCreated is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricValidationCreated(tags interface{}) *SinkMock_MetricValidationCreated_Call {
	return &SinkMock_MetricValidationCreated_Call{Call: _e.mock.On("MetricValidationCreated", tags)}
}

func (_c *SinkMock_MetricValidationCreated_Call) Run(run func(tags []string)) *SinkMock_MetricValidationCreated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricValidationCreated_Call) Return(_a0 error) *SinkMock_MetricValidationCreated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricValidationCreated_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricValidationCreated_Call {
	_c.Call.Return(run)
	return _c
}

// MetricValidationDeleted provides a mock function with given fields: tags
func (_m *SinkMock) MetricValidationDeleted(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricValidationDeleted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricValidationDeleted'
type SinkMock_MetricValidationDeleted_Call struct {
	*mock.Call
}

// MetricValidationDeleted is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricValidationDeleted(tags interface{}) *SinkMock_MetricValidationDeleted_Call {
	return &SinkMock_MetricValidationDeleted_Call{Call: _e.mock.On("MetricValidationDeleted", tags)}
}

func (_c *SinkMock_MetricValidationDeleted_Call) Run(run func(tags []string)) *SinkMock_MetricValidationDeleted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricValidationDeleted_Call) Return(_a0 error) *SinkMock_MetricValidationDeleted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricValidationDeleted_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricValidationDeleted_Call {
	_c.Call.Return(run)
	return _c
}

// MetricValidationFailed provides a mock function with given fields: tags
func (_m *SinkMock) MetricValidationFailed(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricValidationFailed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricValidationFailed'
type SinkMock_MetricValidationFailed_Call struct {
	*mock.Call
}

// MetricValidationFailed is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricValidationFailed(tags interface{}) *SinkMock_MetricValidationFailed_Call {
	return &SinkMock_MetricValidationFailed_Call{Call: _e.mock.On("MetricValidationFailed", tags)}
}

func (_c *SinkMock_MetricValidationFailed_Call) Run(run func(tags []string)) *SinkMock_MetricValidationFailed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricValidationFailed_Call) Return(_a0 error) *SinkMock_MetricValidationFailed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricValidationFailed_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricValidationFailed_Call {
	_c.Call.Return(run)
	return _c
}

// MetricValidationUpdated provides a mock function with given fields: tags
func (_m *SinkMock) MetricValidationUpdated(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SinkMock_MetricValidationUpdated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetricValidationUpdated'
type SinkMock_MetricValidationUpdated_Call struct {
	*mock.Call
}

// MetricValidationUpdated is a helper method to define mock.On call
//   - tags []string
func (_e *SinkMock_Expecter) MetricValidationUpdated(tags interface{}) *SinkMock_MetricValidationUpdated_Call {
	return &SinkMock_MetricValidationUpdated_Call{Call: _e.mock.On("MetricValidationUpdated", tags)}
}

func (_c *SinkMock_MetricValidationUpdated_Call) Run(run func(tags []string)) *SinkMock_MetricValidationUpdated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *SinkMock_MetricValidationUpdated_Call) Return(_a0 error) *SinkMock_MetricValidationUpdated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SinkMock_MetricValidationUpdated_Call) RunAndReturn(run func([]string) error) *SinkMock_MetricValidationUpdated_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSinkMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewSinkMock creates a new instance of SinkMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSinkMock(t mockConstructorTestingTNewSinkMock) *SinkMock {
	mock := &SinkMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}