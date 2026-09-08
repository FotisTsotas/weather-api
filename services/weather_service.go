package services

import (
	"weather-api/models"
	"weather-api/repositories"
)

type WeatherService struct {
	weatherRepo *repositories.WeatherRepository
}

func NewWeatherService(weatherRepo *repositories.WeatherRepository) *WeatherService {
	return &WeatherService{weatherRepo: weatherRepo}
}

func (s *WeatherService) GetWeather(cityID int) (*models.Weather, error) {
	return s.weatherRepo.FindByCity(cityID)
}

func (s *WeatherService) CreateWeather(cityID int, condition string, temperatureC float64) (*models.Weather, error) {
	return s.weatherRepo.Create(cityID, condition, temperatureC)
}

func (s *WeatherService) UpdateWeather(cityID int, newCityID int, condition string, temperatureC float64) (*models.Weather, error) {
	return s.weatherRepo.Update(cityID, newCityID, condition, temperatureC)
}
