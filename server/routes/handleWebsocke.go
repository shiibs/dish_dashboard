package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shiibs/nosh_dashboard/database"
	"github.com/shiibs/nosh_dashboard/models"
)

// clients keeps track of all connected WebSocket clients.
var clients = make(map[*websocket.Conn]bool)

// broadcast is a channel used to send updated dishes to all clients.
var broadcast = make(chan models.Dish)

// upgrader upgrades HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections.
		return true
	},
}

// HandleConnections upgrades the HTTP connection to a WebSocket connection and manages.
func HandleConnections(c *gin.Context) {
	// Upgrade the HTTP connection to a WebSocket connection.
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer ws.Close()

	// Register the new client.
	clients[ws] = true

	for {
		var dish models.Dish
		// Read JSON message from WebSocket client.
		err := ws.ReadJSON(&dish)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			delete(clients, ws)
			break
		}

		// Update the dish in the database.
		if err := database.DBConn.Model(&models.Dish{}).Where(
			"id = ?", dish.ID,
		).Updates(models.Dish{IsPublished: dish.IsPublished}).Error; err != nil {
			log.Printf("Database update error: %v", err)
			continue
		}

		// Send the updated dish to the broadcast channel.
		broadcast <- dish
	}
}

// HandleMessages listens for messages on the broadcast channel and sends them to all connected clients.
func HandleMessages() {
	for {
		// Receive updated dish from the broadcast channel.
		dish := <-broadcast
		// Send the updated dish to all clients.
		for client := range clients {
			err := client.WriteJSON(dish)
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
