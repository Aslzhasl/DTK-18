package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenStr string) error {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token: %w", err)
	}

	return nil
}
