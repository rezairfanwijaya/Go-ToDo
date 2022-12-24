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
	r.POST("/activity-groups", activityHandler.CreateActivity)
}
