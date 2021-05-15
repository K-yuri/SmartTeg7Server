// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	contracts "github.com/aws/amazon-ssm-agent/agent/session/contracts"
	datachannel "github.com/aws/amazon-ssm-agent/agent/session/datachannel"

	mock "github.com/stretchr/testify/mock"
)

// IPortSession is an autogenerated mock type for the IPortSession type
type IPortSession struct {
	mock.Mock
}

// HandleStreamMessage provides a mock function with given fields: _a0, streamDataMessage
func (_m *IPortSession) HandleStreamMessage(streamDataMessage contracts.AgentMessage) error {
	ret := _m.Called(streamDataMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(contracts.AgentMessage) error); ok {
		r0 = rf(streamDataMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsConnectionAvailable provides a mock function with given fields:
func (_m *IPortSession) IsConnectionAvailable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InitializeSession provides a mock function with given fields: _a0, serverPort
func (_m *IPortSession) InitializeSession() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *IPortSession) Stop() {
	_m.Called()
}

// WritePump provides a mock function with given fields: _a0, channel
func (_m *IPortSession) WritePump(channel datachannel.IDataChannel) int {
	ret := _m.Called(channel)

	var r0 int
	if rf, ok := ret.Get(0).(func(datachannel.IDataChannel) int); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
