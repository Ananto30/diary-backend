package handler

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golpo/dto"
	"github.com/golpo/internalError"
	"github.com/golpo/util"
)

func mapError(c *fiber.Ctx, iError *internalError.IError) {
	statusCode := iError.ErrorCode / 100

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

func errorHandler(c *fiber.Ctx, err error) {
	switch err.(type) {
	case *internalError.IError:
		ierr := err.(*internalError.IError)
		util.LogWithTrack(c, ierr.Message)
		mapError(c, ierr)
		return
	default:
		util.LogWithTrack(c, err.Error())
		c.Status(500).JSON(dto.ServerError(c))
		return
	}
}
