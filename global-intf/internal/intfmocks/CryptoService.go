// Code generated by mockery v2.27.1. DO NOT EDIT.

package intfmocks

import mock "github.com/stretchr/testify/mock"

// CryptoService is an autogenerated mock type for the CryptoService type
type CryptoService struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hashed, salt, password
func (_m *CryptoService) ComparePassword(hashed []byte, salt []byte, password []byte) error {
	ret := _m.Called(hashed, salt, password)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte, []byte) error); ok {
		r0 = rf(hashed, salt, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenUlid provides a mock function with given fields:
func (_m *CryptoService) GenUlid() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// HashPassword provides a mock function with given fields: password
func (_m *CryptoService) HashPassword(password []byte) ([]byte, string, error) {
	ret := _m.Called(password)

	var r0 []byte
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) string); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func([]byte) error); ok {
		r2 = rf(password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewCryptoService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCryptoService creates a new instance of CryptoService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCryptoService(t mockConstructorTestingTNewCryptoService) *CryptoService {
	mock := &CryptoService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
