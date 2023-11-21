package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/phamdinhha/go-chat-server/internal/models"
)

type chatRepo struct {
	db *sqlx.DB
}

func NewChatRepo(db *sqlx.DB) ChatRepo {
	return &chatRepo{
		db: db,
	}
}

func (r *chatRepo) ListByRoomId(ctx context.Context, roomID string) ([]*models.Chat, error) {
	chats := []*models.Chat{}
	if err := r.db.SelectContext(ctx, &chats, LIST_CHAT_BY_ROOM_ID, roomID); err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepo) CreateMessge(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	createdChat := &models.Chat{}
	if err := r.db.QueryRowxContext(
		ctx,
		CREATE_CHAT_QUERY,
		chat.RoomID,
		chat.UserID,
		chat.Message,
		chat.CreatedAt,
	).StructScan(createdChat); err != nil {
		return nil, err
	}
	return createdChat, nil
}
