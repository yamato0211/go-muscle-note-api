package services

import (
	"fiber-muscles/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwt(userID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := t.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJwt(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
}
