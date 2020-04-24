package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func ValidateToken(c *fiber.Ctx, token string) (*Claims, bool) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		LogWithTrack(c, err.Error())
		return nil, false
	}
	if !tkn.Valid {
		LogWithTrack(c, "Token not valid")
		return nil, false
	}
	return claims, true
}
