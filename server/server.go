package server

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/golpo/middleware"
	"github.com/golpo/util"
)

func StartServer() error {
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

	InitiateRoutes(app)

	if err := app.Listen(3000); err != nil {
		return err
	}
	return nil
}

