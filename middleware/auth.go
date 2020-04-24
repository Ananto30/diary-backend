package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/golpo/util"
	"strings"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func Auth() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		token := c.Get("Authorization")
		tokenSplit := strings.Fields(string(token))
		if len(tokenSplit) == 0 || tokenSplit[0] != "Bearer" {
			c.Status(500).Send("Authorization failed")
			return
		}
		claims, done := util.ValidateToken(c, tokenSplit[1])
		if !done {
			c.Status(500).Send("Invalid token")
			return
		}
		//log.Println(tokenSplit[1])
		c.Locals("user", claims.UserID)
		c.Next()
	}
}
