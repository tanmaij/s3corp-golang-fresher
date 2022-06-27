package service

import (
	error2 "s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/utils"
)

type UserService interface {

	// Login with username and password
	Login(username string, password string) (*models.User, string, error2.Error)

	// GetUserByUsername Return one user with id parameter
	GetUserByUsername(username string) (*models.User, error2.Error)

	// GetUsers Return a user Slice
	GetUsers(queriesParams map[string]string) (models.UserSlice, *utils.Pagination, error2.Error)

	// CreateUser Insert data by user parameter
	CreateUser(user models.User) error2.Error

	// UpdateUser Update one record by user parameter
	UpdateUser(user models.User) error2.Error

	// DeleteUser Delete one record by username parameter
	DeleteUser(username string) error2.Error
}
