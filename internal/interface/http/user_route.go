package http

import (
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(api gin.IRouter, userService UserService) {
	userHandler := NewUserServiceHandler(userService)
	userRoutes := api.Group("/user")
	userRoutes.GET("/", userHandler.GetAllUser)
	userRoutes.GET("/id/:id", userHandler.GetUserByID)
	userRoutes.GET("/by-email", userHandler.GetUserByEmail)
	userRoutes.POST("/", userHandler.InsertNewUser)

}
