package router

import (
	"gotodo/activity"
	"gotodo/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, r *gin.Engine) {
	// repo activity
	activityRepo := activity.NewActivityRespository(db)
	// service activity
	activityService := activity.NewActivityService(activityRepo)
	// handler activity
	activityHandler := handler.NewActivityHandler(activityService)

	// endpoints
	r.GET("/activity-groups", activityHandler.GetAllActivity)
	r.GET("/activity-groups/:id", activityHandler.GetActivityByID)
	r.POST("/activity-groups", activityHandler.CreateActivity)
	r.DELETE("/activity-groups/:id", activityHandler.DeleteByID)
	r.PATCH("/activity-groups/:id", activityHandler.UpdateByID)
}
