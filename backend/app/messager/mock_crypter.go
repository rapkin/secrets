// Code generated by mockery v1.0.0. DO NOT EDIT.

package messager

import mock "github.com/stretchr/testify/mock"

// MockCrypter is an autogenerated mock type for the Crypter type
type MockCrypter struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: req
func (_m *MockCrypter) Decrypt(req Request) ([]byte, error) {
	ret := _m.Called(req)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(Request) []byte); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: req
func (_m *MockCrypter) Encrypt(req Request) ([]byte, error) {
	ret := _m.Called(req)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(Request) []byte); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}