// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	mock "github.com/stretchr/testify/mock"
)

// WithdrawalCredentials is an autogenerated mock type for the WithdrawalCredentials type
type WithdrawalCredentials struct {
	mock.Mock
}

type WithdrawalCredentials_Expecter struct {
	mock *mock.Mock
}

func (_m *WithdrawalCredentials) EXPECT() *WithdrawalCredentials_Expecter {
	return &WithdrawalCredentials_Expecter{mock: &_m.Mock}
}

// ToExecutionAddress provides a mock function with given fields:
func (_m *WithdrawalCredentials) ToExecutionAddress() (common.ExecutionAddress, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ToExecutionAddress")
	}

	var r0 common.ExecutionAddress
	var r1 error
	if rf, ok := ret.Get(0).(func() (common.ExecutionAddress, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() common.ExecutionAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.ExecutionAddress)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawalCredentials_ToExecutionAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToExecutionAddress'
type WithdrawalCredentials_ToExecutionAddress_Call struct {
	*mock.Call
}

// ToExecutionAddress is a helper method to define mock.On call
func (_e *WithdrawalCredentials_Expecter) ToExecutionAddress() *WithdrawalCredentials_ToExecutionAddress_Call {
	return &WithdrawalCredentials_ToExecutionAddress_Call{Call: _e.mock.On("ToExecutionAddress")}
}

func (_c *WithdrawalCredentials_ToExecutionAddress_Call) Run(run func()) *WithdrawalCredentials_ToExecutionAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithdrawalCredentials_ToExecutionAddress_Call) Return(_a0 common.ExecutionAddress, _a1 error) *WithdrawalCredentials_ToExecutionAddress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WithdrawalCredentials_ToExecutionAddress_Call) RunAndReturn(run func() (common.ExecutionAddress, error)) *WithdrawalCredentials_ToExecutionAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewWithdrawalCredentials creates a new instance of WithdrawalCredentials. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWithdrawalCredentials(t interface {
	mock.TestingT
	Cleanup(func())
}) *WithdrawalCredentials {
	mock := &WithdrawalCredentials{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
