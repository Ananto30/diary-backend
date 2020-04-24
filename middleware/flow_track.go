package middleware

import (
	"fmt"
	"github.com/gofiber/fiber"
	//"github.com/satori/go.uuid"
	//"log"
)

func Tracker() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		//trackID, err := uuid.NewV4()
		//if err != nil {
		//	log.Printf("Something went wrong: %s\n", err)
		//	return
		//}
		c.Locals("track-id", fmt.Sprintf("%v", c.Fasthttp.ID()))
		c.Next()
	}
}
