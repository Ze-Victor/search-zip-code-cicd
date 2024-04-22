package pkg

import (
	"net/http"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Search CEP
// @Description Search Address by CEP
// @Tags CEP
// @Produce json
// @Param Authorization header string true "Token"
// @Param cep path string true "CEP"
// @Success 200 {object} SendSuccessCEPResponse
// @Failure 400 {object} SendErrorCEPResponse
// @Failure 401 {object} SendErrorAuthResponse
// @Failure 404 {object} SendErrorCEPResponse
// @Failure 500 {object} SendErrorCEPResponse
// @Router /cep/{cep} [get]
func SearchCEPHandler(ctx *gin.Context) {
	logger := config.GetLogger("cep_search_handler")

	cep := ctx.Param("cep")

	if err := services.ValidateCEP(cep); err != nil {
		logger.Errorf("Invalid CEP")
		sendErrorCEPResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	address, err := SearchAddressByCEP(ctx, cep)
	if err != nil {
		logger.Errorf("Error searching for CEP: %v", err)
		sendErrorCEPResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucessCEPResponse(ctx, address)
}
