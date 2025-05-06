package handlers

import (
	"room-reservation-backend/config"
	"room-reservation-backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllReservations(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var reservations []models.Reservation
	if err := config.DB.Preload("User").Preload("Room").Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservations})
}
