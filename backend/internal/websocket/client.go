package websocket

import (
	"fmt"
	"time"
	"github.com/ChewCEC/chat-app/internal/db"
	
	"github.com/gorilla/websocket"
	"github.com/ChewCEC/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID   primitive.ObjectID
	Coon *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Coon.Close()
	}()

	for {
		messageType, p, err := c.Coon.ReadMessage()
		if messageType !=1 {
			fmt.Errorf("Bad message type: %s", err)
		}

		if err != nil {
			fmt.Errorf("error: %s", err)
			return
		}
		
        message := models.Message{
            ID: primitive.NewObjectID(), // Add proper ID generation
            Content: string(p), 
            Timestamp: time.Now(),
        }

        // Save message to DB before broadcasting
        if err := db.SaveMessage(message); err != nil {
            fmt.Printf("Error saving message: %v\n", err)
            continue
        }

		c.Pool.Broadcast <- message

	}
}
