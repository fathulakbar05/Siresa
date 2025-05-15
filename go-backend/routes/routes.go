package routes

import (
	"go-backend/controllers"
	"go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rute Reservasi (Dimodifikasi)
	reservasi := router.Group("/reservasi")
	{
		reservasi.POST("/ajukan", controllers.AjukanReservasi)
		reservasi.PUT("/status", controllers.UpdateStatusReservasi) // Hapus atau komentari rute ini
		reservasi.GET("/by-status", controllers.GetReservasiByStatus)
		reservasi.GET("/by-user/:id_user", controllers.GetReservasiByUser)
		reservasi.PUT("/setujui/:id_reservasi", middleware.AdminOnly(), controllers.SetujuiReservasi) // Hapus atau komentari rute ini
		reservasi.PUT("/tolak/:id_reservasi", middleware.AdminOnly(), controllers.TolakReservasi)   // Hapus atau komentari rute ini
		reservasi.PUT("/:aksi/:id_reservasi", middleware.AdminOnly(), controllers.UpdateStatusReservasi) // Rute baru
	}

	// Rute Ruangan (Tidak Berubah)
	ruangan := router.Group("/ruangan")
	{
		ruangan.GET("/", controllers.GetAllRuangan)
		ruangan.GET("/filter", controllers.FilterRuangan)
		ruangan.GET("/:id", controllers.GetDetailRuangan)
	}

	// Rute Auth (Tidak Berubah)
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	// Rute Admin (Ditambahkan)
	admin := router.Group("/api/admin")
	{
		admin.GET("/ringkasan", middleware.AdminOnly(), controllers.GetRingkasanSistem)
	}
}