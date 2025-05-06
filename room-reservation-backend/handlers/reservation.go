package handlers

import (
	"room-reservation-backend/config"
	"room-reservation-backend/models"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context) {
	var input struct {
		RoomID    uint      `json:"room_id"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)

	reservation := models.Reservation{
		UserID:    userID,
		RoomID:    input.RoomID,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
	}

	if err := config.DB.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create reservation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": reservation})
}
