package mocks

import (
	"github.com/stretchr/testify/mock"
	"s3corp-golang-fresher/internal/models"
)

type UserRepoMock struct {
	mock.Mock
}

func (userRepoMock *UserRepoMock) Login(username string) (*models.User, error) {
	return nil, nil
}

func (userRepoMock *UserRepoMock) GetUserByUsername(username string) (*models.User, error) {
	return nil, nil
}

func (userRepoMock *UserRepoMock) GetUsers() (models.UserSlice, error) {
	return nil, nil
}

func (userRepoMock *UserRepoMock) CreateUser(user models.User) error {
	return nil
}

func (userRepoMock *UserRepoMock) UpdateUser(user models.User) (int64, error) {
	return 0, nil
}

func (userRepoMock *UserRepoMock) DeleteUser(username string) (int64, error) {
	return 0, nil
}
func (userRepoMock *UserRepoMock) GetByUsernameOrEmail(username string, email string) (*models.User, error) {
	return nil, nil
}
func (userRepoMock *UserRepoMock) GetUsersByYear(year int) (models.UserSlice, error) {

	args := userRepoMock.Called(year)

	return args.Get(0).(models.UserSlice), args.Error(1)
}
