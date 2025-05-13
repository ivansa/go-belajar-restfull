package service

import (
	"context"
	"go-belajar-restfull/model/dto"
)

type EmployeeService interface {
	FindAll(ctx context.Context) []dto.EmployeeResponse
	FindById(ctx context.Context, id int) dto.EmployeeResponse
	Create(ctx context.Context, request dto.EmployeeCreateRequest) dto.EmployeeResponse
	Update(ctx context.Context, request dto.EmployeeCreateRequest, id int) dto.EmployeeResponse
	Delete(ctx context.Context, id int)
}
