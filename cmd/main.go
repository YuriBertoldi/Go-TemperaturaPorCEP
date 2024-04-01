package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/YuriBertoldi/Go-TemperaturaPorCEP/internal/conversoes"
	"github.com/YuriBertoldi/Go-TemperaturaPorCEP/internal/services"
)

func main() {
	http.HandleFunc("/clima", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		cep := queryParams.Get("cep")

		if len(cep) != 8 {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		cityName, err := consultarCidade(cep)
		if err != nil {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		temperaturaCelsius, err := consultarTemperatura(cityName)
		if err != nil {
			http.Error(w, "error fetching weather information", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(montarResponse(temperaturaCelsius))
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func consultarCidade(cep string) (string, error) {
	CepService := services.NewCepService()
	cidade, err := CepService.BuscarNomeCidadePeloCep(cep)
	return cidade, err
}

func consultarTemperatura(cidade string) (float64, error) {
	TemperaturaService := services.NewTemperaturaService()
	tempCelsius, err := TemperaturaService.BuscarTemperatura(cidade)
	return tempCelsius, err
}

func montarResponse(temp_C float64) map[string]float64 {

	temperaturaFahrenheit := conversoes.CelsiusToFahrenheit(temp_C)
	temperaturaKelvin := conversoes.CelsiusToKelvin(temp_C)

	response := map[string]float64{
		"temp_C": temp_C,
		"temp_F": temperaturaFahrenheit,
		"temp_K": temperaturaKelvin,
	}

	return response
}
