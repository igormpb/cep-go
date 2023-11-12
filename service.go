package cepgo

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Service struct {
	cep string
}

// InitClient -
func newService(cep string) *Service {
	return &Service{
		cep: cep,
	}
}

// Providers - provedores disponiveis
func (s *Service) Providers() []Provider {
	ViaCep := Provider{
		Type: ProviderViaCep,
		URL:  fmt.Sprintf("https://viacep.com.br/ws/%s/json", s.cep),
	}

	BrasilAPI := Provider{
		Type: ProviderBrasilAPI,
		URL:  fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", s.cep),
	}

	BrasilAberto := Provider{
		Type: ProviderBrasilAberto,
		URL:  fmt.Sprintf("https://brasilaberto.com/api/v1/zipcode/%s", s.cep),
	}

	OpenCEP := Provider{
		Type: ProviderOpenCEP,
		URL:  fmt.Sprintf("https://opencep.com/v1/%s", s.cep),
	}

	return []Provider{ViaCep, BrasilAPI, BrasilAberto, OpenCEP}

}

// GetCep - busca informação do cep de forma concorrente
func (s *Service) GetCep() ProviderResponse {
	providerResponse := make(chan ProviderResponse, 0)
	client := newClient()
	providers := s.Providers()
	for _, provider := range providers {
		go client.Do(provider.URL, providerResponse, provider.Type)
	}

	errorCount := 0

	for range providers {
		response := <-providerResponse
		if len(response.Data) > 0 {
			return response
		}
		errorCount++
	}

	if errorCount == len(providers) {
		return ProviderResponse{
			Data:  nil,
			Type:  "",
			Error: errors.New("Todas as solicitações falharam"),
		}
	}

	return ProviderResponse{
		Data:  nil,
		Type:  "",
		Error: errors.New("Erro inesperado"),
	}

}

// ConvertToCepResponse - converte para a estrutura padrão CepResponse
func (s *Service) ConvertToCepResponse(p ProviderResponse) (CepResponse, error) {
	if p.Type == ProviderBrasilAPI {
		var response BrasilAPIRespose
		if e := json.Unmarshal(p.Data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}

	if p.Type == ProviderViaCep {
		var response ViaCepResponse
		if e := json.Unmarshal(p.Data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}

	if p.Type == ProviderBrasilAberto {
		var response BrasilAbertoRespose
		if e := json.Unmarshal(p.Data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}

	if p.Type == ProviderOpenCEP {
		var response OpenCEPRespose
		if e := json.Unmarshal(p.Data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}
	return CepResponse{}, errors.New("CEP NÃO ENCONTRADO")
}
