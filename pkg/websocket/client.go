package websocket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Connection *websocket.Conn
	Pool       *Pool
	Email      string
	UserID     string
}

func (c *Client) Send(message Message) error {
	return c.Connection.WriteJSON(message)
}

func (c *Client) Read(bodyChan chan []byte) {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()

	for {
		messageType, p, err := c.Connection.ReadMessage()
		if err != nil {
			c.Pool.Unregister <- c
			c.Connection.Close()
			break
		}
		var body Body
		err = json.Unmarshal(p, &body)
		if err != nil {
			c.Pool.Unregister <- c
			c.Connection.Close()
			break
		}
		body.UserID = c.Email
		message := Message{
			Type: messageType,
			Body: body,
		}
		c.Pool.Broadcast <- message
		bodyChan <- p
		// save to db

	}
}
