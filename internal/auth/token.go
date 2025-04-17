package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string, secret string) (string, error) {
	claims := &JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
