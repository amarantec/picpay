package models

type User struct {
	Id        int64
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
