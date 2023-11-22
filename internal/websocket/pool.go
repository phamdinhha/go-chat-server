package websocket

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Message struct {
	Type     int `json:"Type,omitempty"`
	Body     Body
	ClientID string `json:"client_id"`
}

type Body struct {
	RoomID  uuid.UUID `json:"room_id"`
	Message string    `json:"message"`
	UserID  string    `json:"user_id"`
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Broadcast  chan Message
	clientsMux *sync.RWMutex
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan Message),
		clientsMux: &sync.RWMutex{},
	}
}

func (p *Pool) AddClient(client *Client) {
	p.clientsMux.Lock()
	defer p.clientsMux.Unlock()
	p.Clients[client.ID] = client
	fmt.Println("adding client: ", p.Clients)
}

func (p *Pool) RemoveClient(client *Client) {
	p.clientsMux.Lock()
	defer p.clientsMux.Unlock()
	delete(p.Clients, client.ID)
}

func (p *Pool) GetClientByID(clientID string) *Client {
	p.clientsMux.Lock()
	defer p.clientsMux.Unlock()
	return p.Clients[clientID]
}

func (p *Pool) BroadcastMessage(message Message) {
	p.clientsMux.Lock()
	defer p.clientsMux.Unlock()
	fmt.Println("client: ", p.Clients)
	var targetClients []*Client
	for _, client := range p.Clients {
		if client.RoomID == message.Body.RoomID && client.ID != message.ClientID {
			targetClients = append(targetClients, client)
		}
	}
	for _, client := range targetClients {
		client.Send(message)
	}
}

func (p *Pool) Start() {
	// TODO: recover here
	for {
		select {
		case client := <-p.Register:
			fmt.Println("registering client")
			fmt.Println(client.ID)
			p.AddClient(client)
		case client := <-p.Unregister:
			p.RemoveClient(client)
		case message := <-p.Broadcast:
			p.BroadcastMessage(message)
		}
	}
}
