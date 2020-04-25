package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/dto"
	"github.com/golpo/repository"
	"github.com/golpo/service"
	"github.com/golpo/util"
)


func UserList(c *fiber.Ctx) {
	userRepo := repository.UserRepoGorm{DB: config.DB}
	userService := service.UserServiceImpl{UserRepo: userRepo}
	res, err := userService.GetUsers()
	if err != nil {
		util.LogWithTrack(c, err.Message)
		c.Status(500).JSON(dto.ServerError(c))
		return
	}

	if err := c.JSON(res); err != nil {
		util.LogWithTrack(c, err.Error())
		c.Status(500).JSON(dto.ServerError(c))
		return
	}
}

func CreateUser(c *fiber.Ctx) {
	u := new(dto.User)
	if err := c.BodyParser(u); err != nil {
		c.Status(400).Send(err)
		return
	}
	userRepo := repository.UserRepoGorm{DB: config.DB}
	userService := service.UserServiceImpl{UserRepo: userRepo}
	res, err := userService.CreateUser(u)
	if err != nil {
		util.LogWithTrack(c, err.Message)
		c.Status(500).JSON(dto.ServerError(c))
		return
	}

	if err := c.Status(201).JSON(res); err != nil {
		util.LogWithTrack(c, err.Error())
		c.Status(500).JSON(dto.ServerError(c))
		return
	}
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
