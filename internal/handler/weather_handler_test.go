package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalcolmFuchs/Go-FinExity/internal/api"
)

type MockWeatherService struct{}

func (m *MockWeatherService) GetWeather(lat, lon string) (*api.APIResponse, error) {
	return &api.APIResponse{
		Latitude:  52.52,
		Longitude: 13.405,
		CurrentWeather: api.CurrentWeather{
			Temperature: 22.0,
			Windspeed:   4.5,
		},
	}, nil
}

func TestGetWeatherHandler(t *testing.T) {
	mockService := &MockWeatherService{}
	handler := NewWeatherHandler(mockService)

	reqBody := map[string]string{
		"lat": "52.52",
		"lon": "13.405",
	}
	body, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/weather", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Fehler beim Erstellen der Anfrage: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.GetWeather(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler gab Statuscode %v zurück, erwartet wurde %v", status, http.StatusOK)
	}

	var resp api.APIResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("Fehler beim Decodieren der Antwort: %v", err)
	}

	if resp.CurrentWeather.Temperature != 22.0 {
		t.Errorf("Erwartete Temperatur 22.0, erhielt %.2f", resp.CurrentWeather.Temperature)
	}
}
