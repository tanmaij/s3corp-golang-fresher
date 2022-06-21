package repository

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/null/v8"
	"s3corp-golang-fresher/internal/models"
	"time"
)

type DocumentRepositoryMock struct {
	DocumentId string
	Subject    string
	Createdat  time.Time
}

func (documentRepositoryMock DocumentRepositoryMock) GetAll() (models.DocumentSlice, error) {
	return []*models.Document{
			{DocumentId: string(uuid.V4), Subject: null.String{String: "Làm màu"}, Createdat: null.Time{Time: time.Date(
				2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true}},
			{DocumentId: string(uuid.V4), Subject: null.String{String: "Làm màu"}, Createdat: null.Time{Time: time.Date(
				2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true}}},
		nil
}
func (documentRepositoryMock DocumentRepositoryMock) GetOneById(id string) (*models.Document, error) {

	documents := []models.Document{
		{DocumentId: "238064cd-6b9b-49f5-848d-1c48c6ab1562\n", Subject: null.String{String: "Làm màu"}, Createdat: null.Time{Time: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true}},
		{DocumentId: "bc46ed59-c84f-480a-94f9-b01ff69c0fa1\n", Subject: null.String{String: "Làm màu 2"}, Createdat: null.Time{Time: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true}},
		{DocumentId: "ff5fb12a-f8d3-4fe0-9a45-ff473b2e00b2\n", Subject: null.String{String: "Làm màu"}, Createdat: null.Time{Time: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true}},
		{DocumentId: "6b71d523-0e3b-41e9-a693-4a0cb9b4fb19\n", Subject: null.String{String: "Làm màu"}, Createdat: null.Time{Time: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Valid: true}}}

	for _, v := range documents {
		if v.DocumentId == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("data is not exist")
}

func (documentRepositoryMock DocumentRepositoryMock) CreateOne(document models.Document) error {
	return nil
}

func (documentRepositoryMock DocumentRepositoryMock) UpdateOne(document models.Document) (int64, error) {
	return 0, nil
}

func (documentRepositoryMock DocumentRepositoryMock) DeleteOneById(documentId string) (int64, error) {
	return 0, nil
}
