// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

// BuildNumber provides a mock function with given fields:
func (_m *Task) BuildNumber() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasStarted provides a mock function with given fields:
func (_m *Task) HasStarted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Poll provides a mock function with given fields:
func (_m *Task) Poll() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitToStart provides a mock function with given fields: pollInterval
func (_m *Task) WaitToStart(pollInterval time.Duration) (int64, error) {
	ret := _m.Called(pollInterval)

	var r0 int64
	if rf, ok := ret.Get(0).(func(time.Duration) int64); ok {
		r0 = rf(pollInterval)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Duration) error); ok {
		r1 = rf(pollInterval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
