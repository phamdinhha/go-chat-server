package repositories

const (
	CREATE_USER_QUERY    = `INSERT INTO users (id, user_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *`
	GET_USER_BY_EMAIL    = `SELECT id, user_name, email, password FROM users WHERE email = $1`
	LIST_CHAT_BY_ROOM_ID = `SELECT id, room_id, user_id, message, created_at FROM chats WHERE room_id = $1`
	CREATE_CHAT_QUERY    = `INSERT INTO chats (id, room_id, user_id, message, created_at) VALUES (uuid_generate_v4(), $1, $2, $3, $4) RETURNING id, room_id, user_id, message, created_at RETURNING *`
	LIST_ROOM_QUERY      = `SELECT id, name FROM chat_rooms`
	CREATE_ROOM_QUERY    = `INSERT INTO chat_rooms (id, name) VALUES (uuid_generate_v4(), $1) RETURNING *`
)
