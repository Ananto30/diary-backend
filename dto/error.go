package dto

import (
	"fmt"
	"github.com/gofiber/fiber"
)

type ErrorResponse struct {
	RequestID string `json:"request_id"`
	ErrorCode uint32 `json:"error_code"`
	Message   string `json:"message"`
}

func ServerError(c *fiber.Ctx) *ErrorResponse {
	return &ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: 50001,
		Message:   "Server error",
	}
}

func InvalidToken(c *fiber.Ctx) *ErrorResponse {
	return &ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: 40301,
		Message:   "Invalid access token",
	}
}

func InvalidCredentials(c *fiber.Ctx) *ErrorResponse {
	return &ErrorResponse{
		RequestID: fmt.Sprintf("%v", c.Fasthttp.ID()),
		ErrorCode: 40302,
		Message:   "Invalid credentials",
	}
}