package models

import "github.com/golang-jwt/jwt"

type Token struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	*jwt.StandardClaims
}
