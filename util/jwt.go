package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"time"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GenerateToken(c *fiber.Ctx, userID string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: "a1b2c3",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		LogWithTrack(c, err.Error())
		return "", err
	}
	return tokenString, nil
}

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
