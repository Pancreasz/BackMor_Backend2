package http

import (
	"github.com/gin-gonic/gin"
)

func SetUpActivityRoutes(api gin.IRouter, activityService ActivityService) {
	activityHandler := NewActivityHandler(activityService)
	activityRoutes := api.Group("/activity")
	activityRoutes.GET("/", activityHandler.ListActivities)
	activityRoutes.GET("/id/:id", activityHandler.GetActivityByID)
	activityRoutes.POST("/", activityHandler.CreateActivity)
	activityRoutes.GET("/:id/members", activityHandler.GetActivityMembers)
	activityRoutes.POST("/:id/join", activityHandler.JoinActivity)
	activityRoutes.GET("/by-user", activityHandler.ListActivitiesByUser)
}
