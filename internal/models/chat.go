package models

type ChatRoom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Chat struct {
	RoomID  string `json:"room_id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
