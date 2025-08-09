package http

import (
	// "github.com/Pancreasz/BackMor_Backend2/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(api fiber.Router, userService UserService) {
	userHandler := NewUserServiceHandler(userService)
	userRoutes := api.Group("/user")
	userRoutes.Get("/", userHandler.GetAllUser)
	userRoutes.Get("/id/:id", userHandler.GetUserByID)
}
