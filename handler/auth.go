package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (h AuthHandler) Login(c *fiber.Ctx) {
	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		c.Status(400).Send(err)
		return
	}
	res, err := h.AuthService.Login(req)
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(res)
}
