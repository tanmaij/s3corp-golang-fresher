package repository

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type UserRepoImpl struct {
	Data *data.Data
}

func NewUserRepo(data *data.Data) UserRepo {
	return &UserRepoImpl{Data: data}
}

func (userRepoImpl UserRepoImpl) Login(username string) (*models.User, error) {
	user, err := models.Users(qm.Where("username=?", username), qm.Select("username", "password")).One(context.Background(), userRepoImpl.Data.DB)
	return user, err
}

func (userRepoImpl UserRepoImpl) GetUserByUsername(username string) (*models.User, error) {
	user, err := models.Users(qm.Where("username=?", username), qm.Select("username", "email", "name")).One(context.Background(), userRepoImpl.Data.DB)
	return user, err
}

func (userRepoImpl UserRepoImpl) GetUsers() (models.UserSlice, error) {
	return models.Users().All(context.Background(), userRepoImpl.Data.DB)
}

func (userRepoImpl UserRepoImpl) CreateUser(user models.User) error {
	return user.Insert(context.Background(), userRepoImpl.Data.DB, boil.Infer())
}

func (userRepoImpl UserRepoImpl) UpdateUser(user models.User) (int64, error) {
	return user.Update(context.Background(), userRepoImpl.Data.DB, boil.Infer())
}

func (userRepoImpl UserRepoImpl) DeleteUser(username string) (int64, error) {
	return models.Users(qm.Where("username=?", username)).DeleteAll(context.Background(), userRepoImpl.Data.DB)
}
