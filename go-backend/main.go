package main

import (
	"go-backend/config"
	"go-backend/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    config.InitDB()

    r := gin.Default()
    r.Static("/static", "./static")

    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })

    // Konfigurasi CORS
    configCors := cors.DefaultConfig()
    configCors.AllowOrigins = []string{"http://127.0.0.1:5500", "http://localhost:5500", "http://localhost:8080", "http://127.0.0.1:3000", "http://localhost:3000"}
    configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "X-Requested-With", "Authorization"} // Tambahkan "Authorization" di sini
    configCors.AllowCredentials = true
    corsMiddleware := cors.New(configCors)
    log.Println("CORS Middleware dibuat:", corsMiddleware)
    r.Use(corsMiddleware)

    routes.SetupRoutes(r)
    r.Run(":8080")
}
