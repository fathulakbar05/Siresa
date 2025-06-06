package controllers

import (
	"net/http"
	"strconv"

	"go-backend/config"
	"go-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllRuangan(c *gin.Context) {
	var ruangan []models.Ruangan
	if err := config.DB.Find(&ruangan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ruangan})
}

func FilterRuangan(c *gin.Context) {
	// Ambil parameter dari query string
	tanggal := c.Query("tanggal")
	jamMulai := c.Query("jam_mulai")
	jamSelesai := c.Query("jam_selesai")
	gedung := c.Query("gedung")
	kapasitasStr := c.Query("kapasitas")

	// Validasi parameter (Tambahkan validasi yang lebih kuat)
	if tanggal == "" || jamMulai == "" || jamSelesai == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tanggal, jam_mulai, dan jam_selesai harus diisi"})
		return
	}

	var ruangan []models.Ruangan
	query := config.DB.Model(&models.Ruangan{})

	// Filter by Gedung
	if gedung != "" {
		query = query.Where("lokasi_gedung = ?", gedung)
	}

	// Filter by Kapasitas
	if kapasitasStr != "" {
		kapasitas, err := strconv.Atoi(kapasitasStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Kapasitas tidak valid"})
			return
		}
		query = query.Where("kapasitas >= ?", kapasitas)
	}

	// Filter by Tanggal dan Waktu (Ketersediaan)
	// Logika kompleks untuk menghindari tumpang tindih
	query = query.Where("id_ruangan NOT IN (?)",
		config.DB.Model(&models.Reservasi{}).
			Where("tanggal = ?", tanggal).
			Where("(jam_mulai < ? AND jam_selesai > ?) OR (jam_mulai < ? AND jam_selesai < ?) OR (jam_mulai > ? AND jam_mulai < ? AND jam_selesai > ?)",
				jamSelesai, jamMulai, jamMulai, jamSelesai, jamMulai, jamSelesai, jamSelesai).
			Select("id_ruangan"),
	)

	if err := query.Find(&ruangan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ruangan})
}

// GetRuanganByID mengambil detail ruangan berdasarkan ID
func GetDetailRuangan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID ruangan tidak valid"})
		return
	}

	var ruangan models.Ruangan
	if err := config.DB.Where("id_ruangan = ?", uint(id)).First(&ruangan).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ruangan tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data ruangan: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ruangan)
}
