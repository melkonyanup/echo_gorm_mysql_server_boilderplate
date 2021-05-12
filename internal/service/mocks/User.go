// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "myapp/internal/models"

	mock "github.com/stretchr/testify/mock"

	payloads "myapp/internal/shared/payloads"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// GetUserProfile provides a mock function with given fields: email
func (_m *User) GetUserProfile(email string) (*models.User, error) {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignIn provides a mock function with given fields: payload
func (_m *User) SignIn(payload *payloads.SignInPayload) (string, error) {
	ret := _m.Called(payload)

	var r0 string
	if rf, ok := ret.Get(0).(func(*payloads.SignInPayload) string); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*payloads.SignInPayload) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: payload
func (_m *User) SignUp(payload *payloads.SignUpPayload) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*payloads.SignUpPayload) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
