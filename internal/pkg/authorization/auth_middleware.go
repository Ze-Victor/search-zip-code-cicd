package pkg

import (
	"net/http"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := config.GetLogger("Authorization_Middleware")
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			sendErrorAuthResponse(ctx, http.StatusUnauthorized, "Authorization token is required")
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret_key), nil
		})

		if err != nil {
			if !token.Valid {
				logger.Errorf("Invalid or expired token")
				sendErrorAuthResponse(ctx, http.StatusUnauthorized, err.Error())
				ctx.Abort()
			} else {
				logger.Errorf("Internal server error")
				sendErrorAuthResponse(ctx, http.StatusInternalServerError, err.Error())

			}
			return
		}

		ctx.Next()
	}
}
