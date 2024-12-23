package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByCity(city string) (float64, error) {
	apiKey := "cd31c61e814d465fb67204954240508"
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
		return 0, fmt.Errorf("500 Internal Server Error")
	}

	return weatherAPIResponse.Current.TempC, nil
}
