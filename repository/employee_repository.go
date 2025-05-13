package repository

import (
	"context"
	"database/sql"
	"go-belajar-restfull/model/domain"
)

// Ini adalah kontrak untuk repository Employee
type EmployeeRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Employee, error)
	Save(ctx context.Context, tx *sql.Tx, employee *domain.Employee)
	Update(ctx context.Context, tx *sql.Tx, employee *domain.Employee, id int)
	Delete(ctx context.Context, tx *sql.Tx, id int)
}
