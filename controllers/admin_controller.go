package controllers

import (
	"net/http"

	"go-backend/config"
	"go-backend/models"

	"github.com/gin-gonic/gin"
)

// GetRingkasanSistem mengembalikan ringkasan dashboard admin
func GetRingkasanSistem(c *gin.Context) {
	var totalRuangan int64
	var totalMenunggu int64
	var totalDisetujui int64
	var totalDitolak int64

	// Hitung total ruangan
	config.DB.Model(&models.Ruangan{}).Count(&totalRuangan)

	// Hitung jumlah reservasi berdasarkan status
	config.DB.Model(&models.Reservasi{}).Where("status = ?", "menunggu").Count(&totalMenunggu)
	config.DB.Model(&models.Reservasi{}).Where("status = ?", "disetujui").Count(&totalDisetujui)
	config.DB.Model(&models.Reservasi{}).Where("status = ?", "ditolak").Count(&totalDitolak)

	c.JSON(http.StatusOK, gin.H{
		"permintaan_baru": totalMenunggu,
		"disetujui":       totalDisetujui,
		"ditolak":         totalDitolak,
		"total_ruangan":   totalRuangan,
	})
}
