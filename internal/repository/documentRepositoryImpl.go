package repository

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	_ "s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type DocumentRepositoryImpl struct {
	Data *data.Data
}

func (documentRepositoryImpl DocumentRepositoryImpl) GetOneById(id string) (*models.Document, error) {
	return models.Documents(qm.Where("DocumentId=?", id)).One(context.Background(), documentRepositoryImpl.Data.DB)
}

func (documentRepositoryImpl DocumentRepositoryImpl) GetAll() (models.DocumentSlice, error) {
	return models.Documents().All(context.Background(), documentRepositoryImpl.Data.DB)
}

func (documentRepositoryImpl DocumentRepositoryImpl) CreateOne(document models.Document) error {
	return document.Insert(context.Background(), documentRepositoryImpl.Data.DB, boil.Infer())
}

func (documentRepositoryImpl DocumentRepositoryImpl) UpdateOne(document models.Document) (int64, error) {
	return document.Update(context.Background(), documentRepositoryImpl.Data.DB, boil.Infer())
}

func (documentRepositoryImpl DocumentRepositoryImpl) DeleteOneById(documentId string) (int64, error) {
	return models.Documents(qm.Where("DocumentId=?", documentId)).DeleteAll(context.Background(), documentRepositoryImpl.Data.DB)
}
