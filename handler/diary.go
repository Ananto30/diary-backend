package handler

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/repository"
)

type DiaryHandler struct {
	DiaryRepo repository.DiaryRepo
}

func (h DiaryHandler) DiaryList(c *fiber.Ctx) {
	result, err := h.DiaryRepo.List()
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(result)
}

func (h DiaryHandler) CreateDiary(c *fiber.Ctx) {
	d := new(dto.Diary)
	if err := c.BodyParser(d); err != nil {
		c.Status(400).Send(err)
		return
	}
	d.AuthorID = fmt.Sprintf("%v", c.Locals("user"))
	if err := h.DiaryRepo.Create(d); err != nil {
		errorHandler(c, err)
		return
	}

	c.JSON(dto.StatusResponse{Status: "Created"})
}

func (h DiaryHandler) GetDiaryById(c *fiber.Ctx) {
	result, err := h.DiaryRepo.GetByID(c.Params("id"))
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(result)
}

func (h DiaryHandler) GetDiaryByUserId(c *fiber.Ctx) {
	result, err := h.DiaryRepo.GetByUserID(fmt.Sprintf(c.Params("id")))
	if err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(result)
}
