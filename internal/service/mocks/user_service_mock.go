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
func (userServiceMock *UserServiceMock) CreateUser(user models.User) error {

	args := userServiceMock.Called(user)

	return args.Error(0)

}

// DeleteUser provides a mock function with given fields: username
func (userServiceMock *UserServiceMock) DeleteUser(username string) error {

	args := userServiceMock.Called(username)

	return args.Error(0)
}

// GetUserByUsername provides a mock function with given fields: username
func (userServiceMock *UserServiceMock) GetUserByUsername(username string) (models.User, error) {

	args := userServiceMock.Called(username)

	return args.Get(0).(models.User), args.Get(1).(error)

}

// GetUsers provides a mock function with given fields: queriesParams
func (userServiceMock *UserServiceMock) GetUsers(queriesParams map[string]int) ([]models.User, utils.Pagination, error) {

	args := userServiceMock.Called(queriesParams)

	users := args.Get(0).([]models.User)

	if len(users) == 0 {
		users = []models.User{}
	}

	return users, args.Get(1).(utils.Pagination), args.Error(2)

}

// Login provides a mock function with given fields: username, password
func (userServiceMock *UserServiceMock) Login(username string, password string) (models.User, string, error) {

	args := userServiceMock.Called(username, password)

	return args.Get(0).(models.User), args.String(1), args.Error(2)

}

// UpdateUser provides a mock function with given fields: user
func (userServiceMock *UserServiceMock) UpdateUser(user models.User) error {

	args := userServiceMock.Called(user)

	return args.Error(0)

}
func (userServiceMock *UserServiceMock) UsersStatsCSVFile(year int) ([]byte, error) {
	args := userServiceMock.Called()

	return args.Get(0).([]byte), args.Error(1)
}
