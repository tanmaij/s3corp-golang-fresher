package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"s3corp-golang-fresher/internal/handler"
	"s3corp-golang-fresher/internal/repository"
	"s3corp-golang-fresher/internal/service"
	"s3corp-golang-fresher/pkg"
)

func main() {
	err := godotenv.Load()
	db := pkg.NewPsqlDB()

	var r *chi.Mux = chi.NewRouter()
	var port string = ":3333"

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Server has started on Port", port)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	docItemRepo := repository.NewDocItemRepo(db)
	docItemService := service.NewDocItemService(docItemRepo)
	docItemHandler := handler.NewDocItemHandler(docItemService)

	docRepo := repository.NewDocRepo(db)
	docService := service.NewDocService(docRepo)
	docHandler := handler.NewDocHandler(docService)

	r.Route("/api", func(r chi.Router) {
		r.Route("/document", docHandler.DocHandler)
		r.Route("/document-item", docItemHandler.DocItemHandler)
		r.Route("/user", userHandler.UserHandler)
	})

	http.ListenAndServe(port, r)

}
