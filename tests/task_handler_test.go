package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yasseryazid/technical-test/config"
	"github.com/yasseryazid/technical-test/handlers"
	"github.com/yasseryazid/technical-test/models"
	"github.com/yasseryazid/technical-test/repositories"
	"github.com/yasseryazid/technical-test/usecases"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTasks(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.GET("/api/tasks", taskHandler.GetTasks)

	req, _ := http.NewRequest("GET", "/api/tasks?page=1&limit=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")

	var response struct {
		Tasks      []models.Task `json:"tasks"`
		Pagination struct {
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
			TotalTasks  int `json:"total_tasks"`
		} `json:"pagination"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response JSON should be valid")

	assert.NotNil(t, response.Tasks, "Tasks list should not be nil")
	assert.GreaterOrEqual(t, len(response.Tasks), 0, "Tasks list should have at least 0 elements")

	assert.GreaterOrEqual(t, response.Pagination.TotalPages, 0, "Total pages should be >= 0")
	assert.GreaterOrEqual(t, response.Pagination.TotalTasks, 0, "Total tasks should be >= 0")

	if len(response.Tasks) > 0 {
		firstTask := response.Tasks[0]
		assert.NotZero(t, firstTask.ID, "Task ID should not be zero")
		assert.NotEmpty(t, firstTask.Title, "Task Title should not be empty")
		assert.NotEmpty(t, firstTask.Description, "Task Description should not be empty")
		assert.Contains(t, []string{"pending", "completed"}, firstTask.Status, "Task Status should be valid")
	}
}

func setupTestHandler() *handlers.TaskHandler {
	projectRoot, _ := os.Getwd()
	envPath := filepath.Join(projectRoot, "../.env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("[!] Warning: .env file not found at %s, using system environment variables", envPath)
	} else {
		log.Println("[V] .env file loaded successfully")
	}

	gin.SetMode(gin.TestMode)

	if config.DB == nil {
		config.ConnectDatabase()
	}

	taskRepo := repositories.NewTaskRepository()
	taskService := usecases.NewTaskService(taskRepo)
	return &handlers.TaskHandler{Service: taskService}
}
