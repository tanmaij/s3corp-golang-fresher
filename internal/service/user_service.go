package service

import (
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/utils"
)

type UserService interface {

	// Login with username and password
	Login(username string, password string) (models.User, string, error)

	// GetUserByUsername Return one user with id parameter
	GetUserByUsername(username string) (models.User, error)

	// GetUsers Return a user Slice
	GetUsers(queriesParams map[string]int) ([]models.User, utils.Pagination, error)

	// CreateUser Insert data by user parameter
	CreateUser(user models.User) error

	// UpdateUser Update one record by user parameter
	UpdateUser(user models.User) error

	// DeleteUser Delete one record by username parameter
	DeleteUser(username string) error
}
