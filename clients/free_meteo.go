package clients

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"weather-api/repositories"
)

const baseURL = "https://api.open-meteo.com/v1/forecast"

type WeatherClient struct {
	httpClient *http.Client
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{httpClient: &http.Client{}}
}

func (c *WeatherClient) GetWeatherData(city repositories.City) ([]byte, error) {
	url := fmt.Sprintf("%s?latitude=%f&longitude=%f&current_weather=true", baseURL, city.Latitude, city.Longitude)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		log.Println("Error fetching weather data:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	return body, nil
}
