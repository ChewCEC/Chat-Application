package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ChewCEC/chat-app/internal/db"
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
		fmt.Errorf("❌ Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		fmt.Errorf("❌ MONGO_URI not set in .env")
		return
	}

	fmt.Printf("Mongo URI: %s\n", mongoURI)
	
	db.InitMongo(mongoURI)
	

	pool := websocket.NewPool()
	go pool.Start()

	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		serveWs(pool, c) // ✅ Use the shared pool for all clients
	})

	router.Run(":8080") // This will listen on 0.0.0.0:8080

}
