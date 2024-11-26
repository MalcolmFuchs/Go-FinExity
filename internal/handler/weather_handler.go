package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MalcolmFuchs/Go-FinExity/internal/service"
)

type RequestData struct {
	Lat string `json="lat"`
	Lon string `json="lon"`
}

type WeatherHandler struct {
	service service.WeatherService
}

func NewWeatherHandler(service service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (wh *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Nur POST-Anfragen sind erlaubt", http.StatusMethodNotAllowed)
		return
	}

	var req RequestData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ungültige JSON-Daten", http.StatusBadRequest)
		log.Printf("Fehler beim Decodieren der Anfrage: %v", err)
		return
	}

	data, err := wh.service.GetWeather(req.Lat, req.Lon)
	if err != nil {
		http.Error(w, "Fehler beim Abrufden der Wetterdaten", http.StatusInternalServerError)
		log.Printf("Fehler beim Abrufen der Wetterdaten %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Fehler beim Senden der Antwort", http.StatusInternalServerError)
		log.Printf("Fehler beim Kodieren der Antwort: %v", err)
	}
}
