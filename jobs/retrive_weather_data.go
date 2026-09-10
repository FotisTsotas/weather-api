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

type openMeteoResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Humidity    float64 `json:"relative_humidity_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current"`
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

		err = SaveWeatherData(city.ID, data)
		if err != nil {
			slog.Error("Failed to save weather data for city", "city_id", city.ID, "error", err)
			continue
		}

		if err := SaveWeatherHistory(city.ID, json.RawMessage(data)); err != nil {
			slog.Error("Failed to save weather history for city", "city_id", city.ID, "error", err)
			continue
		}
		slog.Info("Retrieved weather data for city", "city_id", city.ID, "data", json.RawMessage(data))
	}
}

func SaveWeatherData(cityID int, data json.RawMessage) error {
	retriveData := NewRetriveWeatherData()
	var response openMeteoResponse
	if err := json.Unmarshal(data, &response); err != nil {
		slog.Error("Failed to parse weather data", "city_id", cityID, "err", err)
		return err
	}

	condition := repositories.WeatherCodeToCondition(response.Current.WeatherCode)
	temperature := response.Current.Temperature
	err := retriveData.weather.SaveWeatherData(cityID, condition, temperature)
	if err != nil {
		slog.Error("Failed to save weather data for city", "city_id", cityID, "error", err)
		return err
	}
	return nil
}

func SaveWeatherHistory(cityID int, data json.RawMessage) error {
	retriveData := NewRetriveWeatherData()
	var response openMeteoResponse
	if err := json.Unmarshal(data, &response); err != nil {
		slog.Error("Failed to parse weather history data", "city_id", cityID, "error", err)
		return err
	}

	_, err := retriveData.weather.SaveWeatherHistory(cityID, response.Current.Temperature, response.Current.Humidity, response.Current.WindSpeed)
	if err != nil {
		slog.Error("Failed to save weather history for city", "city_id", cityID, "error", err)
		return err
	}
	return nil
}
