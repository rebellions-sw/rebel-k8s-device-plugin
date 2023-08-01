// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	types "github.com/rebellions-sw/sriov-network-device-plugin/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// VdpaDevice is an autogenerated mock type for the VdpaDevice type
type VdpaDevice struct {
	mock.Mock
}

// GetParent provides a mock function with given fields:
func (_m *VdpaDevice) GetParent() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPath provides a mock function with given fields:
func (_m *VdpaDevice) GetPath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetType provides a mock function with given fields:
func (_m *VdpaDevice) GetType() types.VdpaType {
	ret := _m.Called()

	var r0 types.VdpaType
	if rf, ok := ret.Get(0).(func() types.VdpaType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.VdpaType)
	}

	return r0
}

type NewVdpaDeviceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewVdpaDevice creates a new instance of VdpaDevice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVdpaDevice(t NewVdpaDeviceT) *VdpaDevice {
	mock := &VdpaDevice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
