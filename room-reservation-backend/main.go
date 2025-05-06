package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"room-reservation-backend/config"
	"room-reservation-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}