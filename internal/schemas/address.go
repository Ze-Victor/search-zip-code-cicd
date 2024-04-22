package schemas

type Address struct {
	CEP          string
	Street       string
	Neighborhood string
	City         string
	State        string
}

type AddressResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}
