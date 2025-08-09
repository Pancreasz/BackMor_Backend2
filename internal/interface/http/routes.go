package http

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	app *gin.Engine,
	userService UserService,
) {
	api := app.Group("/v1/api")
	SetUpUserRoutes(api, userService)
}
