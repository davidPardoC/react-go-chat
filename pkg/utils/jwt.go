package utils

import (
	"fmt"

	"github.com/davidPardoC/go-chat/pkg/constants"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	jwt.RegisteredClaims
}

func SignToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.TOKEN_SECRET))
}

func IsTokenValid(tokenString string) (isValid bool, claims jwt.MapClaims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected string token")
		}
		return []byte(constants.TOKEN_SECRET), nil
	})

	if err != nil {
		return false, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return true, claims
	}
	return false, nil
}
