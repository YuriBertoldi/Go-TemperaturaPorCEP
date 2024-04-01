package services_test

import (
	"net/http"
	"sync"
	"testing"

	"time"

	"github.com/YuriBertoldi/Go-TemperaturaPorCEP/internal/services"
)

func TestRequestCep(t *testing.T) {
	var wg sync.WaitGroup
	cep := "29709090"
	totalRequests := 20
	requestsPerSecond := 1
	wg.Add(totalRequests)

	makeRequest := func() {
		defer wg.Done()

		CepService := services.NewCepService()

		cityName, err := CepService.BuscarNomeCidadePeloCep(cep)
		if err != nil {
			t.Error("can not find zipcode", http.StatusNotFound)
		}

		if cityName == "" {
			t.Error("Error in load city", cityName)
		}

	}

	for i := 0; i < totalRequests; i++ {
		go makeRequest()
		time.Sleep(time.Second / time.Duration(requestsPerSecond))
	}

	wg.Wait()
}
