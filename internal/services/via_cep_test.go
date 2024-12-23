package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockViaCepResponse(cep string) string {
	return fmt.Sprintf(`{
        "cep": "%s",
        "logradouro": "Rua Teste",
        "complemento": "Apartamento 101",
        "bairro": "Bairro Teste",
        "localidade": "Cidade Teste",
        "uf": "TT",
        "ibge": "1234567",
        "gia": "",
        "ddd": "99",
        "siafi": "1234"
    }`, cep)
}

func TestViaCepRepository_GetAddressByCEP_Success(t *testing.T) {
	cep := "01153000"
	expectedCEP := "São Paulo"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockViaCepResponse(expectedCEP)))
	}))
	defer ts.Close()

	originalURL := viaCepURL
	viaCepURL = ts.URL + "/%s/json/"
	defer func() {
		viaCepURL = originalURL
	}()

	city, err := GetCityByCEP(cep)
	if err != nil {
		t.Fatalf("Não esperava erro, mas obteve: %v", err)
	}
	if city != "São Paulo" {
		t.Errorf("Esperava 'São Paulo', obteve '%s'", city)
	}
}
func TestViaCepRepository_GetAddressByCEP_Unknown(t *testing.T) {
	unknownCEP := "99999999"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockViaCepResponse(unknownCEP)))
	}))
	defer ts.Close()

	originalURL := viaCepURL
	viaCepURL = ts.URL + "/%s/json/"
	defer func() {
		viaCepURL = originalURL
	}()

	city, err := GetCityByCEP(unknownCEP)
	if err == nil {
		t.Fatalf("Esperava erro 'CEP não encontrado', mas não ocorreu erro")
	}
	if err.Error() != "CEP não encontrado" {
		t.Errorf("Esperava erro 'CEP não encontrado', obteve '%s'", err.Error())
	}
	if city != "" {
		t.Errorf("Esperava city vazio, obteve '%s'", city)
	}
}

var viaCepURL = "https://viacep.com.br/ws/"
