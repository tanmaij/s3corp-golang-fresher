package service

import (
	"net/http"
	error2 "s3corp-golang-fresher/internal/errors"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
)

type DocServiceImpl struct {
	DocRepo repository.DocRepo
}

func NewDocService(docRepo repository.DocRepo) DocService {
	return &DocServiceImpl{DocRepo: docRepo}
}

func (docServiceImpl DocServiceImpl) GetDocById(docId string) (*models.Document, error2.Error) {
	document, err := docServiceImpl.DocRepo.GetDocById(docId)

	if err != nil {
		return document, error2.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	return document, error2.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) GetDocsByUsername(username string) (models.DocumentSlice, error2.Error) {
	documents, err := docServiceImpl.DocRepo.GetDocsByUsername(username)
	if err != nil {
		return documents, error2.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if len(documents) == 0 {
		return []*models.Document{}, error2.NewError("Successfully", http.StatusOK)
	}
	return documents, error2.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) GetDocs() (models.DocumentSlice, error2.Error) {
	docs, err := docServiceImpl.DocRepo.GetDocs()
	if err != nil {
		return docs, error2.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if len(docs) == 0 {
		return []*models.Document{}, error2.NewError("Successfully", http.StatusOK)
	}
	return docs, error2.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) CreateDoc(document models.Document) error2.Error {
	err := docServiceImpl.DocRepo.CreateDoc(document)
	if err != nil {
		return error2.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	return error2.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) UpdateDoc(document models.Document) error2.Error {
	affectedRows, err := docServiceImpl.DocRepo.UpdateDoc(document)

	if err != nil {
		return error2.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	if affectedRows <= 0 {
		return error2.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return error2.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) DeleteDocById(documentId string) error2.Error {
	affectedRows, err := docServiceImpl.DocRepo.DeleteDocById(documentId)

	if err != nil {
		return error2.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	if affectedRows <= 0 {
		return error2.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return error2.NewError("Successfully", http.StatusOK)
}
