package main

import (
	"github.com/golpo/service"
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

	// Get all records from postgreSQL
	app.Get("/employee", func(c *fiber.Ctx) {
		// Insert Employee into database
		service.GetEmployees(c)
	})

	// Add record into postgreSQL
	app.Post("/employee", func(c *fiber.Ctx) {
		service.CreateEmployee(c)
	})

	// Update record into postgreSQL
	app.Put("/employee", func(c *fiber.Ctx) {
		service.UpdateEmployee(c)
	})

	// Delete record from postgreSQL
	app.Delete("/employee", func(c *fiber.Ctx) {
		service.DeleteEmployee(c)
	})

	app.Listen(3000)
}




