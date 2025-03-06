package routes

import (
	"github.com/yasseryazid/technical-test/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(api *gin.RouterGroup, taskHandler *handlers.TaskHandler) {
	{
		api.GET("", taskHandler.GetTasks)
		api.POST("", taskHandler.CreateTask)
		api.GET("/:id", taskHandler.GetTaskByID)
		api.PUT("/:id", taskHandler.UpdateTask)
		api.DELETE("/:id", taskHandler.DeleteTask)
	}
}
