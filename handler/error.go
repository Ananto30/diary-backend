package handler

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
)

func mapError(c *fiber.Ctx, iError *internalError.IError) {
	statusCode := iError.ErrorCode / 100
	//log.Println(statusCode)
	if statusCode == 500 || statusCode == 503 || statusCode == 422 {
		c.Status(int(statusCode)).JSON(&dto.ErrorResponse{
			RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
			ErrorCode: iError.ErrorCode,
			Message:   "Server error",
		})
		return
	}
	c.Status(int(statusCode)).JSON(&dto.ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: iError.ErrorCode,
		Message:   iError.Message,
	})
	return
}
