package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yasseryazid/technical-test/handlers"
	"github.com/yasseryazid/technical-test/middlewares"
	"github.com/yasseryazid/technical-test/repositories"
)

func RegisterAPIRoutes(router *gin.Engine, taskHandler *handlers.TaskHandler) {
	api := router.Group("/api")

	// ✅ Register authentication routes
	userRepo := repositories.NewUserRepository()
	authHandler := &handlers.AuthHandler{UserRepo: userRepo} // Use pointer

	api.POST("/register", authHandler.Register) // Public route
	api.POST("/login", authHandler.Login)       // Public route

	// ✅ Secure Task Routes with JWT Middleware
	taskRoutes := api.Group("/tasks")
	taskRoutes.Use(middlewares.AuthMiddleware()) // Apply JWT auth middleware
	{
		RegisterTaskRoutes(taskRoutes, taskHandler) // Use existing function
	}
}
