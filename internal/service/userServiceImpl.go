package service

import (
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/utils"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &UserServiceImpl{userRepo}
}

func (userServiceImpl UserServiceImpl) Login(username string, password string) (*models.User, utils.Error) {
	user, err := userServiceImpl.UserRepo.Login(username)
	if err != nil {
		return user, utils.NewError(utils.InternalServerError, http.StatusInternalServerError)
	}
	if user.Password.String != password {
		return user, utils.NewError(utils.PasswordIsIncorrect, http.StatusUnauthorized)
	}
	return user, utils.NewError(utils.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) GetUserByUsername(username string) (*models.User, utils.Error) {
	user, err := userServiceImpl.UserRepo.GetUserByUsername(username)

	if err != nil {
		return user, utils.NewError(utils.InternalServerError, http.StatusInternalServerError)
	}
	return user, utils.NewError(utils.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) GetUsers() (models.UserSlice, utils.Error) {
	users, err := userServiceImpl.UserRepo.GetUsers()
	if err != nil {
		return users, utils.NewError(utils.InternalServerError, http.StatusInternalServerError)
	}
	return users, utils.NewError(utils.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) CreateUser(user models.User) utils.Error {

	_, err := userServiceImpl.UserRepo.GetUserByUsername(user.Username)

	if err == nil {
		return utils.NewError(utils.UserAlreadyExist, http.StatusBadRequest)
	}

	err = userServiceImpl.UserRepo.CreateUser(user)
	if err != nil {
		return utils.NewError(utils.InternalServerError, http.StatusInternalServerError)
	}
	return utils.NewError(utils.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) UpdateUser(user models.User) utils.Error {
	affectedRows, err := userServiceImpl.UserRepo.UpdateUser(user)

	if affectedRows <= 0 {
		return utils.NewError(utils.NotExist, http.StatusNotFound)
	}
	if err != nil {
		return utils.NewError(utils.Successfully, http.StatusInternalServerError)
	}
	return utils.NewError(utils.Successfully, http.StatusOK)
}

func (userServiceImpl UserServiceImpl) DeleteUser(username string) utils.Error {
	affectedRows, err := userServiceImpl.UserRepo.DeleteUser(username)
	if affectedRows <= 0 {
		return utils.NewError(utils.NotExist, http.StatusNotFound)
	}
	if err != nil {
		return utils.NewError(utils.InternalServerError, http.StatusInternalServerError)
	}
	return utils.NewError(utils.Successfully, http.StatusOK)
}
