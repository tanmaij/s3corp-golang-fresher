package repository

import (
	"s3corp-golang-fresher/internal/models"
)

type UserRepo interface {
	Login(username string) (*models.User, error)
	// GetUserByUsername Return one user with id parameter
	GetUserByUsername(username string) (*models.User, error)

	// GetUsers Return a user Slice
	GetUsers() (models.UserSlice, error)

	// CreateUser Insert data by user parameter
	CreateUser(user models.User) error

	// UpdateUser Update one record by user parameter
	UpdateUser(user models.User) (int64, error)

	// DeleteUser Delete one record by username parameter
	DeleteUser(username string) (int64, error)

	GetByUsernameOrEmail(username string, email string) (*models.User, error)
}
