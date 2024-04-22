package services

import (
	"fmt"
	"strconv"

	"github.com/Ze-Victor/search-zip-code/config"
)

func ValidateCEP(cep string) error {

	logger := config.GetLogger("cep_validate")

	logger.Debug("Validating for CEP...")

	if len(cep) != 8 {
		return fmt.Errorf("CEP must have 8 digits")
	}

	_, err := strconv.Atoi(cep)
	if err != nil {
		return fmt.Errorf("CEP must have only digits")
	}

	return nil
}
