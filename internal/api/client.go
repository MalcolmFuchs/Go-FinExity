package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIResponse struct {
	Latitude            float64             `json:"latitude"`
	Longitude           float64             `json:"longitude"`
	Timezone            string              `json:"timezone"`
	CurrentWeatherUnits CurrentWeatherUnits `json:"current_weather_units"`
	CurrentWeather      CurrentWeather      `json:"current_weather"`
}

type CurrentWeatherUnits struct {
	Time        string `json:"time"`
	Temperature string `json:"temperature"`
	Windspeed   string `json:"windspeed"`
}

type CurrentWeather struct {
	Time        string  `json:"time"`
	Temperature float64 `json:"temperature"`
	Windspeed   float64 `json:"windspeed"`
}

type WeatherClient interface {
	GetWeatherData(lat, lon string) (*APIResponse, error)
}

type weatherClient struct {
	client *http.Client
}

func NewWeatherClient() WeatherClient {
	return &weatherClient{
		client: &http.Client{},
	}
}

func (wc *weatherClient) GetWeatherData(lat, lon string) (*APIResponse, error) {

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lat, lon)
	resp, err := wc.client.Get(url)
	if err != nil {
		log.Printf("Fehler beim Aubrufden der Wetterdaten: %v", err)
		return nil, fmt.Errorf("Fehler beim Abrufen der Wetterdaten: %w", err)

	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		log.Printf("Fehler beim Decodieren der Antwort: %v", err)
		return nil, fmt.Errorf("Fehler beim Decodieren der Antwort %w", err)
	}

	return &apiResp, nil
}
