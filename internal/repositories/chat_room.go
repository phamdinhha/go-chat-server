package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/phamdinhha/go-chat-server/internal/models"
)

type chatRoomRepo struct {
	db *sqlx.DB
}

func NewChatRoomRepo(db *sqlx.DB) ChatRoomRepo {
	return &chatRoomRepo{
		db: db,
	}
}

func (r *chatRoomRepo) CreateRoom(ctx context.Context, room *models.ChatRoom) (*models.ChatRoom, error) {
	createdRoom := &models.ChatRoom{}
	if err := r.db.QueryRowxContext(
		ctx,
		CREATE_ROOM_QUERY,
		room.ID,
		room.Name,
	).StructScan(createdRoom); err != nil {
		return nil, err
	}
	return createdRoom, nil
}

func (r *chatRoomRepo) ListRoom(ctx context.Context) ([]*models.ChatRoom, error) {
	rooms := []*models.ChatRoom{}
	if err := r.db.SelectContext(ctx, &rooms, LIST_ROOM_QUERY); err != nil {
		return nil, err
	}
	return rooms, nil
}
