package controllers

import (
    "net/http"
    "time"
    "strconv"

    "github.com/gin-gonic/gin"
    "go-backend/config"
    "go-backend/models"
)

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

func UpdateStatusReservasi(c *gin.Context) {
    var input struct {
        IDReservasi uint   `json:"id_reservasi"`
        Status      string `json:"status"` // "disetujui" atau "ditolak"
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var reservasi models.Reservasi
    if err := config.DB.First(&reservasi, input.IDReservasi).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reservasi tidak ditemukan"})
        return
    }

    reservasi.Status = input.Status
    config.DB.Save(&reservasi)

    // Update notifikasi
    var notif models.Notifikasi
    config.DB.Where("id_reservasi = ?", input.IDReservasi).First(&notif)
    notif.StatusReservasi = input.Status
    config.DB.Save(&notif)

    c.JSON(http.StatusOK, gin.H{"message": "Status reservasi diperbarui"})
}

func GetReservasiByStatus(c *gin.Context) {
    status := c.Query("status")
    var data []models.Reservasi

    if err := config.DB.Where("status = ?", status).Find(&data).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, data)
}

func GetReservasiByUser(c *gin.Context) {
    idUser := c.Param("id_user")
    var data []models.Reservasi

    if err := config.DB.Where("id_user = ?", idUser).Find(&data).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, data)
}

func SetujuiReservasi(c *gin.Context) {
  idReservasiStr := c.Param("id_reservasi")
  idReservasi, err := strconv.Atoi(idReservasiStr)
  if err != nil {
   c.JSON(http.StatusBadRequest, gin.H{"error": "ID Reservasi tidak valid"})
   return
  }

  var reservasi models.Reservasi
  if err := config.DB.First(&reservasi, idReservasi).Error; err != nil {
   c.JSON(http.StatusNotFound, gin.H{"error": "Reservasi tidak ditemukan"})
   return
  }

  reservasi.Status = "disetujui"
  config.DB.Save(&reservasi)

  var notif models.Notifikasi
  config.DB.Where("id_reservasi = ?", idReservasi).First(&notif)
  notif.StatusReservasi = "disetujui"
  config.DB.Save(&notif)

  c.JSON(http.StatusOK, gin.H{"message": "Reservasi disetujui"})
 }

 func TolakReservasi(c *gin.Context) {
  idReservasiStr := c.Param("id_reservasi")
  idReservasi, err := strconv.Atoi(idReservasiStr)
  if err != nil {
   c.JSON(http.StatusBadRequest, gin.H{"error": "ID Reservasi tidak valid"})
   return
  }

  var reservasi models.Reservasi
  if err := config.DB.First(&reservasi, idReservasi).Error; err != nil {
   c.JSON(http.StatusNotFound, gin.H{"error": "Reservasi tidak ditemukan"})
   return
  }

  reservasi.Status = "ditolak"
  config.DB.Save(&reservasi)

  // Update notifikasi (jika perlu)
  var notif models.Notifikasi
  config.DB.Where("id_reservasi = ?", idReservasi).First(&notif)
  notif.StatusReservasi = "ditolak"
  config.DB.Save(&notif)

  c.JSON(http.StatusOK, gin.H{"message": "Reservasi ditolak"})
 }