package service

import (
	"fmt"
	"log"

	"github.com/MalcolmFuchs/Go-FinExity/internal/api"
)

type WeatherService interface {
	GetWeather(lat, lon string) (*api.APIResponse, error)
}

type weatherService struct {
	client api.WeatherClient
}

func NewWeatherService(client api.WeatherClient) WeatherService {
	return &weatherService{client: client}
}

func (ws weatherService) GetWeather(lat, lon string) (*api.APIResponse, error) {
	data, err := ws.client.GetWeatherData(lat, lon)
	if err != nil {
		log.Printf("Fehler im WeatherService %v", err)
		return nil, fmt.Errorf("Fehler beim Abrufden der Wetterdaten: %w", err)
	}

	return data, nil
}
