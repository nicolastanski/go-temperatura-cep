package services

func ConvertCelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

func ConvertCelsiusToKelvin(c float64) float64 {
	return c + 273
}
