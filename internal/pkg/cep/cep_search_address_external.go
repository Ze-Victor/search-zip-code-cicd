package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/schemas"
)

func CEPSearchAddressExternal(cep string) schemas.Address {

	logger := config.GetLogger("External_Search")

	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return schemas.Address{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}

	defer resp.Body.Close()

	var address schemas.AddressResponse

	err = json.Unmarshal(body, &address)
	if err != nil {
		logger.Error(err)
	}

	cepFormated := strings.ReplaceAll(address.CEP, "-", "")

	addressResponse := schemas.Address{
		CEP:          cepFormated,
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
	}

	return addressResponse
}
