package main

import (
	"log"
	"weather-api/config"
	"weather-api/db"
	"weather-api/handlers"
	"weather-api/repositories"
	"weather-api/routes"
	"weather-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg := config.Load()

	db.InitDB(cfg)

	userRepo := repositories.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	userHandler := handlers.NewUserHandler(authService)

	server := gin.Default()
	routes.RegisterRoutes(server, userHandler)

	server.Run(":" + cfg.ServerPort)
}
