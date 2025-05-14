package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-belajar-restfull/helper"
	"go-belajar-restfull/model/domain"
)

func NewEmployeeRepository() *EmployeeRepositoryImpl {
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
	SQL := "select id, name, position, email from m_employee where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, id)
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
	// for PostgreSQL using RETURNING id (untuk ambil id setelah insert)
	// untuk PostgreSQL menggunakan query context untuk ambil balikan query
	SQL := "insert into m_employee (name,position,email) values ($1,$2,$3) RETURNING id"
	rows, err := tx.QueryContext(ctx, SQL, employee.Name, employee.Position, employee.Email)
	helper.CheckError(err)

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&employee.Id)
		helper.CheckError(err)
	}

	// for MySQL untuk ambil id setelah insert
	// id, err := result.LastInsertId()
	// helper.CheckError(err)
	// employee.Id = int(id)
}

func (e EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee *domain.Employee, id int) {
	SQL := "update m_employee set name = $1, position = $2, email = $3 where id = $4"
	_, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Position, employee.Email, id)
	helper.CheckError(err)
	employee.Id = id
}

func (e EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "delete from m_employee where id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.CheckError(err)
}
