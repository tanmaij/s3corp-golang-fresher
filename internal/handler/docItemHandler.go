package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"net/http"
	"s3corp-golang-fresher/internal/models"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/utils"
)

type DocItemHandler struct {
	DocItemService service.DocItemService
}

func NewDocItemHandler(docItemService service.DocItemService) DocItemHandler {
	return DocItemHandler{DocItemService: docItemService}
}

func (docItemHandler DocItemHandler) DocItemHandler(r chi.Router) {
	r.Get("/{id}", docItemHandler.GetDocItemById)
	r.Get("/", docItemHandler.GetDocItems)
	r.Put("/", docItemHandler.UpdateDocItem)
	r.Delete("/", docItemHandler.DeleteDocItemById)
	r.Post("/", docItemHandler.CreateDocItem)
}
func (docItemHandler DocItemHandler) GetDocItemById(w http.ResponseWriter, r *http.Request) {
	docItemId := chi.URLParam(r, "id")

	docItem, err := docItemHandler.DocItemService.GetDocItemById(docItemId)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: docItem})
}
func (docItemHandler DocItemHandler) GetDocItems(w http.ResponseWriter, r *http.Request) {

	docItems, err := docItemHandler.DocItemService.GetDocItems()
	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: docItems})
}
func (docItemHandler DocItemHandler) CreateDocItem(w http.ResponseWriter, r *http.Request) {
	docItem := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&docItem)

	title, ok := docItem["title"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'title' is NOT FOUND"})
		return
	}
	content, ok := docItem["content"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'content' is NOT FOUND"})
		return
	}
	documentId, ok := docItem["documentId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'documentId' is NOT FOUND"})
		return
	}
	newDocItem := models.DocumentItem{
		Title:      null.String{String: title.(string), Valid: true},
		Content:    null.String{String: content.(string), Valid: true},
		Documentid: null.String{String: documentId.(string), Valid: true}}
	err := docItemHandler.DocItemService.CreateDocItem(newDocItem)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Create data successfully"})
}
func (docItemHandler DocItemHandler) UpdateDocItem(w http.ResponseWriter, r *http.Request) {
	docItemId := chi.URLParam(r, "id")
	docItem := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&docItem)

	title, ok := docItem["title"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'title' is NOT FOUND"})
		return
	}
	content, ok := docItem["content"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'content' is NOT FOUND"})
		return
	}
	documentId, ok := docItem["documentId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{Data: "The field 'documentId' is NOT FOUND"})
		return
	}
	newDocItem := models.DocumentItem{
		DocumentItemId: docItemId,
		Title:          null.String{String: title.(string), Valid: true},
		Content:        null.String{String: content.(string), Valid: true},
		Documentid:     null.String{String: documentId.(string), Valid: true}}
	err := docItemHandler.DocItemService.UpdateDocItem(newDocItem)

	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Update data successfully"})
}
func (docItemHandler DocItemHandler) DeleteDocItemById(w http.ResponseWriter, r *http.Request) {
	docItemId := chi.URLParam(r, "id")
	err := docItemHandler.DocItemService.DeleteDocItemById(docItemId)
	if err.StatusCode != http.StatusOK {
		err.Response(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{Success: true, Data: "Delete data successfully"})
}
