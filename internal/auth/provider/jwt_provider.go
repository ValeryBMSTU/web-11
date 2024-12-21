package provider

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secretKey string
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTProvider(secretKey string) *JWTProvider {
	return &JWTProvider{secretKey: secretKey}
}
