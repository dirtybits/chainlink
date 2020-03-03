// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	eth "chainlink/core/eth"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	models "chainlink/core/store/models"
)

// LogRequest is an autogenerated mock type for the LogRequest type
type LogRequest struct {
	mock.Mock
}

// BlockNumber provides a mock function with given fields:
func (_m *LogRequest) BlockNumber() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// ForLogger provides a mock function with given fields: kvs
func (_m *LogRequest) ForLogger(kvs ...interface{}) []interface{} {
	var _ca []interface{}
	_ca = append(_ca, kvs...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(...interface{}) []interface{}); ok {
		r0 = rf(kvs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// GetInitiator provides a mock function with given fields:
func (_m *LogRequest) GetInitiator() models.Initiator {
	ret := _m.Called()

	var r0 models.Initiator
	if rf, ok := ret.Get(0).(func() models.Initiator); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Initiator)
	}

	return r0
}

// GetJobSpecID provides a mock function with given fields:
func (_m *LogRequest) GetJobSpecID() *models.ID {
	ret := _m.Called()

	var r0 *models.ID
	if rf, ok := ret.Get(0).(func() *models.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ID)
		}
	}

	return r0
}

// GetLog provides a mock function with given fields:
func (_m *LogRequest) GetLog() eth.Log {
	ret := _m.Called()

	var r0 eth.Log
	if rf, ok := ret.Get(0).(func() eth.Log); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(eth.Log)
	}

	return r0
}

// JSON provides a mock function with given fields:
func (_m *LogRequest) JSON() (models.JSON, error) {
	ret := _m.Called()

	var r0 models.JSON
	if rf, ok := ret.Get(0).(func() models.JSON); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.JSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunRequest provides a mock function with given fields:
func (_m *LogRequest) RunRequest() (models.RunRequest, error) {
	ret := _m.Called()

	var r0 models.RunRequest
	if rf, ok := ret.Get(0).(func() models.RunRequest); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.RunRequest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToDebug provides a mock function with given fields:
func (_m *LogRequest) ToDebug() {
	_m.Called()
}

// Validate provides a mock function with given fields:
func (_m *LogRequest) Validate() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateRequester provides a mock function with given fields:
func (_m *LogRequest) ValidateRequester() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
