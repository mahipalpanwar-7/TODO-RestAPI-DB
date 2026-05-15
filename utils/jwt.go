package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key")

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
