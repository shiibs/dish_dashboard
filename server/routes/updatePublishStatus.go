package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shiibs/nosh_dashboard/database"
	"github.com/shiibs/nosh_dashboard/models"
)

// ToggleDishStatus handles PUT requests to toggle the publication status of a dish.
func ToggleDishStatus(c *gin.Context) {
	// Get the dish ID from the URL parameters.
	id := c.Param("id")

	var dish models.Dish
	// Find the dish by ID in the database.
	if err := database.DBConn.First(&dish, id).Error; err != nil {
		// If there's an error (e.g., dish not found), respond with an internal server error and the error message.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Toggle the publication status of the dish.
	dish.IsPublished = !dish.IsPublished

	// Save the updated dish back to the database.
	database.DBConn.Save(&dish)

	// Broadcast the updated dish
	broadcast <- dish

	// Respond with the updated dish
	c.JSON(http.StatusOK, dish)
}
