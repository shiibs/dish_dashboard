package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shiibs/nosh_dashboard/database"
	"github.com/shiibs/nosh_dashboard/models"
)

// GetDishes handles GET requests to retrieve all dishes from the database.
func GetDishes(c *gin.Context) {
	var dishes []models.Dish

	// Retrieve all dishes from the database.
	if err := database.DBConn.Find(&dishes).Error; err != nil {
		// If there's an error, respond with an internal server error and the error message.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the list of dishes in JSON format.
	c.JSON(http.StatusOK, dishes)
}
