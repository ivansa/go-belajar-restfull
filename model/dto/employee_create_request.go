package dto

type EmployeeCreateRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Position string `json:"position" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=150"`
}
