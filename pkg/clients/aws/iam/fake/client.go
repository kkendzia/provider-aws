// Code generated by mockery v1.0.0. DO NOT EDIT.

package fake

import mock "github.com/stretchr/testify/mock"
import serviceiam "github.com/aws/aws-sdk-go-v2/service/iam"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreatePolicyAndAttach provides a mock function with given fields: username, policyName, policyDocument
func (_m *Client) CreatePolicyAndAttach(username string, policyName string, policyDocument string) (string, error) {
	ret := _m.Called(username, policyName, policyDocument)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(username, policyName, policyDocument)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(username, policyName, policyDocument)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: username
func (_m *Client) CreateUser(username string) (*serviceiam.AccessKey, error) {
	ret := _m.Called(username)

	var r0 *serviceiam.AccessKey
	if rf, ok := ret.Get(0).(func(string) *serviceiam.AccessKey); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.AccessKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicyAndDetach provides a mock function with given fields: username, policyName
func (_m *Client) DeletePolicyAndDetach(username string, policyName string) error {
	ret := _m.Called(username, policyName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(username, policyName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: username
func (_m *Client) DeleteUser(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPolicyVersion provides a mock function with given fields: policyName
func (_m *Client) GetPolicyVersion(policyName string) (string, error) {
	ret := _m.Called(policyName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(policyName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(policyName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePolicy provides a mock function with given fields: policyName, policyDocument
func (_m *Client) UpdatePolicy(policyName string, policyDocument string) (string, error) {
	ret := _m.Called(policyName, policyDocument)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(policyName, policyDocument)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(policyName, policyDocument)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
