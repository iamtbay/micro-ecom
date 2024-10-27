package main

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	//ch        *amqp091.Channel
	clients   = make(map[*websocket.Conn]bool)
	upgrader  = websocket.Upgrader{}
	mu        sync.Mutex
	secretKey = []byte("scretkey")
)

func main() {
	r := gin.Default()

	connectRabbitMQ()
	r.GET("/ws", handleWebSocket)

	log.Println("Starting http server on :8086")

	r.Run(":8086")

}

// ws
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection")
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
	}()

	for {
		_, _, err := conn.NextReader()
		if err != nil {
			break
		}
	}
}
