package http

import (
	"github.com/gofiber/fiber/v2"
)

func SayHi(c *fiber.Ctx) error {
	return c.SendString("Hi!")
}
