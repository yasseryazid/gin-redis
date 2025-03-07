package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
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

var createdTaskID uint

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

	createdTaskID = response.Task.ID
	assert.NotZero(t, createdTaskID, "Task ID should not be zero")
}

func TestGetTaskByID(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.GET("/api/tasks/:id", taskHandler.GetTaskByID)

	taskID := strconv.Itoa(int(createdTaskID))

	req, _ := http.NewRequest("GET", "/api/tasks/"+taskID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")

	var response models.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Response JSON should be valid")
	assert.Equal(t, createdTaskID, response.ID, "Task ID should match")
}

func TestUpdateTask(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.PUT("/api/tasks/:id", taskHandler.UpdateTask)

	taskID := strconv.Itoa(int(createdTaskID))

	updateTask := models.Task{
		Title:       "Updated Test Task",
		Description: "Updated description",
		Status:      "completed",
		DueDate:     "2025-04-10",
	}

	body, _ := json.Marshal(updateTask)
	req, _ := http.NewRequest("PUT", "/api/tasks/"+taskID, bytes.NewBuffer(body))
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
	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")
	assert.Equal(t, "Task updated successfully", response.Message)
	assert.Equal(t, "Updated Test Task", response.Task.Title)
	assert.Equal(t, "completed", response.Task.Status)
}

func TestDeleteTask(t *testing.T) {
	router := gin.Default()
	taskHandler := setupTestHandler()
	router.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	taskID := strconv.Itoa(int(createdTaskID))

	req, _ := http.NewRequest("DELETE", "/api/tasks/"+taskID, nil)
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
