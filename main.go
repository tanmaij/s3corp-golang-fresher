package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	data "s3corp-golang-fresher/data"
	"s3corp-golang-fresher/internal/handler"
	"s3corp-golang-fresher/internal/reponsitory"
	"s3corp-golang-fresher/internal/service"
)

func main() {

	data := data.Data{}

	data.Init()

	var r *chi.Mux = chi.NewRouter()
	var port string = ":3333"
	fmt.Println("Server has started on Port", port)

	documentReponsitory := reponsitory.DocumentReponsitory{Data: &data}
	documentService := service.DocumentService{DocumentReponsitory: &documentReponsitory}
	documentHandle := handler.DocumentHandler{DocumentService: &documentService}

	r.Route("/api", func(r chi.Router) {
		r.Route("/document", documentHandle.DocumentHandler)
	})

	http.ListenAndServe(port, r)

}
