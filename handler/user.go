package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/repository"
	"github.com/golpo/util"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (h UserHandler) UserList(c *fiber.Ctx) {
	result, err := h.UserRepo.List()
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(result)
}

func (h UserHandler) CreateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	pStr := util.HashPassword(c, *u.Password)
	u.Password = &pStr
	if err := h.UserRepo.Create(u); err != nil {
		errorHandler(c, err)
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
	err := h.UserRepo.Update(u)
	if err != nil {
		errorHandler(c, err)
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
	err := h.UserRepo.Delete(u.ID)
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.Status(202).JSON(dto.StatusResponse{Status: "Deleted"})
}
