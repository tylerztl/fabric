// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	validation "github.com/hyperledger/fabric/core/handlers/validation/api/identities"
)

// IdentityDeserializer is an autogenerated mock type for the IdentityDeserializer type
type IdentityDeserializer struct {
	mock.Mock
}

// DeserializeIdentity provides a mock function with given fields: serializedIdentity
func (_m *IdentityDeserializer) DeserializeIdentity(serializedIdentity []byte) (validation.Identity, error) {
	ret := _m.Called(serializedIdentity)

	var r0 validation.Identity
	if rf, ok := ret.Get(0).(func([]byte) validation.Identity); ok {
		r0 = rf(serializedIdentity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(serializedIdentity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
