package services

import (
	"context"

	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/internal/models"
	"github.com/phamdinhha/go-chat-server/internal/repositories"
)

type chatService struct {
	chatRepo repositories.ChatRepo
	roomRepo repositories.ChatRoomRepo
	cfg      config.Config
}

func NewChatService(
	chatRepo repositories.ChatRepo,
	roomRepo repositories.ChatRoomRepo,
	cfg config.Config,
) ChatService {
	return &chatService{
		chatRepo: chatRepo,
		roomRepo: roomRepo,
		cfg:      cfg,
	}
}

func (a *chatService) CreateRoom(ctx context.Context, name string) (*models.ChatRoom, error) {
	room := &models.ChatRoom{
		Name: name,
	}
	return a.roomRepo.CreateRoom(ctx, room)
}

func (a *chatService) ListByRoomId(ctx context.Context, roomID string) ([]*models.Chat, error) {
	return a.chatRepo.ListByRoomId(ctx, roomID)
}

func (a *chatService) CreateMessge(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	return a.chatRepo.CreateMessge(ctx, chat)
}
