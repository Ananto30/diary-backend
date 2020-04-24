package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
)

func Login(c *fiber.Ctx) {
	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		c.Status(400).Send(err)
		return
	}
	service.Login(c, req)
}
