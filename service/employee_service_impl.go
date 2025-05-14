package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-belajar-restfull/exception"
	"go-belajar-restfull/helper"
	"go-belajar-restfull/model/domain"
	"go-belajar-restfull/model/dto"
	"go-belajar-restfull/repository"
)

func NewEmployeeService(db *sql.DB, employeeRepository repository.EmployeeRepository, validate *validator.Validate) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		DB:                 db,
		EmployeeRepository: employeeRepository,
		Validate:           validate,
	}
}

type EmployeeServiceImpl struct {
	DB                 *sql.DB
	EmployeeRepository repository.EmployeeRepository
	Validate           *validator.Validate
}

func (e EmployeeServiceImpl) FindAll(ctx context.Context) []dto.EmployeeResponse {
	tx, err := e.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employees := e.EmployeeRepository.FindAll(ctx, tx)
	var employeeResponses []dto.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, helper.EmployeeToResponse(employee))
	}

	// return array kosong, jgn nil
	if employeeResponses == nil {
		employeeResponses = make([]dto.EmployeeResponse, 0)
	}
	return employeeResponses
}

func (e EmployeeServiceImpl) FindById(ctx context.Context, id int) dto.EmployeeResponse {
	tx, err := e.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := e.EmployeeRepository.FindById(ctx, tx, id)
	helper.CheckError(exception.NewNotFoundError(err.Error()))

	return helper.EmployeeToResponse(employee)
}

func (e EmployeeServiceImpl) Create(ctx context.Context, request dto.EmployeeCreateRequest) dto.EmployeeResponse {
	// validasi
	err := e.Validate.Struct(request)
	helper.CheckError(err)

	tx, err := e.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employee := domain.Employee{
		Name:     request.Name,
		Position: request.Position,
		Email:    request.Email,
	}

	e.EmployeeRepository.Save(ctx, tx, &employee)
	return helper.EmployeeToResponse(employee)

}

func (e EmployeeServiceImpl) Update(ctx context.Context, request dto.EmployeeCreateRequest, id int) dto.EmployeeResponse {
	err := e.Validate.Struct(request)
	helper.CheckError(err)

	tx, err := e.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := e.EmployeeRepository.FindById(ctx, tx, id)
	helper.CheckError(exception.NewNotFoundError(err.Error()))

	employee.Name = request.Name
	employee.Position = request.Position
	employee.Email = request.Email

	e.EmployeeRepository.Update(ctx, tx, &employee, id)
	return helper.EmployeeToResponse(employee)
}

func (e EmployeeServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := e.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := e.EmployeeRepository.FindById(ctx, tx, id)
	helper.CheckError(exception.NewNotFoundError(err.Error()))

	e.EmployeeRepository.Delete(ctx, tx, employee.Id)
}
