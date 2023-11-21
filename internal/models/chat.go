package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type ChatRoom struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

type Chat struct {
	ID        uuid.UUID `json:"id" db:"id"`
	RoomID    string    `json:"room_id" db:"room_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
