package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

			for client, _ := range pool.Clients {
				client.Coon.WriteJSON(Message{Type: 1, Body: "New User joined chat"})
				client.Coon.WriteJSON(Message{Type: 1, Body: fmt.Sprintf("client %s", client.ID)})

			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)

			for client, _ := range pool.Clients {
				client.Coon.WriteJSON(Message{Type: 1, Body: "User left chat"})
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
