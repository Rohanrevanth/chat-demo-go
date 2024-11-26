package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/Rohanrevanth/chat-demo-go/models"
)

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // Broadcast channel

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message defines a simple message structure
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// HandleWebSocket upgrades the HTTP connection to a WebSocket and handles communication
func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket Read Error: %v", err)
			delete(clients, ws)
			break
		}
		fmt.Println(msg)
		broadcast <- msg
	}
}

// HandleMessages listens for new messages and broadcasts them to connected clients
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("WebSocket Write Error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func SendNotificationToClients(message models.SocketMessage) {
	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("WebSocket SendNotificationToClients Error: %v", err)
			client.Close()
			delete(clients, client) // Remove disconnected client
		}
	}
}

func PingClients() {
	message := models.Message{
		Message: "ping",
	}
	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("WebSocket PingClients Error: %v", err)
			client.Close()
			delete(clients, client) // Remove disconnected client
		}
	}
}
