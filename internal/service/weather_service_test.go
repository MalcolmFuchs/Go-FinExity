package service

import (
	"testing"

	"github.com/MalcolmFuchs/Go-FinExity/internal/api"
)

type MockWeatherClient struct{}

func (m *MockWeatherClient) GetWeatherData(lat, lon string) (*api.APIResponse, error) {
	return &api.APIResponse{
		Latitude:  52.52,
		Longitude: 13.405,
		CurrentWeather: api.CurrentWeather{
			Temperature: 22.0,
			Windspeed:   4.5,
		},
	}, nil
}

func TestGetWeather(t *testing.T) {
	mockClient := &MockWeatherClient{}
	weatherService := NewWeatherService(mockClient)

	data, err := weatherService.GetWeather("52.52", "13.405")
	if err != nil {
		t.Fatalf("Unerwarteter Fehler: %v", err)
	}

	if data.CurrentWeather.Temperature != 22.0 {
		t.Errorf("Erwartete Temperatur 22.0, erhielt %.2f", data.CurrentWeather.Temperature)
	}

	if data.CurrentWeather.Windspeed != 4.5 {
		t.Errorf("Erwartete Windgeschwindigkeit 4.5, erhielt %.2f", data.CurrentWeather.Windspeed)
	}
}
