package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Coon *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Coon.Close()
	}()

	for {
		messageType, p, err := c.Coon.ReadMessage()
		if err != nil {
			fmt.Errorf("error: %s", err)
		}

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
		fmt.Printf("Tamano de pool: %d\n", len(c.Pool.Clients))

	}
}
