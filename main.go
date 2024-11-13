package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestData struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

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

func getWeather(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Nur POST-Anfragen sind erlaubt", http.StatusMethodNotAllowed)
		return
	}

	var req RequestData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ungültige JSON-Daten", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", req.Lat, req.Lon)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Fehler beim Abrufen der Wetterdaten", http.StatusInternalServerError)
		log.Println("Fehler beim Abrufen der URL:", err)
		return
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		http.Error(w, "Fehler beim Parsen der API-Antwort", http.StatusInternalServerError)
		log.Println("Fehler beim Parsen der API-Antwort:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(apiResp); err != nil {
		http.Error(w, "Fehler beim Senden der Antwort", http.StatusInternalServerError)
		log.Printf("Fehler beim Kodieren der Antwort: %v", err)
	}
}

func main() {
	port := ":8080"
	http.HandleFunc("/weather", getWeather)

	log.Printf("Server running on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
