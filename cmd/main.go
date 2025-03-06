package main

import (
	"log"

	"github.com/yasseryazid/technical-test/config"
	"github.com/yasseryazid/technical-test/handlers"
	"github.com/yasseryazid/technical-test/migrations"
	"github.com/yasseryazid/technical-test/repositories"
	"github.com/yasseryazid/technical-test/routes"
	"github.com/yasseryazid/technical-test/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	config.ConnectDatabase()
	migrations.RunMigration()

	router := gin.Default()

	taskRepo := repositories.NewTaskRepository()
	taskService := usecases.NewTaskService(taskRepo)
	taskHandler := &handlers.TaskHandler{Service: taskService}

	routes.RegisterAPIRoutes(router, taskHandler)

	log.Println("🚀 Server running on port 3000")
	log.Fatal(router.Run(":3000"))
}
