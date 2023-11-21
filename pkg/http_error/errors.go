package http_error

import (
	"errors"
)

var (
	ErrInvalidCredentials  = errors.New("invalid login credentials")
	ErrInRequestMarshaling = errors.New("invalid/bad request paramenters")
	ErrDuplicateEmail      = errors.New("email already exists")
	ErrMalformedToken      = errors.New("malformed jwt token")
)

func Error(e error) {
	panic(e)
}
