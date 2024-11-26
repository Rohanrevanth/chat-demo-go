package main

import (
	"github.com/Rohanrevanth/chat-demo-go/database"
	"github.com/Rohanrevanth/chat-demo-go/http"
	// websocket "github.com/Rohanrevanth/chat-demo-go/websockets"
)

func main() {
	// Get a greeting message and print it.
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// database.ConnectToDB()
	database.ConnectDatabase()

	// go websocket.HandleMessages()

	http.StartServer()
}
