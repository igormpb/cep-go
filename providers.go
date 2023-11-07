package cepgo

import (
	"fmt"
	"net/http"
)

// Providers - serviços
func Providers(cep string) (providers []Provider) {
	ViaCep := Provider{
		Type: ProviderViaCep,
		URL:  fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep),
	}

	BrasilAPI := Provider{
		Type: ProviderBrasilAPI,
		URL:  fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep),
	}

	BrasilAberto := Provider{
		Type: ProviderBrasilAberto,
		URL:  fmt.Sprintf("https://brasilaberto.com/api/v1/zipcode/%s", cep),
	}

	OpenCEP := Provider{
		Type: ProviderOpenCEP,
		URL:  fmt.Sprintf("https://opencep.com/v1/%s", cep),
	}

	providers = append(providers, ViaCep, BrasilAPI, BrasilAberto, OpenCEP)
	return providers

}

// NewRequest -
func (p Provider) NewRequest() (*http.Request, error) {
	request, e := http.NewRequest("GET", p.URL, nil)
	if e != nil {
		return nil, e
	}
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}
