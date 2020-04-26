package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	util2 "github.com/golpo/service/util"
	"github.com/golpo/util"
	"strings"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

func Auth() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		token := c.Get("Authorization")
		tokenSplit := strings.Fields(string(token))
		if len(tokenSplit) == 0 || tokenSplit[0] != "Bearer" {
			c.Status(403).Send("Forbidden")
			return
		}
		claims, done, err := util2.ValidateToken(tokenSplit[1])
		if err != nil {
			util.LogWithTrack(c, err.Error())
			c.Status(403).JSON(dto.InvalidToken(c))
			return
		}
		if !done {
			c.Status(403).JSON(dto.InvalidToken(c))
			return
		}
		c.Locals("user", claims.UserID)
		c.Next()
	}
}
