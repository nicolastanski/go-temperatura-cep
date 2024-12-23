package services

import (
	"testing"
)

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "Zero graus Celsius",
			input:    0,
			expected: 32,
		},
		{
			name:     "Temperatura ambiente (~25ºC)",
			input:    25,
			expected: 77,
		},
		{
			name:     "Temperatura negativa (-40ºC)",
			input:    -40,
			expected: -40,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvertCelsiusToFahrenheit(tc.input)
			if got != tc.expected {
				t.Errorf("ConvertCelsiusToFahrenheit(%.2f) = %.2f; esperava %.2f",
					tc.input, got, tc.expected)
			}
		})
	}
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "Zero graus Celsius",
			input:    0,
			expected: 273,
		},
		{
			name:     "Temperatura ambiente (~25ºC)",
			input:    25,
			expected: 298,
		},
		{
			name:     "Temperatura negativa (-40ºC)",
			input:    -40,
			expected: 233,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConvertCelsiusToKelvin(tc.input)
			if got != tc.expected {
				t.Errorf("ConvertCelsiusToKelvin(%.2f) = %.2f; expected %.2f",
					tc.input, got, tc.expected)
			}
		})
	}
}
