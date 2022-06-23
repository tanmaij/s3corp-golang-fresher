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

	userRepo := repository.NewUserRepo(&data)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	docItemRepo := repository.NewDocItemRepo(&data)
	docItemService := service.NewDocItemService(docItemRepo)
	docItemHandler := handler.NewDocItemHandler(docItemService)

	docRepo := repository.NewDocRepo(&data)
	docService := service.NewDocService(docRepo)
	docHandler := handler.NewDocHandler(docService)

	r.Route("/api", func(r chi.Router) {
		r.Route("/document", docHandler.DocHandler)
		r.Route("/document-item", docItemHandler.DocItemHandler)
		r.Route("/user", userHandler.UserHandler)
	})

	http.ListenAndServe(port, r)

}
