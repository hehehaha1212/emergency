package main

import (
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"emergency-app/configs"
	"emergency-app/internal/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.ConnectDB()
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Println("Starting server on port:", port)
	router.Run(":" + port)
}