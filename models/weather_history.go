package models


import "time"

type WeatherHistory struct {
	ID          int       `json:"id"`
	CityID      int       `json:"city_id"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	RecordedAt  time.Time `json:"recorded_at"`
}