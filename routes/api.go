package routes

import (
	"github.com/yasseryazid/technical-test/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	api := router.Group("/api")

	RegisterTaskRoutes(api, taskHandler)
}
