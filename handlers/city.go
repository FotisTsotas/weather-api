package handlers

import (
	"strconv"
	"weather-api/services"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(service *services.CityService) *CityHandler {
	return &CityHandler{service: service}
}

func (h *CityHandler) GetCities(c *gin.Context) {
	cities, err := h.service.GetCities()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve cities"})
		return
	}
	c.JSON(200, cities)
}

func (h *CityHandler) GetCityByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid city ID"})
		return
	}

	city, err := h.service.GetCityByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve city"})
		return
	}
	if city == nil {
		c.JSON(404, gin.H{"error": "City not found"})
		return
	}
	c.JSON(200, city)
}

func (h *CityHandler) CreateCity(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	city, err := h.service.CreateCity(request.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create city"})
		return
	}
	c.JSON(201, city)
}

func (h *CityHandler) UpdateCity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid city ID"})
		return
	}

	var request struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	city, err := h.service.UpdateCity(id, request.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update city"})
		return
	}
	c.JSON(200, city)
}

func (h *CityHandler) DeleteCity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid city ID"})
		return
	}

	err = h.service.DeleteCity(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete city"})
		return
	}
	c.JSON(204, nil)
}
