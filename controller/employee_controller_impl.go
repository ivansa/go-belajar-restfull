package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-belajar-restfull/helper"
	"go-belajar-restfull/model/dto"
	"go-belajar-restfull/service"
	"net/http"
	"strconv"
)

// EmployeeControllerImpl adalah implementasi dari EmployeeController
func NewEmployeeController(service service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		Service: service,
	}
}

type EmployeeControllerImpl struct {
	Service service.EmployeeService
}

func (e EmployeeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// panggil service
	responseData := e.Service.FindAll(r.Context())

	// set base response
	responseBody := helper.CreateResponse(200, "success", responseData)

	// write response
	helper.WriteResponse(w, responseBody)
}

func (e EmployeeControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// ambil id dari url
	id, err := strconv.Atoi(p.ByName("id"))
	helper.CheckError(err)

	// panggil service
	responseData := e.Service.FindById(r.Context(), id)

	// set base response
	responseBody := helper.CreateResponse(200, "success", responseData)

	// write response
	helper.WriteResponse(w, responseBody)
}

func (e EmployeeControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// mapping request body ke struct
	requestBody := dto.EmployeeCreateRequest{}
	helper.RequestBodyToStruct[dto.EmployeeCreateRequest](r, &requestBody)
	fmt.Println(requestBody)

	// panggil service
	responseData := e.Service.Create(r.Context(), requestBody)

	// set base response
	responseBody := helper.CreateResponse(200, "success", responseData)

	// write response
	helper.WriteResponse(w, responseBody)
}

func (e EmployeeControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// mapping request body ke struct
	requestBody := dto.EmployeeCreateRequest{}
	helper.RequestBodyToStruct[dto.EmployeeCreateRequest](r, &requestBody)

	// ambil id dari url
	id, err := strconv.Atoi(p.ByName("id"))
	helper.CheckError(err)

	// panggil service
	responseData := e.Service.Update(r.Context(), requestBody, id)

	// set base response
	responseBody := helper.CreateResponse(200, "success", responseData)

	// write response
	helper.WriteResponse(w, responseBody)
}

func (e EmployeeControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// ambil id dari url
	id, err := strconv.Atoi(p.ByName("id"))
	helper.CheckError(err)

	// panggil service
	e.Service.Delete(r.Context(), id)

	// set base response
	responseBody := helper.CreateResponse(200, "success", nil)

	// write response
	helper.WriteResponse(w, responseBody)
}
