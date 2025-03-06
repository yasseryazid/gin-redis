package routes

import (
	"github.com/yasseryazid/technical-test/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskHandler.GetTasks)
		tasks.POST("", taskHandler.CreateTask)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}
}
