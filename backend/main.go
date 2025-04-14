package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ChewCEC/chat-app/internal/websocket"
)

func serveWs(pool *websocket.Pool, c *gin.Context) {
	conn, err := websocket.Upgrade(c)
	if err != nil {
		fmt.Errorf("%+v\n", err)
	}

	client := &websocket.Client{
		Coon: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func main() {

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	pool := websocket.NewPool()
	go pool.Start()

	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		serveWs(pool, c)
	})

	router.Run(":8080")

}
