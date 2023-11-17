package models

import "github.com/golang-jwt/jwt"

type Token struct {
	UserID   string
	UserName string
	Email    string
	*jwt.StandardClaims
}
