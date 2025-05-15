package routes
 
 import (
  "github.com/gin-gonic/gin"
  "go-backend/controllers"
 )
 
 func SetupRoutes(router *gin.Engine) {
  // Rute Reservasi
  reservasi := router.Group("/reservasi")
  {
   reservasi.POST("/ajukan", controllers.AjukanReservasi)
   reservasi.PUT("/status", controllers.UpdateStatusReservasi)
   reservasi.GET("/by-status", controllers.GetReservasiByStatus)
   reservasi.GET("/by-user/:id_user", controllers.GetReservasiByUser)
  }
 
  // Rute Ruangan
  ruangan := router.Group("/ruangan")
  {
   ruangan.GET("/", controllers.GetAllRuangan)
  }
 
  // Rute Auth
  auth := router.Group("/auth")
{
      auth.POST("/login", controllers.Login)
      auth.POST("/register", controllers.Register)
}
}