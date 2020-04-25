package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
	"github.com/golpo/util"
)

type UserHandler struct {
	UserService service.UserService
}

func (h UserHandler) UserList(c *fiber.Ctx) {
	res, err := h.UserService.ListUsers()
	if err != nil {
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.JSON(res)
}

func (h UserHandler) CreateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	err := h.UserService.CreateUser(u)
	if err != nil {
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.Status(201).JSON(dto.StatusResponse{Status: "Created"})
}

func (h UserHandler) UpdateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	err := h.UserService.UpdateUser(u)
	if err != nil {
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.Status(202).JSON(dto.StatusResponse{Status: "Updated"})
}

func (h UserHandler) DeleteUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	err := h.UserService.DeleteUser(u.ID)
	if err != nil {
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.Status(202).JSON(dto.StatusResponse{Status: "Deleted"})
}
