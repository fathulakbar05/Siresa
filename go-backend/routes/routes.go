package routes
 
 import (
  "github.com/gin-gonic/gin"
  "go-backend/controllers"
  "go-backend/middleware"
 )
 
 func SetupRoutes(router *gin.Engine) {
  // Rute Reservasi
  reservasi := router.Group("/reservasi")
  {
   reservasi.POST("/ajukan", controllers.AjukanReservasi)
   reservasi.PUT("/status", controllers.UpdateStatusReservasi)
   reservasi.GET("/by-status", controllers.GetReservasiByStatus)
   reservasi.GET("/by-user/:id_user", controllers.GetReservasiByUser)
   reservasi.PUT("/setujui/:id_reservasi", middleware.AdminOnly(), controllers.SetujuiReservasi)
   reservasi.PUT("/tolak/:id_reservasi", middleware.AdminOnly(), controllers.TolakReservasi)
  }
 
  // Rute Ruangan
ruangan := router.Group("/ruangan")
{
    ruangan.GET("/", controllers.GetAllRuangan)
    ruangan.GET("/filter", controllers.FilterRuangan) // Route baru untuk filter
    ruangan.GET("/:id", controllers.GetDetailRuangan)
}
 
  // Rute Auth
  auth := router.Group("/auth")
{
      auth.POST("/login", controllers.Login)
      auth.POST("/register", controllers.Register)
}

}
