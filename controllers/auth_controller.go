package controllers

import (
	"net/http"

	"go-backend/config"
	"go-backend/models"
	"go-backend/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input struct {
		Nama     string `json:"nama" binding:"required"`
		NIM_NIP  string `json:"nim_nip" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password dengan SHA256
	hashedPassword := utils.HashPasswordSHA256(input.Password)

	// Buat user baru
	user := models.User{
		Nama:     input.Nama,
		NIM_NIP:  input.NIM_NIP,
		Password: hashedPassword, // Simpan hash SHA256, bukan bcrypt
		Role:     input.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "user": user})
}

func Login(c *gin.Context) {
	var input struct {
		NIM_NIP  string `json:"nim_nip"`
		Password string `json:"password"`
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

	if !utils.VerifyPasswordSHA256(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"user":    user,
	})
}
