package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type ChatRoom struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Chat struct {
	ID        uuid.UUID `json:"id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
