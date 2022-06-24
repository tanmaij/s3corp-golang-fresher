package repository

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	_ "s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type DocRepoImpl struct {
	db *sql.DB
}

func NewDocRepo(db *sql.DB) DocRepo {
	return &DocRepoImpl{db}
}

func (docRepoImpl DocRepoImpl) GetDocById(id string) (*models.Document, error) {
	doc, err := models.Documents(qm.Where("DocumentId=?", id)).One(context.Background(), docRepoImpl.db)
	return doc, err
}
func (docRepoImpl DocRepoImpl) GetDocs() (models.DocumentSlice, error) {
	return models.Documents().All(context.Background(), docRepoImpl.db)
}

func (docRepoImpl DocRepoImpl) GetDocsByUsername(username string) (models.DocumentSlice, error) {
	return models.Documents(qm.Where("Username=?", username)).All(context.Background(), docRepoImpl.db)
}

func (docRepoImpl DocRepoImpl) CreateDoc(document models.Document) error {
	return document.Insert(context.Background(), docRepoImpl.db, boil.Infer())
}

func (docRepoImpl DocRepoImpl) UpdateDoc(document models.Document) (int64, error) {
	return document.Update(context.Background(), docRepoImpl.db, boil.Infer())
}

func (docRepoImpl DocRepoImpl) DeleteDocById(documentId string) (int64, error) {
	return models.Documents(qm.Where("DocumentId=?", documentId)).DeleteAll(context.Background(), docRepoImpl.db)
}
