package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yasseryazid/technical-test/handlers"
	"github.com/yasseryazid/technical-test/middlewares"
	"github.com/yasseryazid/technical-test/repositories"
)

func RegisterAPIRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	api := router.Group("/api")

	userRepo := repositories.NewUserRepository()
	authHandler := &handlers.AuthHandler{UserRepo: userRepo}

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
	api.POST("/logout", authHandler.Logout)

	taskRoutes := api.Group("/tasks")
	taskRoutes.Use(middlewares.AuthMiddleware())
	{
		RegisterTaskRoutes(taskRoutes, taskHandler)
	}
}
