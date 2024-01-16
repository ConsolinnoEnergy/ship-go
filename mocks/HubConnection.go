// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/ship-go/api"
	mock "github.com/stretchr/testify/mock"
)

// HubConnection is an autogenerated mock type for the HubConnection type
type HubConnection struct {
	mock.Mock
}

type HubConnection_Expecter struct {
	mock *mock.Mock
}

func (_m *HubConnection) EXPECT() *HubConnection_Expecter {
	return &HubConnection_Expecter{mock: &_m.Mock}
}

// AllowWaitingForTrust provides a mock function with given fields: ski
func (_m *HubConnection) AllowWaitingForTrust(ski string) bool {
	ret := _m.Called(ski)

	if len(ret) == 0 {
		panic("no return value specified for AllowWaitingForTrust")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ski)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HubConnection_AllowWaitingForTrust_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllowWaitingForTrust'
type HubConnection_AllowWaitingForTrust_Call struct {
	*mock.Call
}

// AllowWaitingForTrust is a helper method to define mock.On call
//   - ski string
func (_e *HubConnection_Expecter) AllowWaitingForTrust(ski interface{}) *HubConnection_AllowWaitingForTrust_Call {
	return &HubConnection_AllowWaitingForTrust_Call{Call: _e.mock.On("AllowWaitingForTrust", ski)}
}

func (_c *HubConnection_AllowWaitingForTrust_Call) Run(run func(ski string)) *HubConnection_AllowWaitingForTrust_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubConnection_AllowWaitingForTrust_Call) Return(_a0 bool) *HubConnection_AllowWaitingForTrust_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HubConnection_AllowWaitingForTrust_Call) RunAndReturn(run func(string) bool) *HubConnection_AllowWaitingForTrust_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteSKIConnected provides a mock function with given fields: ski
func (_m *HubConnection) RemoteSKIConnected(ski string) {
	_m.Called(ski)
}

// HubConnection_RemoteSKIConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteSKIConnected'
type HubConnection_RemoteSKIConnected_Call struct {
	*mock.Call
}

// RemoteSKIConnected is a helper method to define mock.On call
//   - ski string
func (_e *HubConnection_Expecter) RemoteSKIConnected(ski interface{}) *HubConnection_RemoteSKIConnected_Call {
	return &HubConnection_RemoteSKIConnected_Call{Call: _e.mock.On("RemoteSKIConnected", ski)}
}

func (_c *HubConnection_RemoteSKIConnected_Call) Run(run func(ski string)) *HubConnection_RemoteSKIConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubConnection_RemoteSKIConnected_Call) Return() *HubConnection_RemoteSKIConnected_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubConnection_RemoteSKIConnected_Call) RunAndReturn(run func(string)) *HubConnection_RemoteSKIConnected_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteSKIDisconnected provides a mock function with given fields: ski
func (_m *HubConnection) RemoteSKIDisconnected(ski string) {
	_m.Called(ski)
}

// HubConnection_RemoteSKIDisconnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteSKIDisconnected'
type HubConnection_RemoteSKIDisconnected_Call struct {
	*mock.Call
}

// RemoteSKIDisconnected is a helper method to define mock.On call
//   - ski string
func (_e *HubConnection_Expecter) RemoteSKIDisconnected(ski interface{}) *HubConnection_RemoteSKIDisconnected_Call {
	return &HubConnection_RemoteSKIDisconnected_Call{Call: _e.mock.On("RemoteSKIDisconnected", ski)}
}

func (_c *HubConnection_RemoteSKIDisconnected_Call) Run(run func(ski string)) *HubConnection_RemoteSKIDisconnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HubConnection_RemoteSKIDisconnected_Call) Return() *HubConnection_RemoteSKIDisconnected_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubConnection_RemoteSKIDisconnected_Call) RunAndReturn(run func(string)) *HubConnection_RemoteSKIDisconnected_Call {
	_c.Call.Return(run)
	return _c
}

// ServicePairingDetailUpdate provides a mock function with given fields: ski, detail
func (_m *HubConnection) ServicePairingDetailUpdate(ski string, detail *api.ConnectionStateDetail) {
	_m.Called(ski, detail)
}

// HubConnection_ServicePairingDetailUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServicePairingDetailUpdate'
type HubConnection_ServicePairingDetailUpdate_Call struct {
	*mock.Call
}

// ServicePairingDetailUpdate is a helper method to define mock.On call
//   - ski string
//   - detail *api.ConnectionStateDetail
func (_e *HubConnection_Expecter) ServicePairingDetailUpdate(ski interface{}, detail interface{}) *HubConnection_ServicePairingDetailUpdate_Call {
	return &HubConnection_ServicePairingDetailUpdate_Call{Call: _e.mock.On("ServicePairingDetailUpdate", ski, detail)}
}

func (_c *HubConnection_ServicePairingDetailUpdate_Call) Run(run func(ski string, detail *api.ConnectionStateDetail)) *HubConnection_ServicePairingDetailUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*api.ConnectionStateDetail))
	})
	return _c
}

