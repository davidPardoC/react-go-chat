package utils

import (
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
