// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	eip4844 "github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"

	mock "github.com/stretchr/testify/mock"
)

// BlobsBundle is an autogenerated mock type for the BlobsBundle type
type BlobsBundle struct {
	mock.Mock
}

type BlobsBundle_Expecter struct {
	mock *mock.Mock
}

func (_m *BlobsBundle) EXPECT() *BlobsBundle_Expecter {
	return &BlobsBundle_Expecter{mock: &_m.Mock}
}

// GetBlobs provides a mock function with given fields:
func (_m *BlobsBundle) GetBlobs() []*eip4844.Blob {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlobs")
	}

	var r0 []*eip4844.Blob
	if rf, ok := ret.Get(0).(func() []*eip4844.Blob); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*eip4844.Blob)
		}
	}

	return r0
}

// BlobsBundle_GetBlobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlobs'
type BlobsBundle_GetBlobs_Call struct {
	*mock.Call
}

// GetBlobs is a helper method to define mock.On call
func (_e *BlobsBundle_Expecter) GetBlobs() *BlobsBundle_GetBlobs_Call {
	return &BlobsBundle_GetBlobs_Call{Call: _e.mock.On("GetBlobs")}
}

func (_c *BlobsBundle_GetBlobs_Call) Run(run func()) *BlobsBundle_GetBlobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobsBundle_GetBlobs_Call) Return(_a0 []*eip4844.Blob) *BlobsBundle_GetBlobs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobsBundle_GetBlobs_Call) RunAndReturn(run func() []*eip4844.Blob) *BlobsBundle_GetBlobs_Call {
	_c.Call.Return(run)
	return _c
}

// GetCommitments provides a mock function with given fields:
func (_m *BlobsBundle) GetCommitments() []eip4844.KZGCommitment {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCommitments")
	}

	var r0 []eip4844.KZGCommitment
	if rf, ok := ret.Get(0).(func() []eip4844.KZGCommitment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]eip4844.KZGCommitment)
		}
	}

	return r0
}

// BlobsBundle_GetCommitments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCommitments'
type BlobsBundle_GetCommitments_Call struct {
	*mock.Call
}

// GetCommitments is a helper method to define mock.On call
func (_e *BlobsBundle_Expecter) GetCommitments() *BlobsBundle_GetCommitments_Call {
	return &BlobsBundle_GetCommitments_Call{Call: _e.mock.On("GetCommitments")}
}

func (_c *BlobsBundle_GetCommitments_Call) Run(run func()) *BlobsBundle_GetCommitments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobsBundle_GetCommitments_Call) Return(_a0 []eip4844.KZGCommitment) *BlobsBundle_GetCommitments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobsBundle_GetCommitments_Call) RunAndReturn(run func() []eip4844.KZGCommitment) *BlobsBundle_GetCommitments_Call {
	_c.Call.Return(run)
	return _c
}

// GetProofs provides a mock function with given fields:
func (_m *BlobsBundle) GetProofs() []bytes.B48 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProofs")
	}

	var r0 []bytes.B48
	if rf, ok := ret.Get(0).(func() []bytes.B48); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bytes.B48)
		}
	}

	return r0
}

// BlobsBundle_GetProofs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProofs'
type BlobsBundle_GetProofs_Call struct {
	*mock.Call
}

// GetProofs is a helper method to define mock.On call
func (_e *BlobsBundle_Expecter) GetProofs() *BlobsBundle_GetProofs_Call {
	return &BlobsBundle_GetProofs_Call{Call: _e.mock.On("GetProofs")}
}

func (_c *BlobsBundle_GetProofs_Call) Run(run func()) *BlobsBundle_GetProofs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobsBundle_GetProofs_Call) Return(_a0 []bytes.B48) *BlobsBundle_GetProofs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobsBundle_GetProofs_Call) RunAndReturn(run func() []bytes.B48) *BlobsBundle_GetProofs_Call {
	_c.Call.Return(run)
	return _c
}

// NewBlobsBundle creates a new instance of BlobsBundle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobsBundle(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobsBundle {
	mock := &BlobsBundle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
