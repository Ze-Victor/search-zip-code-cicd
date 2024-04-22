package pkg

import (
	"net/http"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/Ze-Victor/search-zip-code/internal/services"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Sumary Authorization API
// @Description Generate token jwt for API authorization
// @Tags Authorization Token
// @Accept json
// @Produce json
// @Param request body CredentialsAuth true "Request body"
// @Success 200 {object} SendSuccessAuthResponse
// @Failure 400 {object} SendErrorAuthResponse
// @Failure 500 {object} SendErrorAuthResponse
// @Router /auth [post]
func CreateTokenHandler(ctx *gin.Context) {

	logger := config.GetLogger("Authorization")

	request := CredentialsAuth{}
	err := ctx.BindJSON(&request)

	if err != nil {
		logger.Errorf("Internal server error")
		sendErrorAuthResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Validate error")
		sendErrorAuthResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := services.CreateToken(request.Email, request.Password)
	if err != nil {
		logger.Errorf("Internal server error")
		sendErrorAuthResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucessAuthResponse(ctx, token)

}
