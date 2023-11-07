package cepgo

import (
	"encoding/json"
	"errors"
)

// CEP -
func CEP(cep string) (CepResponse, error) {
	var cepResponse CepResponse
	providerResponse := InitClient(cep).RequestHTTPProviders()
	if len(providerResponse.Data) == 0 {
		return cepResponse, errors.New("CEP não foi encontrado")
	}
	cepResponse, e := convertToCepResponse(providerResponse.Data, providerResponse.Type)
	if e != nil {
		return cepResponse, errors.New("Ocorreu um error na conversão do CEP")
	}
	return cepResponse, nil
}

func convertToCepResponse(data []byte, providerType ProviderType) (CepResponse, error) {

	if providerType == ProviderBrasilAPI {
		var response BrasilAPIRespose
		if e := json.Unmarshal(data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}

	if providerType == ProviderViaCep {
		var response ViaCepResponse
		if e := json.Unmarshal(data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}

	if providerType == ProviderBrasilAberto {
		var response BrasilAbertoRespose
		if e := json.Unmarshal(data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}

	if providerType == ProviderOpenCEP {
		var response OpenCEPRespose
		if e := json.Unmarshal(data, &response); e != nil {
			return CepResponse{}, e
		}

		return response.convert(), nil
	}
	return CepResponse{}, errors.New("CEP NÃO ENCONTRADO")
}
