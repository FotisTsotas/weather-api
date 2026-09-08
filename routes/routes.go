package routes

import (
	"database/sql"

	"weather-api/handlers"
	"weather-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRoutes(server *gin.Engine, db *sql.DB) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	AuthRoutes(server, handlers.InitUserHandler(db))
	CityRoutes(server, authenticated, handlers.InitCityHandler(db))
	WeatherRoutes(server, authenticated, handlers.InitWeatherHandler(db))
}

func CityRoutes(
	server *gin.Engine,
	authenticated *gin.RouterGroup,
	h *handlers.CityHandler,
) {
	authenticated.GET("/cities", h.GetCities)
	authenticated.GET("/cities/:id", h.GetCityByID)
	authenticated.POST("/cities", h.CreateCity)
	authenticated.PUT("/cities/:id", h.UpdateCity)
	authenticated.DELETE("/cities/:id", h.DeleteCity)
}

func WeatherRoutes(
	server *gin.Engine,
	authenticated *gin.RouterGroup,
	h *handlers.WeatherHandler,
) {
	authenticated.GET("/weather/:city_id", h.GetWeather)
	authenticated.POST("/weather", h.CreateWeather)
	authenticated.PUT("/weather/:city_id", h.UpdateWeather)
	authenticated.DELETE("/weather/:city_id", h.DeleteWeather)
}

func AuthRoutes(server *gin.Engine, h *handlers.UserHandler) {
	server.POST("/signup", h.Signup)
	server.POST("/login", h.Login)
}
