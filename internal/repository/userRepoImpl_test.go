package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/volatiletech/null/v8"
	"s3corp-golang-fresher/internal/models"
	"testing"
)

var userRepo UserRepo

func TestUserRepo_UpdateUser(t *testing.T) {

	db, mock, _ := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "main"."user"`).WithArgs("1", "mai@gmail.com", "Mãi", "mai").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	db.Begin()
	userRepo = NewUserRepo(db)
	affectedRows, err := userRepo.UpdateUser(models.User{
		Username: "mai",
		Password: null.String{String: "1", Valid: true},
		Email:    null.String{String: "mai@gmail.com", Valid: true},
		Name:     null.String{String: "Mãi", Valid: true}})

	t.Log(err, affectedRows)
	if err != nil {
		t.Errorf("Update is failed ")
	}

	db.Close()
}
func TestUserRepo_CreateUser(t *testing.T) {

	db, mock, _ := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "main"."user"`).WithArgs("mai", "1", "mai@gmail.com", "Mãi").WillReturnResult(sqlmock.NewErrorResult(nil))
	mock.ExpectCommit()

	db.Begin()

	userRepo = NewUserRepo(db)
	err := userRepo.CreateUser(models.User{
		Password: null.String{String: "1", Valid: true},
		Username: "mai",
		Email:    null.String{String: "mai@gmail.com", Valid: true},
		Name:     null.String{String: "Mãi", Valid: true}})

	t.Log(err)
	if err != nil {
		t.Errorf("Insert is failed ")
	}
	db.Close()
}
func TestUser_DeleteUser(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "main"."user"`).WithArgs("mai").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	db.Begin()

	userRepo = NewUserRepo(db)
	affectedRows, err := userRepo.DeleteUser("mai")

	t.Log(err)
	if err != nil {
		t.Errorf("Delete is failed ")
	}
	if affectedRows != 1 {
		t.Errorf("Delete is failed ")
	}
	db.Close()
}
func TestUserRepo_GetUsers(t *testing.T) {

	db, mock, _ := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT "main"."user".* FROM "main"."user"`).WillReturnRows(sqlmock.NewRows([]string{"username", "password", "email", "name"}).AddRow("mai", "1", "email", "Mai"))
	mock.ExpectCommit()

	db.Begin()

	userRepo = NewUserRepo(db)
	users, err := userRepo.GetUsers()
	if err != nil {
		t.Errorf("Error get data")
	}
	t.Log(err, users[0].Username)
	db.Close()
}
func TestUserRepo_Login(t *testing.T) {

	db, mock, _ := sqlmock.New()

	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT "username", "password" FROM "main"."user"`).WithArgs("mai").WillReturnRows(sqlmock.NewRows([]string{"username", "password"}).AddRow("mai", "1"))
	mock.ExpectCommit()

	db.Begin()

	userRepo = NewUserRepo(db)
	user, err := userRepo.Login("mai")
	if err != nil {
		t.Errorf("Error get data")
	}
	t.Log(err, user)
	db.Close()
}
