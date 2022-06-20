package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"s3corp-golang-fresher/internal/service"
)

type SubDocument struct {
	subDocumentService *service.SubDocumentService
}

func (subDocument SubDocument) SubDocumentHandler(r chi.Router) {
	r.Get("/{id}", subDocument.getOneById)
	r.Get("/", subDocument.getAll)
	r.Put("/", subDocument.updateOneById)
	r.Delete("/", subDocument.deleteOneById)
	r.Post("/", subDocument.createOne)
}
func (subDocument SubDocument) getOneById(w http.ResponseWriter, r *http.Request) {

}
func (subDocument SubDocument) getAll(w http.ResponseWriter, r *http.Request) {

}
func (subDocument SubDocument) createOne(w http.ResponseWriter, r *http.Request) {

}
func (subDocument SubDocument) updateOneById(w http.ResponseWriter, r *http.Request) {

}
func (subDocument SubDocument) deleteOneById(w http.ResponseWriter, r *http.Request) {

}
