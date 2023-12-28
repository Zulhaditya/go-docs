package main

import (
	"belajar_golang_restful_api/app"
	"belajar_golang_restful_api/controller"
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/middleware"
	"belajar_golang_restful_api/repository"
	"belajar_golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
