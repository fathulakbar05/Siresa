package routes

import (
    "github.com/gin-gonic/gin"
    "go-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
    reservasi := router.Group("/reservasi")
    {
        reservasi.POST("/ajukan", controllers.AjukanReservasi)
    }

    ruangan := router.Group("/ruangan")
    {
        ruangan.GET("/", controllers.GetAllRuangan)
    }
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	reservasi.PUT("/status", controllers.UpdateStatusReservasi)
	reservasi.GET("/by-status", controllers.GetReservasiByStatus)
	reservasi.GET("/by-user/:id_user", controllers.GetReservasiByUser)
}
