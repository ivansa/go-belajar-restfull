package config

import (
	"github.com/julienschmidt/httprouter"
	"go-belajar-restfull/controller"
	"go-belajar-restfull/exception"
	"net/http"
)

func NewHttpRouter(controller controller.EmployeeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/employee", controller.FindAll)
	router.GET("/api/v1/employee/:id", controller.FindById)
	router.POST("/api/v1/employee", controller.Create)
	router.PUT("/api/v1/employee/:id", controller.Update)
	router.DELETE("/api/v1/employee/:id", controller.Delete)

	router.PanicHandler = exception.ErrorHandler
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		exception.NotFoundErrotHandler(writer)
	})

	return router
}
