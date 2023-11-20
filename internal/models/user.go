package models

type User struct {
	ID       string `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
