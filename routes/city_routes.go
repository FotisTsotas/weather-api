package routes

import (
	"weather-api/handlers"
	"weather-api/middlewares"

	"github.com/gin-gonic/gin"
)

func CityRoutes(server *gin.Engine, h *handlers.CityHandler) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Define city routes here
	authenticated.GET("/cities", h.GetCities)
	authenticated.GET("/cities/:id", h.GetCityByID)
	authenticated.POST("/cities", h.CreateCity)
	authenticated.PUT("/cities/:id", h.UpdateCity)
	authenticated.DELETE("/cities/:id", h.DeleteCity)
}