func (_c *HubConnection_ServicePairingDetailUpdate_Call) Return() *HubConnection_ServicePairingDetailUpdate_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubConnection_ServicePairingDetailUpdate_Call) RunAndReturn(run func(string, *api.ConnectionStateDetail)) *HubConnection_ServicePairingDetailUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// ServiceShipIDUpdate provides a mock function with given fields: ski, shipID
func (_m *HubConnection) ServiceShipIDUpdate(ski string, shipID string) {
	_m.Called(ski, shipID)
}

// HubConnection_ServiceShipIDUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServiceShipIDUpdate'
type HubConnection_ServiceShipIDUpdate_Call struct {
	*mock.Call
}

// ServiceShipIDUpdate is a helper method to define mock.On call
//   - ski string
//   - shipID string
func (_e *HubConnection_Expecter) ServiceShipIDUpdate(ski interface{}, shipID interface{}) *HubConnection_ServiceShipIDUpdate_Call {
	return &HubConnection_ServiceShipIDUpdate_Call{Call: _e.mock.On("ServiceShipIDUpdate", ski, shipID)}
}

func (_c *HubConnection_ServiceShipIDUpdate_Call) Run(run func(ski string, shipID string)) *HubConnection_ServiceShipIDUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *HubConnection_ServiceShipIDUpdate_Call) Return() *HubConnection_ServiceShipIDUpdate_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubConnection_ServiceShipIDUpdate_Call) RunAndReturn(run func(string, string)) *HubConnection_ServiceShipIDUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// SetupRemoteDevice provides a mock function with given fields: ski, writeI
func (_m *HubConnection) SetupRemoteDevice(ski string, writeI api.SpineDataConnection) api.SpineDataProcessing {
	ret := _m.Called(ski, writeI)

	if len(ret) == 0 {
		panic("no return value specified for SetupRemoteDevice")
	}

	var r0 api.SpineDataProcessing
	if rf, ok := ret.Get(0).(func(string, api.SpineDataConnection) api.SpineDataProcessing); ok {
		r0 = rf(ski, writeI)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.SpineDataProcessing)
		}
	}

	return r0
}

// HubConnection_SetupRemoteDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupRemoteDevice'
type HubConnection_SetupRemoteDevice_Call struct {
	*mock.Call
}

// SetupRemoteDevice is a helper method to define mock.On call
//   - ski string
//   - writeI api.SpineDataConnection
func (_e *HubConnection_Expecter) SetupRemoteDevice(ski interface{}, writeI interface{}) *HubConnection_SetupRemoteDevice_Call {
	return &HubConnection_SetupRemoteDevice_Call{Call: _e.mock.On("SetupRemoteDevice", ski, writeI)}
}

func (_c *HubConnection_SetupRemoteDevice_Call) Run(run func(ski string, writeI api.SpineDataConnection)) *HubConnection_SetupRemoteDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(api.SpineDataConnection))
	})
	return _c
}

func (_c *HubConnection_SetupRemoteDevice_Call) Return(_a0 api.SpineDataProcessing) *HubConnection_SetupRemoteDevice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HubConnection_SetupRemoteDevice_Call) RunAndReturn(run func(string, api.SpineDataConnection) api.SpineDataProcessing) *HubConnection_SetupRemoteDevice_Call {
	_c.Call.Return(run)
	return _c
}

// VisibleMDNSRecordsUpdated provides a mock function with given fields: entries
func (_m *HubConnection) VisibleMDNSRecordsUpdated(entries []*api.MdnsEntry) {
	_m.Called(entries)
}

// HubConnection_VisibleMDNSRecordsUpdated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisibleMDNSRecordsUpdated'
type HubConnection_VisibleMDNSRecordsUpdated_Call struct {
	*mock.Call
}

// VisibleMDNSRecordsUpdated is a helper method to define mock.On call
//   - entries []*api.MdnsEntry
func (_e *HubConnection_Expecter) VisibleMDNSRecordsUpdated(entries interface{}) *HubConnection_VisibleMDNSRecordsUpdated_Call {
	return &HubConnection_VisibleMDNSRecordsUpdated_Call{Call: _e.mock.On("VisibleMDNSRecordsUpdated", entries)}
}

func (_c *HubConnection_VisibleMDNSRecordsUpdated_Call) Run(run func(entries []*api.MdnsEntry)) *HubConnection_VisibleMDNSRecordsUpdated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*api.MdnsEntry))
	})
	return _c
}

func (_c *HubConnection_VisibleMDNSRecordsUpdated_Call) Return() *HubConnection_VisibleMDNSRecordsUpdated_Call {
	_c.Call.Return()
	return _c
}

func (_c *HubConnection_VisibleMDNSRecordsUpdated_Call) RunAndReturn(run func([]*api.MdnsEntry)) *HubConnection_VisibleMDNSRecordsUpdated_Call {
	_c.Call.Return(run)
	return _c
}

// NewHubConnection creates a new instance of HubConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHubConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *HubConnection {
	mock := &HubConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
