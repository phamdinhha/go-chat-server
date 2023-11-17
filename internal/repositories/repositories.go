package repositories

import (
	"context"

	"github.com/phamdinhha/go-chat-server/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type ChatRepo interface {
	ListByRoomId(ctx context.Context, roomID string) ([]*models.Chat, error)
	CreateMessge(ctx context.Context, chat *models.Chat) (*models.Chat, error)
}

type ChatRoomRepo interface {
	CreateRoom(ctx context.Context, room *models.ChatRoom) (*models.ChatRoom, error)
	ListRoom(ctx context.Context) ([]*models.ChatRoom, error)
}
