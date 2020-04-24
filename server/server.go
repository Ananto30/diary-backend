package server

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/golpo/db"
	"github.com/golpo/handler"
	"github.com/golpo/middleware"
	"log"
)

func StartServer() {
	// Connect with database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()
	app.Use(logger.New(logger.Config{Format: "${time} - ${ip} - ${method} ${path} - ${status} - ${body} - ${latency} \t[${ua}]\n"}))
	app.Use(middleware.Tracker())

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


	app.Listen(3000)
}
