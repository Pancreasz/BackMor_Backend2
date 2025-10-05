package http

import (
	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(api gin.IRouter, userService UserService) {
	userHandler := NewUserServiceHandler(userService)

	userRoutes := api.Group("/user")

	userRoutes.GET("/", middleware.AuthRequired(), userHandler.GetAllUser)
	userRoutes.GET("/id/:id", userHandler.GetUserByID)
	userRoutes.GET("/email/:email", userHandler.GetUserByEmail)
	userRoutes.POST("/", userHandler.InsertNewUser)
}
