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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500", "http://localhost:5500", "http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "X-Requested-With"}
	config.AllowCredentials = true
	corsMiddleware := cors.New(config)
	log.Println("CORS Middleware dibuat:", corsMiddleware)
	r.Use(corsMiddleware)

	routes.SetupRoutes(r)
	r.Run(":8080")
}
