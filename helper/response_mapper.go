package helper

import (
	"encoding/json"
	"go-belajar-restfull/model/domain"
	"go-belajar-restfull/model/dto"
	"net/http"
)

func EmployeeToResponse(employee domain.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		Id:       employee.Id,
		Name:     employee.Name,
		Position: employee.Position,
		Email:    employee.Email,
	}
}

func CreateResponse(statusCode int, message string, data interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
}

func WriteResponse(w http.ResponseWriter, responseBody dto.BaseResponse) {
	// set content type
	w.Header().Set("Content-Type", "application/json")

	// mapping struct ke json dan kirim ke writer
	encoder := json.NewEncoder(w)
	err := encoder.Encode(responseBody)
	CheckError(err)
}
