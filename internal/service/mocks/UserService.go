// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	error2 "s3corp-golang-fresher/internal/errors"
	models "s3corp-golang-fresher/internal/models"

	mock "github.com/stretchr/testify/mock"

	utils "s3corp-golang-fresher/utils"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserService) CreateUser(user models.User) error2.Error {
	ret := _m.Called(user)

	var r0 error2.Error
	if rf, ok := ret.Get(0).(func(models.User) error2.Error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(error2.Error)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: username
func (_m *UserService) DeleteUser(username string) error2.Error {
	ret := _m.Called(username)

	var r0 error2.Error
	if rf, ok := ret.Get(0).(func(string) error2.Error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(error2.Error)
	}

	return r0
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *UserService) GetUserByUsername(username string) (*models.User, error2.Error) {
	ret := _m.Called(username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error2.Error
	if rf, ok := ret.Get(1).(func(string) error2.Error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Get(1).(error2.Error)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: queriesParams
func (_m *UserService) GetUsers(queriesParams map[string]string) (models.UserSlice, *utils.Pagination, error2.Error) {
	ret := _m.Called(queriesParams)

	var r0 models.UserSlice
	if rf, ok := ret.Get(0).(func(map[string]string) models.UserSlice); ok {
		r0 = rf(queriesParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.UserSlice)
		}
	}

	var r1 *utils.Pagination
	if rf, ok := ret.Get(1).(func(map[string]string) *utils.Pagination); ok {
		r1 = rf(queriesParams)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.Pagination)
		}
	}

	var r2 error2.Error
	if rf, ok := ret.Get(2).(func(map[string]string) error2.Error); ok {
		r2 = rf(queriesParams)
	} else {
		r2 = ret.Get(2).(error2.Error)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: username, password
func (_m *UserService) Login(username string, password string) (*models.User, string, error2.Error) {
	ret := _m.Called(username, password)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string, string) *models.User); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error2.Error
	if rf, ok := ret.Get(2).(func(string, string) error2.Error); ok {
		r2 = rf(username, password)
	} else {
		r2 = ret.Get(2).(error2.Error)
	}

	return r0, r1, r2
}

// UpdateUser provides a mock function with given fields: user
func (_m *UserService) UpdateUser(user models.User) error2.Error {
	ret := _m.Called(user)

	var r0 error2.Error
	if rf, ok := ret.Get(0).(func(models.User) error2.Error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(error2.Error)
	}

	return r0
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
