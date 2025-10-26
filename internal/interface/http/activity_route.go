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
	activityRoutes.POST("/join", activityHandler.JoinActivity)
	activityRoutes.GET("/by-user", activityHandler.ListActivitiesByUser)
	activityRoutes.DELETE("/delete_activity/:id", activityHandler.DeleteActivity)
	activityRoutes.DELETE("/delete_mem_activity", activityHandler.RemoveActivityMember)
}
