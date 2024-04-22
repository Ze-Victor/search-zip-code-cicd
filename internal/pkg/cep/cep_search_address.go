package pkg

import (
	"net/http"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/schemas"
	"github.com/gin-gonic/gin"
)

func SearchAddressByCEP(ctx *gin.Context, cep string) (schemas.Address, error) {
	logger := config.GetLogger("cep_search")
	logger.Debugf("Searching for CEP %s...", cep)

	address := schemas.Address{}
	adr := config.Db.First(&address, cep)

	if adr.RowsAffected == 0 {
		address, err := searchAddress(ctx, cep)
		if err != nil {
			return schemas.Address{}, err
		}
		return address, nil
	}

	return address, nil
}

func searchAddress(ctx *gin.Context, cep string) (schemas.Address, error) {
	logger := config.GetLogger("cep_search_address")
	logger.Debugf("Searching for external CEP %s...", cep)

	address := CEPSearchAddressExternal(cep)
	if address.CEP == "" {
		for i := len(cep) - 1; i >= 0; i-- {
			if cep[i] != '0' {
				cep = cep[:i] + "0" + cep[i+1:]
				return SearchAddressByCEP(ctx, cep)
			}
		}
		logger.Errorf("CEP not found")
		sendErrorCEPResponse(ctx, http.StatusNotFound, "CEP not found")
		return schemas.Address{}, nil
	}

	if err := config.Db.Create(&address).Error; err != nil {
		logger.Errorf("Error searching CEP: %v", err.Error())
		sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
		return schemas.Address{}, err
	}

	return address, nil
}
