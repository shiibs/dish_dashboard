package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shiibs/nosh_dashboard/database"
	"github.com/shiibs/nosh_dashboard/routes"
)

// init initializes the environment and database connection.
func init() {
	// Load environment variables from .env file.
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file error:", err)
	}

	// Connect to the database.
	database.ConnectDB()
}

func main() {
	// Create a new Gin router.
	router := gin.Default()

	// Set up CORS middleware.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},  // Allow requests from this origin.
		AllowMethods:     []string{"GET", "PUT"},             // Allow only GET and PUT methods.
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Allow these headers in requests.
		ExposeHeaders:    []string{"Content-Length"},         // Expose these headers in responses.
		AllowCredentials: true,                               // Allow credentials such as cookies.
	}))

	// Define routes for handling dishes.
	router.GET("/dishes", routes.GetDishes)                   // GET route to retrieve dishes.
	router.PUT("/dishes/toggle/:id", routes.ToggleDishStatus) // PUT route to toggle dish status.

	// Handle WebSocket connections.
	go routes.HandleMessages() // Handle incoming messages in a separate goroutine.

	router.GET("/ws", func(c *gin.Context) {
		routes.HandleConnections(c) // Handle WebSocket connections.
	})

	// Start the server on port 8005.
	router.Run(":8005")
}
