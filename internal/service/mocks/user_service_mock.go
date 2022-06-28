package mocks

import (
	"github.com/stretchr/testify/mock"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/utils"
)

type UserServiceMock struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (m *UserServiceMock) CreateUser(user models.User) error {

	args := m.Called(user)

	return args.Error(0)

}

// DeleteUser provides a mock function with given fields: username
func (m *UserServiceMock) DeleteUser(username string) error {

	args := m.Called(username)

	return args.Get(0).(error)
}

// GetUserByUsername provides a mock function with given fields: username
func (m *UserServiceMock) GetUserByUsername(username string) (models.User, error) {

	args := m.Called(username)

	return args.Get(0).(models.User), args.Get(1).(error)

}

// GetUsers provides a mock function with given fields: queriesParams
func (m *UserServiceMock) GetUsers(queriesParams map[string]int) (models.UserSlice, utils.Pagination, error) {

	args := m.Called(queriesParams)

	userSlice := args.Get(0).(models.UserSlice)

	if len(userSlice) == 0 {
		userSlice = []*models.User{}
	}

	return userSlice, args.Get(1).(utils.Pagination), args.Error(2)

}

// Login provides a mock function with given fields: username, password
func (m *UserServiceMock) Login(username string, password string) (models.User, string, error) {

	args := m.Called(username, password)

	return args.Get(0).(models.User), args.String(1), args.Error(2)

}

// UpdateUser provides a mock function with given fields: user
func (m *UserServiceMock) UpdateUser(user models.User) error {

	args := m.Called(user)

	return args.Get(0).(error)

}
