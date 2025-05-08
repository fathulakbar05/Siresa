package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "go-backend/config"
    "go-backend/models"
)

func Login(c *gin.Context) {
    var input struct {
        NIM_NIP   string `json:"nim_nip"`
        Password  string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("nim_nip = ?", input.NIM_NIP).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login berhasil",
        "user":    user,
    })
}
