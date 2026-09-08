package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"weather-api/repositories"
	"weather-api/services"

	"strconv"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	weatherService *services.WeatherService
}

func NewWeatherHandler(weatherService *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{weatherService: weatherService}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	city_id := c.Param("city_id")

	cityID, err := strconv.Atoi(city_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	weather, err := h.weatherService.GetWeather(cityID)
	if err != nil {
		if errors.Is(err, repositories.ErrCityNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get weather"})
		return
	}
	c.JSON(http.StatusOK, weather)
}

func (h *WeatherHandler) CreateWeather(c *gin.Context) {
	var req struct {
		CityID       int     `json:"city_id" binding:"required"`
		Condition    string  `json:"condition" binding:"required"`
		TemperatureC float64 `json:"temperature_c" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	weather, err := h.weatherService.CreateWeather(req.CityID, req.Condition, req.TemperatureC)
	if err != nil {
		if errors.Is(err, repositories.ErrCityNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create weather"})
		return
	}
	c.JSON(http.StatusOK, weather)
}


func (h *WeatherHandler) UpdateWeather(c *gin.Context) {
	city_id := c.Param("city_id")

	cityID, err := strconv.ParseInt(city_id, 10, 64)
	if err != nil {
		slog.Error("Invalid city ID", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	var req struct {
		CityID       int     `json:"city_id" binding:"required"`
		Condition    string  `json:"condition" binding:"required"`
		TemperatureC float64 `json:"temperature_c" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	weather, err := h.weatherService.UpdateWeather(int(cityID), req.CityID, req.Condition, req.TemperatureC)
	if err != nil {
		if errors.Is(err, repositories.ErrCityNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update weather"})
		return
	}
	c.JSON(http.StatusOK, weather)
}