package main

import (
	"fmt"

	"github.com/ChewCEC/chat-app/src/websocket"
	"github.com/gin-gonic/gin"
)

func serveWs(pool *websocket.Pool, c *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(c)
	if err != nil {
		fmt.Fprintf(c.Writer, "%+v\n", err)
	}

	client := &websocket.Client{
		Coon: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}


func main() {
	pool := websocket.NewPool()
	go pool.Start()
	
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		serveWs(pool, c) // ✅ Use the shared pool for all clients
	})

	router.Run("localhost:8080")

}
