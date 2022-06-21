package service

import (
	"s3corp-golang-fresher/internal/repository"
	"testing"
)

var documentRepository repository.DocumentRepository
var documentService DocumentService

func TestMain(m *testing.M) {
	documentRepository = repository.DocumentRepositoryMock{}
	documentService = DocumentServiceImpl{documentRepository}
	m.Run()
}
