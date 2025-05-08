package main

import (
    "github.com/gin-gonic/gin"
    "go-backend/config"      // Tambahkan ini
    "go-backend/controllers"
)

func main() {
    config.InitDB() // Tambahkan ini di awal

    r := gin.Default()
    r.POST("/reservasi/ajukan", controllers.AjukanReservasi)
    r.Run(":8080")
}
