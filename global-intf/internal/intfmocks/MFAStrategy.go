// Code generated by mockery v2.27.1. DO NOT EDIT.

package intfmocks

import mock "github.com/stretchr/testify/mock"

// MFAStrategy is an autogenerated mock type for the MFAStrategy type
type MFAStrategy struct {
	mock.Mock
}

type mockConstructorTestingTNewMFAStrategy interface {
	mock.TestingT
	Cleanup(func())
}

// NewMFAStrategy creates a new instance of MFAStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMFAStrategy(t mockConstructorTestingTNewMFAStrategy) *MFAStrategy {
	mock := &MFAStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}