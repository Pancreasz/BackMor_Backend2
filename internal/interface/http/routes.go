package http

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	app *gin.Engine,
	userService UserService,
	activityService ActivityService,
) {
	api := app.Group("/api/v1")
	SetUpUserRoutes(api, userService)
	SetUpActivityRoutes(api, activityService)
}
