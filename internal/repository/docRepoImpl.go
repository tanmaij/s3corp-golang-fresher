package repository

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	_ "s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type DocRepoImpl struct {
	Data *data.Data
}

func NewDocRepo(data *data.Data) DocRepo {
	return &DocRepoImpl{data}
}

func (docRepoImpl DocRepoImpl) GetDocById(id string) (*models.Document, error) {
	doc, err := models.Documents(qm.Where("DocumentId=?", id)).One(context.Background(), docRepoImpl.Data.DB)
	return doc, err
}
func (docRepoImpl DocRepoImpl) GetDocs() (models.DocumentSlice, error) {
	return models.Documents().All(context.Background(), docRepoImpl.Data.DB)
}

func (docRepoImpl DocRepoImpl) GetDocsByUsername(username string) (models.DocumentSlice, error) {
	return models.Documents(qm.Where("Username=?", username)).All(context.Background(), docRepoImpl.Data.DB)
}

func (docRepoImpl DocRepoImpl) CreateDoc(document models.Document) error {
	return document.Insert(context.Background(), docRepoImpl.Data.DB, boil.Infer())
}

func (docRepoImpl DocRepoImpl) UpdateDoc(document models.Document) (int64, error) {
	return document.Update(context.Background(), docRepoImpl.Data.DB, boil.Infer())
}

func (docRepoImpl DocRepoImpl) DeleteDocById(documentId string) (int64, error) {
	return models.Documents(qm.Where("DocumentId=?", documentId)).DeleteAll(context.Background(), docRepoImpl.Data.DB)
}
