package repository

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type DocItemRepoImpl struct {
	Data *data.Data
}

func NewDocItemRepo(data *data.Data) DocItemRepo {
	return &DocItemRepoImpl{Data: data}
}

func (docItemRepoImpl DocItemRepoImpl) GetDocItemById(docItemId string) (*models.DocumentItem, error) {
	doc, err := models.DocumentItems(qm.Where("SubDocumentId=?", docItemId)).One(context.Background(), docItemRepoImpl.Data.DB)
	return doc, err
}

func (docItemRepoImpl DocItemRepoImpl) GetDocItems() (models.DocumentItemSlice, error) {
	return models.DocumentItems().All(context.Background(), docItemRepoImpl.Data.DB)
}

func (docItemRepoImpl DocItemRepoImpl) CreateDocItem(docItem models.DocumentItem) error {
	return docItem.Insert(context.Background(), docItemRepoImpl.Data.DB, boil.Infer())
}

func (docItemRepoImpl DocItemRepoImpl) UpdateDocItem(docItem models.DocumentItem) (int64, error) {
	return docItem.Update(context.Background(), docItemRepoImpl.Data.DB, boil.Infer())
}

func (docItemRepoImpl DocItemRepoImpl) DeleteDocById(docItemId string) (int64, error) {
	return models.DocumentItems(qm.Where("documentItemId=?", docItemId)).DeleteAll(context.Background(), docItemRepoImpl.Data.DB)
}
