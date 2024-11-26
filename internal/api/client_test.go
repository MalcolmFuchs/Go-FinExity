package api

import (
	"testing"
)

type MockWeatherClient struct{}

func (m *MockWeatherClient) GetWeatherData(lat, lon string) (*APIResponse, error) {
	return &APIResponse{
		Latitude:  52.52,
		Longitude: 13.405,
		CurrentWeather: CurrentWeather{
			Temperature: 20.0,
			Windspeed:   5.0,
		},
	}, nil
}

func TestGetWeatherData(t *testing.T) {
	mockClient := &MockWeatherClient{}

	data, err := mockClient.GetWeatherData("52.52", "13.405")
	if err != nil {
		t.Fatalf("Unerwarteter Fehler: %v", err)
	}

	if data.CurrentWeather.Temperature != 20.0 {
		t.Errorf("Erwartete Temperatur 20.0, erhielt %.2f", data.CurrentWeather.Temperature)
	}

	if data.CurrentWeather.Windspeed != 5.0 {
		t.Errorf("Erwartete Windgeschwindigkeit 5.0, erhielt %.2f", data.CurrentWeather.Windspeed)
	}
}
