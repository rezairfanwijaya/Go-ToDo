package router

import (
	"gotodo/activity"
	"gotodo/handler"
	"gotodo/todo"

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

	// repo todo
	todoRepo := todo.NewTodoRepository(db)
	// service todo
	todoService := todo.NewTodoService(todoRepo, activityService)
	// handler todo
	todoHandler := handler.NewTodoHandler(todoService)

	// endpoints
	r.GET("/activity-groups", activityHandler.GetAllActivity)
	r.GET("/activity-groups/:id", activityHandler.GetActivityByID)
	r.POST("/activity-groups", activityHandler.CreateActivity)
	r.DELETE("/activity-groups/:id", activityHandler.DeleteByID)
	r.PATCH("/activity-groups/:id", activityHandler.UpdateByID)

	r.GET("/todo-items/:id", todoHandler.GetTodoByID)
	r.GET("/todo-items", todoHandler.GetAllTodo)
	r.POST("/todo-items", todoHandler.CreateTodo)
}
