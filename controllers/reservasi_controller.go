package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-backend/config"
	"go-backend/models"

	"github.com/gin-gonic/gin"
)

// Struktur untuk menampung data reservasi menunggu yang akan dikirim ke frontend
type ReservasiMenunggu struct {
	IDReservasi  int    `json:"id_reservasi"`
	NamaPeminjam string `json:"nama_peminjam"`
	IDRuangan    int    `json:"id_ruangan"`
	Tanggal      string `json:"tanggal"`
	JamMulai     string `json:"jam_mulai"`
	JamSelesai   string `json:"jam_selesai"`
}

// AjukanReservasi menangani pengajuan reservasi baru
func AjukanReservasi(c *gin.Context) {
	var input models.Reservasi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing []models.Reservasi
	err := config.DB.Where("id_ruangan = ? AND tanggal = ? AND status = 'disetujui' AND jam_mulai < ? AND jam_selesai > ?",
		input.IDRuangan, input.Tanggal, input.JamSelesai, input.JamMulai).Find(&existing).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(existing) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Waktu bentrok dengan reservasi lain"})
		return
	}

	input.Status = "menunggu"
	input.WaktuPengajuan = time.Now()

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Buat notifikasi
	notif := models.Notifikasi{
		IDUser:          input.IDUser,
		IDRuangan:       input.IDRuangan,
		IDReservasi:     input.IDReservasi,
		StatusReservasi: "pending",
		WaktuDibuat:     time.Now(),
	}
	config.DB.Create(&notif)

	c.JSON(http.StatusOK, gin.H{"message": "Reservasi diajukan", "data": input})
}

// UpdateStatusReservasi memperbarui status reservasi (setujui/tolak)
func UpdateStatusReservasi(c *gin.Context) {
	aksi := c.Param("aksi")
	idReservasiStr := c.Param("id_reservasi")

	idReservasi, err := strconv.Atoi(idReservasiStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Reservasi tidak valid"})
		return
	}

	var reservasi models.Reservasi // Menggunakan JadwalReservasi
	if err := config.DB.First(&reservasi, idReservasi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservasi tidak ditemukan"})
		return
	}

	// Validasi aksi
	if aksi != "setujui" && aksi != "tolak" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Aksi tidak valid"})
		return
	}

	// Update status
	reservasi.Status = aksi
	config.DB.Save(&reservasi)

	// Update notifikasi (jika perlu)
	var notif models.Notifikasi
	config.DB.Where("id_reservasi = ?", idReservasi).First(&notif)
	notif.StatusReservasi = aksi
	config.DB.Save(&notif)

	c.JSON(http.StatusOK, gin.H{"message": "Status reservasi diperbarui"})
}

// GetReservasiByStatus mengambil data reservasi berdasarkan status untuk dashboard admin
func GetReservasiByStatus(c *gin.Context) {
	status := c.Query("status")
	if status != "menunggu" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status tidak valid"})
		return
	}

	var reservasi []models.Reservasi // Menggunakan JadwalReservasi
	if err := config.DB.Where("status = ?", status).Find(&reservasi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []ReservasiMenunggu
	for _, r := range reservasi {
		var user models.User // Menggunakan model User
		if err := config.DB.First(&user, r.IDUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User tidak ditemukan untuk reservasi"})
			return
		}
		result = append(result, ReservasiMenunggu{
			IDReservasi:  int(r.IDReservasi),
			NamaPeminjam: user.Nama, // Ambil nama peminjam dari tabel user
			IDRuangan:    int(r.IDRuangan),
			Tanggal:      r.Tanggal,    // Sudah berupa string
			JamMulai:     r.JamMulai,   // Sudah berupa string
			JamSelesai:   r.JamSelesai, // Sudah berupa string
		})
	}

	c.JSON(http.StatusOK, result)
}

// GetReservasiByUser mengambil data reservasi berdasarkan ID pengguna
func GetReservasiByUser(c *gin.Context) {
	idUserStr := c.Param("id_user")
	idUser, err := strconv.Atoi(idUserStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID User tidak valid"})
		return
	}

	var reservasi []models.Reservasi // Menggunakan JadwalReservasi
	if err := config.DB.Where("id_user = ?", idUser).Find(&reservasi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservasi)
}

// SetujuiReservasi menyetujui reservasi
func SetujuiReservasi(c *gin.Context) {
	idReservasiStr := c.Param("id_reservasi")
	idReservasi, err := strconv.Atoi(idReservasiStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Reservasi tidak valid"})
		return
	}

	var reservasi models.Reservasi // Menggunakan JadwalReservasi
	if err := config.DB.First(&reservasi, idReservasi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservasi tidak ditemukan"})
		return
	}

	reservasi.Status = "disetujui"
	config.DB.Save(&reservasi)

	// Update notifikasi
	var notif models.Notifikasi
	config.DB.Where("id_reservasi = ?", idReservasi).First(&notif)
	notif.StatusReservasi = "disetujui"
	config.DB.Save(&notif)

	c.JSON(http.StatusOK, gin.H{"message": "Reservasi disetujui"})
}

// TolakReservasi menolak reservasi
func TolakReservasi(c *gin.Context) {
	idReservasiStr := c.Param("id_reservasi")
	idReservasi, err := strconv.Atoi(idReservasiStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Reservasi tidak valid"})
		return
	}

	var reservasi models.Reservasi // Menggunakan JadwalReservasi
	if err := config.DB.First(&reservasi, idReservasi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservasi tidak ditemukan"})
		return
	}

	reservasi.Status = "ditolak"
	config.DB.Save(&reservasi)

	// Update notifikasi
	var notif models.Notifikasi
	config.DB.Where("id_reservasi = ?", idReservasi).First(&notif)
	notif.StatusReservasi = "ditolak"
	config.DB.Save(&notif)

	c.JSON(http.StatusOK, gin.H{"message": "Reservasi ditolak"})
}
