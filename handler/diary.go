package handler

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golpo/config"
	"github.com/golpo/dto"
)

func DiaryList(c *fiber.Ctx) {
	rows, err := config.DB.Raw("SELECT id, title, author_id, content, created_at FROM diaries order by created_at").Rows()
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	defer rows.Close()
	result := dto.Diaries{}

	for rows.Next() {
		diary := dto.Diary{}
		err := rows.Scan(&diary.ID, &diary.Title, &diary.AuthorID, &diary.Content, &diary.CreatedAt)
		if err != nil {
			c.Status(500).Send(err)
			return
		}
		result.Diaries = append(result.Diaries, diary)
	}
	if err := c.JSON(result); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func CreateDiary(c *fiber.Ctx) {
	d := new(dto.Diary)
	if err := c.BodyParser(d); err != nil {
		c.Status(400).Send(err)
		return
	}
	d.AuthorID = fmt.Sprintf("%v", c.Locals("user"))
	res := config.DB.Create(&d)
	if res.Error != nil {
		c.Status(500).Send("Creation failed")
		return
	}

	if err := c.JSON(dto.NewSuccessResponse()); err != nil {
		c.Status(500).Send(err)
		return
	}
}