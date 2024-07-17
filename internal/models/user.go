package models

type User struct {
	Id        int64
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Document  string   `json:"document"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Balance   float64  `json:"balance"`
	UserId    int64    `json:"user_id"`
	UserType  UserType `json:"user_type"`
}
