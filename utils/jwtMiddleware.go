package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UID string `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateJWTToken(uid, username string) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		uid,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}