package routes

import (
	"weather-api/handlers"
	"weather-api/middlewares"

	"github.com/gin-gonic/gin"
)

func WeatherRoutes(server *gin.Engine, h *handlers.WeatherHandler) {
	
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/weather/:city_id", h.GetWeather)
	authenticated.POST("/weather", h.CreateWeather)
	authenticated.PUT("/weather/:city_id", h.UpdateWeather)
}
