package services

import (
	"time"

	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.Secret_key)

func CreateToken(email string, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    email,
			"password": password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
