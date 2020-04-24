package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
)

func UserList(c *fiber.Ctx) {
	service.GetUsers(c)
}

func CreateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	service.CreateUser(c, u)
}

func UpdateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	service.UpdateUser(c, u)
}

func DeleteUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	service.DeleteUser(c, u)
}
