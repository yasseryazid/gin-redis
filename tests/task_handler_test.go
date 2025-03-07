package tests

import (
	"bytes"
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

func TestCreateTask(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.POST("/api/tasks", taskHandler.CreateTask)

	task := models.Task{
		Title:       "New Test Task",
		Description: "This is a test task",
		Status:      "pending",
		DueDate:     "2025-04-01",
	}

	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status 201 Created")

	var response struct {
		Message string      `json:"message"`
		Task    models.Task `json:"task"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response JSON should be valid")
	assert.Equal(t, "Task created successfully", response.Message)
	assert.NotZero(t, response.Task.ID, "Task ID should not be zero")
}

func TestGetTaskByID(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.GET("/api/tasks/:id", taskHandler.GetTaskByID)

	req, _ := http.NewRequest("GET", "/api/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")

	var response models.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response JSON should be valid")
	assert.NotZero(t, response.ID, "Task ID should not be zero")
}

func TestUpdateTask(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.PUT("/api/tasks/:id", taskHandler.UpdateTask)

	updateTask := models.Task{
		Title:       "Updated Test Task",
		Description: "Updated description",
		Status:      "completed",
		DueDate:     "2025-04-10",
	}

	body, _ := json.Marshal(updateTask)
	req, _ := http.NewRequest("PUT", "/api/tasks/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")

	var response struct {
		Message string      `json:"message"`
		Task    models.Task `json:"task"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response JSON should be valid")
	assert.Equal(t, "Task updated successfully", response.Message)
	assert.Equal(t, "Updated Test Task", response.Task.Title)
	assert.Equal(t, "completed", response.Task.Status)
}

func TestDeleteTask(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	req, _ := http.NewRequest("DELETE", "/api/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")

	var response struct {
		Message string `json:"message"`
	}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response JSON should be valid")
	assert.Equal(t, "Task deleted successfully", response.Message)
}

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
