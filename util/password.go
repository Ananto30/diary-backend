package util

import (
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(c *fiber.Ctx, p string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		LogWithTrack(c, err.Error())
	}
	pStr := string(hashedPassword)
	return pStr
}
