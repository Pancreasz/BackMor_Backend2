package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(
	app *fiber.App,
	userService UserService,
) {
	api := app.Group("/v1/api")
	SetUpUserRoutes(api, userService)
}
