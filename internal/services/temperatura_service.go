package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TemperaturaService struct{}

func NewTemperaturaService() *TemperaturaService {
	return &TemperaturaService{}
}

type TemperaturaResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (ws *TemperaturaService) BuscarTemperatura(NomeCidade string) (float64, error) {
	apiKey := "8887ae192b2343f9a32114928240104"
	returnNomeCidade := url.QueryEscape(NomeCidade)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, returnNomeCidade)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var TemperaturaResponse TemperaturaResponse
	err = json.Unmarshal(body, &TemperaturaResponse)
	if err != nil {
		return 0, err
	}

	return TemperaturaResponse.Current.TempC, nil
}
