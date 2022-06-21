package repository

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type SubDocumentReponsitoryImpl struct {
	Data *data.Data
}

func (subDocumentRepositoryImpl SubDocumentReponsitoryImpl) GetByDocumentId(documentId string) (models.SubDocumentSlice, error) {
	return models.SubDocuments(qm.Where("DocumentId=?", documentId)).All(context.Background(), subDocumentRepositoryImpl.Data.DB)
}

func (subDocumentRepositoryImpl SubDocumentReponsitoryImpl) GetOneById(id int) (*models.SubDocument, error) {
	return models.SubDocuments(qm.Where("SubDocumentId=?", id)).One(context.Background(), subDocumentRepositoryImpl.Data.DB)
}

func (subDocumentRepositoryImpl SubDocumentReponsitoryImpl) GetAll() (models.SubDocumentSlice, error) {
	return models.SubDocuments().All(context.Background(), subDocumentRepositoryImpl.Data.DB)
}

func (subDocumentRepositoryImpl SubDocumentReponsitoryImpl) CreateOne(subDocument models.SubDocument) error {
	return subDocument.Insert(context.Background(), subDocumentRepositoryImpl.Data.DB, boil.Infer())
}

func (subDocumentRepositoryImpl SubDocumentReponsitoryImpl) UpdateOne(subDocument models.SubDocument) (int64, error) {
	return subDocument.Update(context.Background(), subDocumentRepositoryImpl.Data.DB, boil.Infer())
}

func (subDocumentRepositoryImpl SubDocumentReponsitoryImpl) DeleteOneById(subDocumentId int) (int64, error) {
	return models.SubDocuments(qm.Where("subDocumentId=?", subDocumentId)).DeleteAll(context.Background(), subDocumentRepositoryImpl.Data.DB)
}
