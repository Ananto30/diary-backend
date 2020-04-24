package config

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/golpo/handler"
	"github.com/golpo/middleware"
	"github.com/golpo/util"
	"log"
)

func StartServer() {

	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// recover from panic errors
	rCfg := recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			util.LogWithTrack(c, err.Error())
			c.SendString("Server Error")
			c.SendStatus(500)
		},
	}
	app.Use(recover.New(rCfg))

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
