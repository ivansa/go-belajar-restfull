package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-belajar-restfull/helper"
	"go-belajar-restfull/model/domain"
)

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

// Ini adalah Implementasi dari interface Employee Repository
type EmployeeRepositoryImpl struct {
}

func (e EmployeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee {
	SQL := "select id, name, position, email from m_employee"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.CheckError(err)

	defer rows.Close()

	var employees []domain.Employee

	for rows.Next() {
		var employee domain.Employee
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Position, &employee.Email)
		helper.CheckError(err)

		employees = append(employees, employee)
	}

	return employees
}

func (e EmployeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Employee, error) {
	sQL := "select id, name, position, email from m_employee where id = ?"
	rows, err := tx.QueryContext(ctx, sQL, id)
	helper.CheckError(err)

	defer rows.Close()

	var employee domain.Employee = domain.Employee{}
	if rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Position, &employee.Email)
		helper.CheckError(err)
		return employee, nil
	} else {
		return domain.Employee{}, errors.New("data not found")
	}
}

func (e EmployeeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, employee *domain.Employee) {
	SQL := "insert into m_employee (name,position,email) values ($1,$2,$3) RETURNING id"
	rows, err := tx.QueryContext(ctx, SQL, employee.Name, employee.Position, employee.Email)
	helper.CheckError(err)

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&employee.Id)
		helper.CheckError(err)
	}

	// for MySQL
	// id, err := result.LastInsertId()
	// helper.CheckError(err)
	// employee.Id = int(id)
}

func (e EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee *domain.Employee, id int) {
	SQL := "update m_employee set name = ?, position = ?, email = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Position, employee.Email, id)
	helper.CheckError(err)
	employee.Id = id
}

func (e EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	sQL := "delete from m_employee where id = ?"
	_, err := tx.ExecContext(ctx, sQL, id)
	helper.CheckError(err)
}
