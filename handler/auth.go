package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
	"github.com/golpo/util"
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
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.JSON(res)
}
