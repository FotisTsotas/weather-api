package utils

import (
	"errors"
	"time"
	"weather-api/config"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(getSecret())
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return getSecret(), nil
	})
	if err != nil {
		return nil, errors.New("Unexpected signing method")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}
	return claims, nil
}

func getSecret() []byte {
	return []byte(config.Load().JWTSecret)
}
