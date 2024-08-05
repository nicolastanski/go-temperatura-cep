package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolastanski/go-temperatura-cep/internal/services"
)

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cep := vars["cep"]

	if len(cep) != 8 {
		http.Error(w, "CEP inválido", http.StatusUnprocessableEntity)
		return
	}

	city, err := services.GetCityByCEP(cep)
	if err != nil {
		http.Error(w, "CEP não encontrado", http.StatusNotFound)
		return
	}

	tempC, err := services.GetTemperatureByCity(city)
	if err != nil {
		http.Error(w, "Temperatura não encontrada", http.StatusInternalServerError)
		return
	}

	tempF := services.ConvertCelsiusToFahrenheit(tempC)
	tempK := services.ConvertCelsiusToKelvin(tempC)

	response := WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
