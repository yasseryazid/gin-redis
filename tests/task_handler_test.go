package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
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

// ✅ Test Get All Tasks API
func TestGetAllTasks(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler() // ✅ Ensure handler is properly initialized
	router.GET("/api/tasks", taskHandler.GetTasks)

	// Create a test request
	req, _ := http.NewRequest("GET", "/api/tasks?page=1&limit=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// ✅ Assert status code
	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")

	// ✅ Parse JSON response
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

	// ✅ Assert response structure
	assert.NotNil(t, response.Tasks, "Tasks list should not be nil")
	assert.GreaterOrEqual(t, len(response.Tasks), 0, "Tasks list should have at least 0 elements")

	// ✅ Assert pagination fields
	assert.GreaterOrEqual(t, response.Pagination.TotalPages, 0, "Total pages should be >= 0")
	assert.GreaterOrEqual(t, response.Pagination.TotalTasks, 0, "Total tasks should be >= 0")

	// ✅ Assert that tasks contain expected fields
	if len(response.Tasks) > 0 {
		firstTask := response.Tasks[0]
		assert.NotZero(t, firstTask.ID, "Task ID should not be zero")
		assert.NotEmpty(t, firstTask.Title, "Task Title should not be empty")
		assert.NotEmpty(t, firstTask.Description, "Task Description should not be empty")
		assert.Contains(t, []string{"pending", "completed"}, firstTask.Status, "Task Status should be valid")
	}
}

func setupTestHandler() *handlers.TaskHandler {
	// ✅ Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("⚠️ Warning: .env file not found, using system environment variables")
	}

	// ✅ Set Gin to test mode to reduce logging
	gin.SetMode(gin.TestMode)

	// ✅ Initialize database if not already set
	if config.DB == nil {
		config.ConnectDatabase()
	}

	taskRepo := repositories.NewTaskRepository()
	taskService := usecases.NewTaskService(taskRepo)
	return &handlers.TaskHandler{Service: taskService}
}
