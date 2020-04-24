package handler

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
)

func DiaryList(c *fiber.Ctx) {
	service.GetDiaries(c)
}

func CreateDiary(c *fiber.Ctx) {
	d := new(dto.Diary)
	if err := c.BodyParser(d); err != nil {
		c.Status(400).Send(err)
		return
	}
	service.CreateDiary(c, d)
}