package main

import (
	"log/slog"
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
		slog.Warn("No .env file found, relying on environment variables")
	}

	cfg := config.Load()
	config.InitLogger()

	db.InitDB(cfg)

	userRepo := repositories.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	userHandler := handlers.NewUserHandler(authService)

	cityRepo := repositories.NewCityRepository(db.DB)
	cityService := services.NewCityService(cityRepo)
	cityHandler := handlers.NewCityHandler(cityService)

	weatherRepo := repositories.NewWeatherRepository(db.DB)
	weatherService := services.NewWeatherService(weatherRepo)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	server := gin.Default()
	routes.AuthRoutes(server, userHandler)
	routes.CityRoutes(server, cityHandler)
	routes.WeatherRoutes(server, weatherHandler)

	server.Run(":" + cfg.ServerPort)
}
