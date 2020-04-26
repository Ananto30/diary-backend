package handler

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/service"
	"github.com/golpo/util"
)

type DiaryHandler struct {
	DiaryService service.DiaryService
}

func (h DiaryHandler) DiaryList(c *fiber.Ctx) {
	res, err := h.DiaryService.ListDiaries()
	if err != nil {
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.JSON(res)
}

func (h DiaryHandler) CreateDiary(c *fiber.Ctx) {
	d := new(dto.Diary)
	if err := c.BodyParser(d); err != nil {
		c.Status(400).Send(err)
		return
	}
	// get user id from token, which is in context
	d.AuthorID = fmt.Sprintf("%v", c.Locals("user"))
	err := h.DiaryService.CreateDiary(d)
	if err != nil {
		util.LogWithTrack(c, err.Message)
		mapError(c, err)
		return
	}
	c.Status(201).JSON(dto.StatusResponse{Status: "Created"})
}