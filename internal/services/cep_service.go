package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CepService struct{}

func NewCepService() *CepService {
	return &CepService{}
}

type cepCodeResponse struct {
	Localidade string `json:"localidade"`
}

func (zc *CepService) BuscarNomeCidadePeloCep(cepCode string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cepCode)
	response, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request to viaCEP: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var cepCodeResponse cepCodeResponse
	err = json.Unmarshal(body, &cepCodeResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if cepCodeResponse.Localidade == "" {
		return "", fmt.Errorf("localidade not found for zipcode: %s", cepCode)
	}

	return cepCodeResponse.Localidade, nil
}
