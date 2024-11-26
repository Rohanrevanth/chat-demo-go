package http

import (
	"github.com/Rohanrevanth/chat-demo-go/routes"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/gin-contrib/cors"
)

// InitRouter initializes the Gin router and registers routes.
func InitRouter() *gin.Engine {
	router := gin.Default()

	// Add middleware (e.g., logging, recovery)
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Frontend URL
		// AllowOrigins:     []string{"https://fresh-cosmic-garfish.ngrok-free.app"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // If you're using cookies or authentication
		MaxAge:           12 * time.Hour, // Preflight request caching
	}))

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://da90-2401-4900-482e-d04c-d879-e020-c080-f1e8.ngrok-free.app"}, // Frontend URL without port
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	// Register application routes
	routes.RegisterRoutes(router)

	return router
}

// StartServer starts the HTTP server on the specified address.
func StartServer() {
	router := InitRouter()

	// router.Run("localhost:8080") // This can be configurable
	router.Run("0.0.0.0:8080") // This can be configurable
}
