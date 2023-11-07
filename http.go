package cepgo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// ClientCep -
type ClientCep struct {
	cep  string
	http *http.Client
}

// InitClient -
func InitClient(cep string) *ClientCep {
	return &ClientCep{
		cep:  cep,
		http: &http.Client{},
	}
}

// RequestHTTPProviders -
func (c *ClientCep) RequestHTTPProviders() ProviderResponse {
	providerResponse := make(chan ProviderResponse, 0)
	for _, provider := range Providers(c.cep) {
		request, e := provider.NewRequest()
		if e != nil {
			continue
		}
		go c.Do(request, providerResponse, provider.Type)
	}

	errorCount := 0

	for range Providers(c.cep) {
		response := <-providerResponse

		if len(response.Data) > 0 {
			return response
		}

		errorCount++
	}

	if errorCount == len(Providers(c.cep)) {
		return ProviderResponse{
			Data:  nil,
			Type:  "",
			Error: errors.New("Todas as solicitações falharam"),
		}
	}

	// Se ocorreu um erro inesperado
	return ProviderResponse{
		Data:  nil,
		Type:  "", // Defina o tipo apropriado aqui
		Error: errors.New("Erro inesperado"),
	}

}

// Do -
func (c *ClientCep) Do(req *http.Request, result chan ProviderResponse, providerType ProviderType) error {
	var response *http.Response
	var e error
	maxRetry := 3
	for retry := 0; retry < maxRetry; retry++ {
		response, e = c.http.Do(req)
		if e == nil {
			break
		}
	}

	if e != nil {
		result <- ProviderResponse{}
		return errors.New("Erro inesperado")
	}

	body, e := ioutil.ReadAll(response.Body)
	if e != nil {
		result <- ProviderResponse{}
		return e
	}

	if response.StatusCode == 404 {
		result <- ProviderResponse{}
		return errors.New("O cep não foi encontrado")

	}
	if response.StatusCode >= 300 {
		result <- ProviderResponse{}
		return errors.New("Erro inesperado")
	}
	defer response.Body.Close()

	result <- ProviderResponse{
		Data: body,
		Type: providerType,
	}
	return nil
}
