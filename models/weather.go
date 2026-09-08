package models

type Weather struct {
	City_ID      int  `json:"city_id"`
	TemperatureC float64 `json:"temperature_c"`
	Condition    string  `json:"condition"`
}
