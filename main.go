package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	data "s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/handler"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/internal/service"
)

func main() {

	data := data.Data{}

	data.Init()

	var r *chi.Mux = chi.NewRouter()
	var port string = ":3333"
	fmt.Println("Server has started on Port", port)

	documentRepository := repository.DocumentRepositoryImpl{Data: &data}
	documentService := service.DocumentServiceImpl{DocumentReponsitory: documentRepository}
	documentHandler := handler.DocumentHandler{DocumentService: documentService}

	subdocumentRepository := repository.SubDocumentReponsitoryImpl{Data: &data}
	subdocumentService := service.SubDocumentServiceImpl{SubDocumentReponsitory: subdocumentRepository}
	subdocumentHandler := handler.SubDocumentHandler{SubDocumentService: subdocumentService}

	r.Route("/api", func(r chi.Router) {
		r.Route("/document", documentHandler.DocumentHandler)
		r.Route("/sub-document", subdocumentHandler.SubDocumentHandler)
	})

	http.ListenAndServe(port, r)

}
