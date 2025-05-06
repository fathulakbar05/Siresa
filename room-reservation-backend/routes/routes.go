package routes

import (
	"github.com/gin-gonic/gin"
	"room-reservation-backend/handlers"
	"room-reservation-backend/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)

	auth := router.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())

	auth.POST("/reservations", handlers.CreateReservation)
	auth.GET("/rooms", handlers.GetAllRooms)
	auth.POST("/rooms", handlers.CreateRoom)
	auth.PUT("/rooms/:id", handlers.UpdateRoom)
	auth.DELETE("/rooms/:id", handlers.DeleteRoom)
	auth.GET("/admin/reservations", handlers.GetAllReservations)
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

}
