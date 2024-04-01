package services_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"time"

	"github.com/YuriBertoldi/Go-TemperaturaPorCEP/internal/services"
)

func TestRequestTemperatura(t *testing.T) {
	var wg sync.WaitGroup
	cityName := "Colatina"
	totalRequests := 5
	requestsPerSecond := 1
	wg.Add(totalRequests)

	makeRequest := func() {
		defer wg.Done()

		TemperaturaService := services.NewTemperaturaService()

		tempCelsius, err := TemperaturaService.BuscarTemperatura(cityName)
		if err != nil {
			t.Error("error fetching weather information", tempCelsius)
			return
		}
	}

	for i := 0; i < totalRequests; i++ {
		go makeRequest()
		time.Sleep(time.Second / time.Duration(requestsPerSecond))
	}

	wg.Wait()
}

func TestAPILocal(t *testing.T) {
	var wg sync.WaitGroup
	cep := "29709090"
	totalRequests := 20
	requestsPerSecond := 1
	wg.Add(totalRequests)

	makeRequest := func() {
		defer wg.Done()

		url := fmt.Sprintf("http://localhost:8080/clima?cep=%s", cep)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Erro ao criar a requisição: %v\n", err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Erro na requisição: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			message := string(body)
			fmt.Printf("Result: %s\n", message)
		}
	}

	for i := 0; i < totalRequests; i++ {
		go makeRequest()
		time.Sleep(time.Second / time.Duration(requestsPerSecond))
	}

	wg.Wait()
}
