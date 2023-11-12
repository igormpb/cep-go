package cepgo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Client -
type Client struct {
	http *http.Client
}

// InitClient -
func newClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

// Do - solicitação http
func (c *Client) Do(url string, result chan ProviderResponse, providerType ProviderType) error {
	var response *http.Response
	var e error

	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		result <- ProviderResponse{}
		return errors.New("Erro inesperado")
	}

	request.Header.Set("Content-Type", "application/json")

	maxRetry := 3
	for retry := 0; retry < maxRetry; retry++ {
		response, e = c.http.Do(request)
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
