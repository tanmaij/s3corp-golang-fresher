package service

import (
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/utils"
)

type DocServiceImpl struct {
	DocRepo repository.DocRepo
}

func NewDocService(docRepo repository.DocRepo) DocService {
	return &DocServiceImpl{DocRepo: docRepo}
}

func (docServiceImpl DocServiceImpl) GetDocById(docId string) (*models.Document, utils.Error) {
	document, err := docServiceImpl.DocRepo.GetDocById(docId)

	if err != nil {
		return document, utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	return document, utils.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) GetDocsByUsername(username string) (models.DocumentSlice, utils.Error) {
	documents, err := docServiceImpl.DocRepo.GetDocsByUsername(username)
	if err != nil {
		return documents, utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if len(documents) == 0 {
		return []*models.Document{}, utils.NewError("Successfully", http.StatusOK)
	}
	return documents, utils.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) GetDocs() (models.DocumentSlice, utils.Error) {
	docs, err := docServiceImpl.DocRepo.GetDocs()
	if err != nil {
		return docs, utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	if len(docs) == 0 {
		return []*models.Document{}, utils.NewError("Successfully", http.StatusOK)
	}
	return docs, utils.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) CreateDoc(document models.Document) utils.Error {
	err := docServiceImpl.DocRepo.CreateDoc(document)
	if err != nil {
		return utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}
	return utils.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) UpdateDoc(document models.Document) utils.Error {
	affectedRows, err := docServiceImpl.DocRepo.UpdateDoc(document)

	if err != nil {
		return utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	if affectedRows <= 0 {
		return utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return utils.NewError("Successfully", http.StatusOK)
}

func (docServiceImpl DocServiceImpl) DeleteDocById(documentId string) utils.Error {
	affectedRows, err := docServiceImpl.DocRepo.DeleteDocById(documentId)

	if err != nil {
		return utils.NewError("Internal Server Error", http.StatusInternalServerError)
	}

	if affectedRows <= 0 {
		return utils.Error{Message: "The record is not exists", StatusCode: 404}
	}
	return utils.NewError("Successfully", http.StatusOK)
}
