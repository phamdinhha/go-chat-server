package dto

import "github.com/phamdinhha/go-chat-server/internal/models"

type LoginResponse struct {
	User     *models.User `json:"User"`
	JwtToken string       `json:"Token"`
}

type SignUpResponse struct {
	User     *models.User `json:"User"`
	JwtToken string       `json:"Token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
