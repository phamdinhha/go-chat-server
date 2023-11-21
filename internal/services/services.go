package services

import (
	"context"

	"github.com/phamdinhha/go-chat-server/internal/dto"
	"github.com/phamdinhha/go-chat-server/internal/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type ChatService interface {
	CreateRoom(ctx context.Context, name string) (*models.ChatRoom, error)
	ListByRoomId(ctx context.Context, roomID string) ([]*models.Chat, error)
	CreateMessge(ctx context.Context, chat *models.Chat) (*models.Chat, error)
}

type AuthenService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	SignUp(ctx context.Context, req dto.SignUpRequest) (dto.SignUpResponse, error)
}
