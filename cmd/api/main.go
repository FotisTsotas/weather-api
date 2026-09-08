package main

import (
	"log/slog"
	"weather-api/config"
	"weather-api/db"
	"weather-api/routes"

	"weather-api/schedulers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, relying on environment variables")
	}

	cfg := config.Load()
	config.InitLogger()
	db.InitDB(cfg)
	scheduler := schedulers.InitScheduler()
	defer scheduler.Stop()

	server := gin.Default()
	routes.SetRoutes(server, db.DB)

	server.Run(":" + cfg.ServerPort)
}
