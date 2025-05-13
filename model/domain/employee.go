package domain

type Employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Email    string `json:"email"`
}
