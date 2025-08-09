package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpCustomerRoutes(api fiber.Router) {
	customerRoutes := api.Group("/customer")
	customerRoutes.Get("/", SayHi)
}
