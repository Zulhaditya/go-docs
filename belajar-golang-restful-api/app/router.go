package app

import (
	"belajar_golang_restful_api/controller"
	"belajar_golang_restful_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.Category) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
