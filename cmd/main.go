package main

import (
	"log"
	"net/http"

	"github.com/MalcolmFuchs/Go-FinExity/internal/api"
	"github.com/MalcolmFuchs/Go-FinExity/internal/handler"
	"github.com/MalcolmFuchs/Go-FinExity/internal/service"
)

func main() {
	client := api.NewWeatherClient()
	weatherService := service.NewWeatherService(client)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	http.HandleFunc("/weather", weatherHandler.GetWeather)

	port := ":8080"
	log.Printf("Server läuft auf Port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
