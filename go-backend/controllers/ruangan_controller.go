package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go-backend/config"
    "go-backend/models"
)

func GetAllRuangan(c *gin.Context) {
    var ruangan []models.Ruangan
    if err := config.DB.Find(&ruangan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": ruangan})
}
