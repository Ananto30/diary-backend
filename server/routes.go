package server

import (
	"github.com/gofiber/fiber"
	"github.com/golpo/handler"
	"github.com/golpo/middleware"
)

func InitiateRoutes(app *fiber.App) {
	userGroup := app.Group("/api/user")
	userGroup.Use(middleware.Auth())
	userGroup.Get("/", handler.UserList)
	userGroup.Post("/", handler.CreateUser)
	userGroup.Put("/", handler.UpdateUser)
	userGroup.Delete("/", handler.DeleteUser)

	authGroup := app.Group("/api/auth")
	authGroup.Post("/login", handler.Login)

	diaryGroup := app.Group("/api/diary")
	diaryGroup.Use(middleware.Auth())
	diaryGroup.Get("/", handler.DiaryList)
	diaryGroup.Post("/", handler.CreateDiary)
}
