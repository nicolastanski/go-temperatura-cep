package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
}

func GetCityByCEP(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var viaCEPResponse ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResponse); err != nil {
		return "", err
	}

	if viaCEPResponse.Localidade == "" {
		return "", fmt.Errorf("CEP não encontrado")
	}

	return viaCEPResponse.Localidade, nil
}
