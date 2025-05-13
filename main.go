package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"go-belajar-restfull/controller"
	"go-belajar-restfull/database"
	"go-belajar-restfull/helper"
	"go-belajar-restfull/repository"
	"go-belajar-restfull/service"
	"net/http"
)

func main() {
	validate := validator.New()
	db := database.InitDB()

	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(db, employeeRepository, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	router := httprouter.New()

	router.GET("/api/v1/employee/list", employeeController.FindAll)
	router.GET("/api/v1/employee/detail/:id", employeeController.FindById)
	router.POST("/api/v1/employee/create", employeeController.Create)
	router.PUT("/api/v1/employee/update/:id", employeeController.Update)
	router.DELETE("/api/v1/employee/delete/:id", employeeController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.CheckError(err)
}
