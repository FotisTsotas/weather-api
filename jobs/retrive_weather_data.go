package jobs

import (
	"encoding/json"
	"log/slog"
	"weather-api/clients"
	"weather-api/db"
	"weather-api/repositories"
)

type RetriveWeatherData struct {
	cities  *repositories.CityRepository
	client  *clients.WeatherClient
	weather *repositories.WeatherRepository
}

func NewRetriveWeatherData() *RetriveWeatherData {
	return &RetriveWeatherData{
		cities:  repositories.NewCityRepository(db.DB),
		client:  clients.NewWeatherClient(),
		weather: repositories.NewWeatherRepository(db.DB),
	}
}

func HandleWeatherDataRetrieval() {
	slog.Info("Retrieving weather data...")
	retriveData := NewRetriveWeatherData()
	cities, err := retriveData.cities.GetCities()
	if err != nil {
		slog.Error("Failed to retrieve cities", "error", err)
		return
	}

	for _, city := range cities {
		data, err := retriveData.client.GetWeatherData(city)
		if err != nil {
			slog.Error("Failed to retrieve weather data for city", "city_id", city.ID, "error", err)
			continue
		}

		err = retriveData.weather.SaveWeatherData(city.ID, data)
		if err != nil {
			slog.Error("Failed to save weather data for city", "city_id", city.ID, "error", err)
			continue
		}
		slog.Info("Retrieved weather data for city", "city_id", city.ID, "data", json.RawMessage(data))
	}
}
