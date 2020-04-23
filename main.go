package main

import (
	"github.com/golpo/handler"
	"log"

	"github.com/gofiber/fiber"
	"github.com/golpo/db"
	_ "github.com/lib/pq"
)

func main() {
	// Connect with database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	userGroup := app.Group("/api/user")
	userGroup.Get("/", handler.UserList)
	userGroup.Post("/", handler.CreateUser)
	userGroup.Put("/", handler.UpdateUser)
	userGroup.Delete("/", handler.DeleteUser)

	app.Listen(3000)
}
