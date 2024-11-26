package routes

import (
	"net/http"

	"github.com/Rohanrevanth/chat-demo-go/auth"
	"github.com/Rohanrevanth/chat-demo-go/controllers"
	websocket "github.com/Rohanrevanth/chat-demo-go/websockets"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the application's routes
func RegisterRoutes(router *gin.Engine) {

	//User routes
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/publicusers", controllers.GetPublicUsers)
	router.POST("/addconversation", controllers.AddNewConversation)
	router.POST("/getconversations", controllers.GetUserConversations)
	router.POST("/getconversation", controllers.GetUserConversation)
	router.POST("/deleteconversation", controllers.DeleteUserConversation)
	router.POST("/message", controllers.AddNewMessage)
	router.POST("/getfile", controllers.GetFile)
	// router.POST("/searchuser", controllers.GetAllUsers)

	// WebSocket route
	router.GET("/ws", websocket.HandleWebSocket) // WebSocket endpoint

	router.GET("/ping", func(c *gin.Context) {
		websocket.PingClients() // Send ping to all clients
		c.JSON(http.StatusOK, gin.H{"message": "Ping sent to clients"})
	})

	// Protected routes
	protected := router.Group("/").Use(auth.JWTAuthMiddleware())
	{
		protected.GET("/users", controllers.GetAllUsers)
	}
}
