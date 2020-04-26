package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) (*Claims, bool, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, false, err
	}
	if !tkn.Valid {
		return nil, false, nil
	}
	return claims, true, nil
}
