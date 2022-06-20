package reponsitory

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type SubDocumentReponsitory struct {
	Data *data.Data
}

// GetByDocumentId Return a slice subdocument by Document paramater
func (subDocumentReponsitory SubDocumentReponsitory) GetByDocumentId(documentId string) (models.SubDocumentSlice, error) {

	subDocuments, err := models.SubDocuments(qm.Where("DocumentId=?", documentId)).All(context.Background(), subDocumentReponsitory.Data.DB)

	if err != nil {
		return subDocuments, fmt.Errorf("internal Server Error")
	}

	return subDocuments, nil
}

// GetOneById Return one subdocument with id parameter
func (subDocumentReponsitory SubDocumentReponsitory) GetOneById(id int) (*models.SubDocument, error) {

	subDocument, err := models.SubDocuments(qm.Where("SubDocumentId=?", id)).One(context.Background(), subDocumentReponsitory.Data.DB)

	if err != nil {
		return subDocument, fmt.Errorf("internal Server Error")
	}

	return subDocument, nil
}

// GetAll Return a subDocument Slice
func (subDocumentReponsitory SubDocumentReponsitory) GetAll() (models.SubDocumentSlice, error) {
	subDocuments, err := models.SubDocuments().All(context.Background(), subDocumentReponsitory.Data.DB)

	if err != nil {
		return subDocuments, fmt.Errorf("internal Server Error")
	}

	return subDocuments, nil
}

// CreateOne Insert data by subdocument parameter
func (subDocumentReponsitory SubDocumentReponsitory) CreateOne(subDocument models.Document) error {

	err := subDocument.Insert(context.Background(), subDocumentReponsitory.Data.DB, boil.Infer())

	if err != nil {
		return fmt.Errorf("internal Server Error")
	}
	return nil
}

// UpdateOne Update one record by subdocument parameter
func (subDocumentReponsitory SubDocumentReponsitory) UpdateOne(subDocument models.SubDocument) (int64, error) {

	affectRows, err := subDocument.Update(context.Background(), subDocumentReponsitory.Data.DB, boil.Infer())
	if err != nil {
		return affectRows, fmt.Errorf("internal Server Error")
	}
	if affectRows <= 0 {
		return affectRows, fmt.Errorf("there isn't any record which is updated")
	}
	return affectRows, err
}

// DeleteOneById Delete one record by subdocument parameter
func (subDocumentReponsitory SubDocumentReponsitory) DeleteOneById(subDocumentId int) (int64, error) {
	affectRows, err := models.SubDocuments(qm.Where("subDocumentId", subDocumentId)).DeleteAll(context.Background(), subDocumentReponsitory.Data.DB)

	if err != nil {
		return affectRows, fmt.Errorf("internal Server Error")
	}

	if affectRows <= 0 {
		return affectRows, fmt.Errorf("there isn't any record which is deleted")
	}

	return affectRows, err
}
