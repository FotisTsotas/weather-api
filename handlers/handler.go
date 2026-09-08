package handlers

import (
	"database/sql"
	"weather-api/repositories"
	"weather-api/services"
)

func InitUserHandler(db *sql.DB) *UserHandler {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	userHandler := NewUserHandler(authService)

	return userHandler
}

func InitWeatherHandler(db *sql.DB) *WeatherHandler {
	weatherRepo := repositories.NewWeatherRepository(db)
	weatherService := services.NewWeatherService(weatherRepo)
	weatherHandler := NewWeatherHandler(weatherService)

	return weatherHandler
}

func InitCityHandler(db *sql.DB) *CityHandler {
	cityRepo := repositories.NewCityRepository(db)
	cityService := services.NewCityService(cityRepo)
	cityHandler := NewCityHandler(cityService)

	return cityHandler
}
