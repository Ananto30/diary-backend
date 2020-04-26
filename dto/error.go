package dto

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/golpo/internalError"
)

type ErrorResponse struct {
	RequestID string `json:"request_id"`
	ErrorCode uint32 `json:"error_code"`
	Message   string `json:"message"`
}

func ServerError(c *fiber.Ctx) *ErrorResponse {
	return &ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: internalError.ServerError,
		Message:   "Server error",
	}
}

func InvalidAccessToken(c *fiber.Ctx) *ErrorResponse {
	return &ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: internalError.InvalidAccessToken,
		Message:   "Invalid access token",
	}
}

func ForbiddenError(c *fiber.Ctx) *ErrorResponse {
	return &ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: internalError.Forbidden,
		Message:   "Forbidden",
	}
}

