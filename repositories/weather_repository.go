package repositories

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"
	"weather-api/models"
)

var ErrCityNotFound = errors.New("city not found")

type WeatherRepository struct {
	db *sql.DB
}

func NewWeatherRepository(db *sql.DB) *WeatherRepository {
	return &WeatherRepository{db: db}
}

func (r *WeatherRepository) FindByCity(city_id int) (*models.Weather, error) {
	weather := &models.Weather{}
	err := r.db.QueryRow("SELECT city_id, temperature_c, `condition` FROM weathers WHERE city_id = ?", city_id).Scan(&weather.City_ID, &weather.TemperatureC, &weather.Condition)
	slog.Info("Querying weather for city:", "city_id", city_id, "err", err)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("City not found", "city_id", city_id)
			return nil, ErrCityNotFound
		}
		return nil, err
	}
	return weather, nil
}

func (r *WeatherRepository) Create(cityID int, condition string, temperatureC float64) (*models.Weather, error) {
	result, err := r.db.Exec("INSERT INTO weathers (city_id, temperature_c, `condition`) VALUES (?, ?, ?)", cityID, temperatureC, condition)
	if err != nil {
		slog.Error("Failed to create weather", "err", err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		slog.Error("Failed to get last insert ID", "err", err)
		return nil, err
	}
	return &models.Weather{
		City_ID:      int(id),
		TemperatureC: temperatureC,
		Condition:    condition,
	}, nil
}

func (r *WeatherRepository) Update(cityID int, newCityID int, condition string, temperatureC float64) (*models.Weather, error) {
	result, err := r.db.Exec("UPDATE weathers SET city_id = ?, temperature_c = ?, `condition` = ? WHERE city_id = ?", newCityID, temperatureC, condition, cityID)
	if err != nil {
		slog.Error("Failed to update weather", "err", err)
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Failed to get rows affected", "err", err)
		return nil, err
	}
	if rowsAffected == 0 {
		slog.Error("City not found", "city_id", cityID)
		return nil, ErrCityNotFound
	}
	return &models.Weather{
		City_ID:      newCityID,
		TemperatureC: temperatureC,
		Condition:    condition,
	}, nil
}

func (r *WeatherRepository) Delete(cityID int) error {
	result, err := r.db.Exec("DELETE FROM weathers WHERE city_id = ?", cityID)
	if err != nil {
		slog.Error("Failed to delete weather", "err", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Failed to get rows affected", "err", err)
		return err
	}
	if rowsAffected == 0 {
		slog.Error("City not found", "city_id", cityID)
		return ErrCityNotFound
	}
	return nil
}

func (r *WeatherRepository) SaveWeatherData(cityID int, condition string, temperature float64) error {
	_, err := r.FindByCity(cityID)
	if err != nil {
		if errors.Is(err, ErrCityNotFound) {
			_, err = r.Create(cityID, condition, temperature)
			return err
		}
		return err
	}

	_, err = r.Update(cityID, cityID, condition, temperature)

	return err
}

func (r *WeatherRepository) SaveWeatherHistory(cityID int, temperature float64, humidity float64, windSpeed float64) (*models.WeatherHistory, error) {
	_, err := r.FindByCity(cityID)
	if err != nil {
		slog.Error("City not found", "city_id", cityID, "err", err)
		return nil, err
	}

	_, err = r.db.Exec("INSERT INTO weather_histories (city_id, temperature, humidity, wind_speed) VALUES (?, ?, ?, ?)", cityID, temperature, humidity, windSpeed)
	if err != nil {
		slog.Error("Failed to save weather history", "err", err)
		return nil, err
	}

	return &models.WeatherHistory{
		CityID:      cityID,
		Temperature: temperature,
		Humidity:    humidity,
		WindSpeed:   windSpeed,
		RecordedAt:  time.Now(),
	}, nil
}
