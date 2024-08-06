package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByCity(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	escapedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, escapedCity)

	fmt.Print(url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var weatherAPIResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherAPIResponse); err != nil {
		return 0, err
	}

	return weatherAPIResponse.Current.TempC, nil
}
