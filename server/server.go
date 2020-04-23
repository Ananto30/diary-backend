package server

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/golpo/db"
	"github.com/golpo/handler"
	"log"
)

func StartServer() {
	// Connect with database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()
	app.Use(logger.New(logger.Config{Format: "${time} - ${ip} - ${method} ${path} - ${body} - ${status} [${ua}]\n"}))

	userGroup := app.Group("/api/user")
	userGroup.Get("/", handler.UserList)
	userGroup.Post("/", handler.CreateUser)
	userGroup.Put("/", handler.UpdateUser)
	userGroup.Delete("/", handler.DeleteUser)

	app.Listen(3000)
}
