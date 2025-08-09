package http

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(
	app *fiber.App,
	userService UserService,
) {
	api := app.Group("/v1/api") // นี่อ่อ
	SetUpUserRoutes(api, userService)
}

//  อ่าา ห้ามเกิน 2 version ได้ๆ ก็แค่่นี้ใช่ป่ะ
