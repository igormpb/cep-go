package cepgo

import (
	"errors"
)

// CEP - Função retorna informações sobre o cep:
//
// Zipcode - cep.
//
// State - estado.
//
// City - cidade.
//
// Neighborhood - bairro.
//
// Street - rua.
//
// Service - serviço utilizado.
func CEP(cep string) (CepResponse, error) {
	var cepResponse CepResponse

	//iniciando o client
	client := newService(cep)

	// buscando informações do cep
	response := client.GetCep()
	if len(response.Data) == 0 {
		return cepResponse, errors.New("CEP não foi encontrado")
	}

	//convertendo para o response padrão
	cepResponse, e := client.ConvertToCepResponse(response)
	if e != nil {
		return cepResponse, errors.New("Ocorreu um error na conversão do CEP")
	}
	return cepResponse, nil
}
