package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/service"
)

func UserList(c *fiber.Ctx) {
	service.GetUsers(c)
}

func CreateUser(c *fiber.Ctx) {
	service.CreateUser(c)
}

func UpdateUser(c *fiber.Ctx) {
	service.UpdateUser(c)
}

func DeleteUser(c *fiber.Ctx) {
	service.DeleteUser(c)
}
