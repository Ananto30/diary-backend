package util

import (
	"fmt"
	"github.com/gofiber/fiber"
	"log"
)

func LogWithTrack(c *fiber.Ctx, s string) {
	log.Println(fmt.Sprintf("%v", c.Locals("track-id")) + " - " + s)
}
