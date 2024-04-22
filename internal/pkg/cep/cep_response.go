package pkg

import (
	"net/http"

	"github.com/Ze-Victor/search-zip-code/internal/schemas"
	"github.com/gin-gonic/gin"
)

func sendErrorCEPResponse(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucessCEPResponse(ctx *gin.Context, address schemas.Address) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Address located",
		"data":    address,
	})
}

type SendSuccessCEPResponse struct {
	Message string          `json:"message"`
	Data    schemas.Address `json:"data"`
}

type SendErrorCEPResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
