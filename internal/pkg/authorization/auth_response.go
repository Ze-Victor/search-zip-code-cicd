package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenJWT struct {
	Token string
}

func sendErrorAuthResponse(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSucessAuthResponse(ctx *gin.Context, token string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

type SendSuccessAuthResponse struct {
	Token string `json:"token"`
}

type SendErrorAuthResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
