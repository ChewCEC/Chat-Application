package websocket

import (
	"fmt"
	"time"

	"github.com/ChewCEC/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan models.Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan models.Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			ID := primitive.NewObjectID()
			for client, _ := range pool.Clients {
				client.Coon.WriteJSON(
					models.Message{
						ID:        ID,
						Content:   "New User Joined...",
						Timestamp: time.Now()})

			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)

			for client, _ := range pool.Clients {
				client.Coon.WriteJSON(
					models.Message{
						ID:        client.ID,
						Content:   "User Disconnected...",
						Timestamp: time.Now()})
				fmt.Printf("%d users connected\n", len(pool.Clients))
			}

		case message := <-pool.Broadcast:
			for client, _ := range pool.Clients {
				err := client.Coon.WriteJSON(message)
				if err != nil {
					fmt.Println(err)
					return
				}

			}
		}
	}
}
