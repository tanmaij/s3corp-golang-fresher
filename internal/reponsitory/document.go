package reponsitory

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"s3corp-golang-fresher/data"
	_ "s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/models"
)

type DocumentReponsitory struct {
	Data *data.Data
}

// GetOneById Return one document with id parameter
func (documentReponsitory DocumentReponsitory) GetOneById(id string) (*models.Document, error) {

	document, err := models.Documents(qm.Where("DocumentId=?", id)).One(context.Background(), documentReponsitory.Data.DB)

	if err != nil {
		return document, fmt.Errorf("internal Server Error")
	}

	return document, nil
}

// GetAll Return a Document Slice
func (documentReponsitory DocumentReponsitory) GetAll() (models.DocumentSlice, error) {
	documents, err := models.Documents().All(context.Background(), documentReponsitory.Data.DB)

	if err != nil {
		return documents, fmt.Errorf("internal Server Error")
	}

	return documents, nil
}

// CreateOne Insert data by document parameter
func (documentReponsitory DocumentReponsitory) CreateOne(document models.Document) error {

	err := document.Insert(context.Background(), documentReponsitory.Data.DB, boil.Infer())

	if err != nil {
		return fmt.Errorf("internal Server Error")
	}
	return nil
}

// UpdateOne Update one record by document parameter
func (documentReponsitory DocumentReponsitory) UpdateOne(document models.Document) (int64, error) {

	affectRows, err := document.Update(context.Background(), documentReponsitory.Data.DB, boil.Infer())
	if err != nil {
		return affectRows, fmt.Errorf("internal Server Error")
	}
	if affectRows <= 0 {
		return affectRows, fmt.Errorf("there isn't any record which is updated")
	}
	return affectRows, err
}

// DeleteOneById Delete one record by id parameter
func (documentReponsitory DocumentReponsitory) DeleteOneById(documentId int) (int64, error) {
	affectRows, err := models.Documents(qm.Where("DocumentId", documentId)).DeleteAll(context.Background(), documentReponsitory.Data.DB)

	if err != nil {
		return affectRows, fmt.Errorf("internal Server Error")
	}

	if affectRows <= 0 {
		return affectRows, fmt.Errorf("there isn't any record which is deleted")
	}

	return affectRows, err
}
