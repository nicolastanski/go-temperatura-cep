package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherProvider_GetTemperatureByCoordinates_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := WeatherAPIResponse{
			Current: struct {
				TempC float64 `json:"temp_c"`
			}{TempC: 19.70},
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	originalFunc := buildWeatherURL
	buildWeatherURL = func(apiKey string, lat, lon float64) string {
		return ts.URL
	}
	defer func() {
		buildWeatherURL = originalFunc
	}()

	temp, err := GetTemperatureByCity("São Paulo")
	if err != nil {
		t.Fatalf("Não esperava erro, mas obteve: %v", err)
	}
	if temp != 19.70 {
		t.Errorf("Temperatura esperada=19.70, obtida=%.2f", temp)
	}
}

func TestWeatherProvider_GetTemperatureByCoordinates_Not200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer ts.Close()

	originalFunc := buildWeatherURL
	buildWeatherURL = func(apiKey string, lat, lon float64) string {
		return ts.URL
	}
	defer func() {
		buildWeatherURL = originalFunc
	}()

	_, err := GetTemperatureByCity("São Paulo")
	if err != nil {
		t.Errorf("500 Internal Server Error '%s'", err.Error())
	}
}

var buildWeatherURL = func(apiKey string, lat, lon float64) string {
	return fmt.Sprintf("https://api.weathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, apiKey)
}
