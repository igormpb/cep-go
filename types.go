package cepgo

// ProviderType -
type ProviderType string

var (
	//ProviderViaCep -
	ProviderViaCep ProviderType = "VIACEP"
	//ProviderBrasilAPI -
	ProviderBrasilAPI ProviderType = "BRASIL_API"
	//ProviderBrasilAberto -
	ProviderBrasilAberto ProviderType = "BRASIL_ABERTO"
	//ProviderOpenCEP -
	ProviderOpenCEP ProviderType = "OPEN_CEP"
)

// Provider -
type Provider struct {
	Type ProviderType
	URL  string
}

// ProviderResponse -
type ProviderResponse struct {
	Data  []byte
	Type  ProviderType
	Error error
}

// CepResponse -
type CepResponse struct {
	Zipcode      string `json:"zipcode"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

// ViaCepResponse -
type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (r ViaCepResponse) convert() CepResponse {
	return CepResponse{
		Zipcode:      r.Cep,
		State:        r.Uf,
		City:         r.Localidade,
		Neighborhood: r.Bairro,
		Street:       r.Logradouro,
	}
}

// BrasilAPIRespose -
type BrasilAPIRespose struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (r BrasilAPIRespose) convert() CepResponse {
	return CepResponse{
		Zipcode:      r.Cep,
		State:        r.State,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		Street:       r.Street,
	}
}

// BrasilAbertoRespose -
type BrasilAbertoRespose struct {
	Result struct {
		Street     string `json:"street"`
		Complement string `json:"complement"`
		District   string `json:"district"`
		DistrictID int    `json:"districtId"`
		City       string `json:"city"`
		CityID     int    `json:"cityId"`
		IbgeID     int    `json:"ibgeId"`
		State      string `json:"state"`
		StateShort string `json:"stateShortname"`
		Zipcode    string `json:"zipcode"`
	} `json:"result"`
}

func (r BrasilAbertoRespose) convert() CepResponse {
	return CepResponse{
		Zipcode:      r.Result.Zipcode,
		State:        r.Result.State,
		City:         r.Result.City,
		Neighborhood: r.Result.District,
		Street:       r.Result.Street,
	}
}

// OpenCEPRespose -
type OpenCEPRespose struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
}

func (r OpenCEPRespose) convert() CepResponse {
	return CepResponse{
		Zipcode:      r.Cep,
		State:        r.UF,
		City:         r.Localidade,
		Neighborhood: r.Bairro,
		Street:       r.Logradouro,
	}
}
