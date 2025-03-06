package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/yasseryazid/technical-test/models"
	"github.com/yasseryazid/technical-test/presenters"
	"github.com/yasseryazid/technical-test/usecases"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Service *usecases.TaskService
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	status, search, page, limit := parseQueryParams(c)

	var tasks []models.Task
	var total int
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(2)

	// Fetch tasks concurrently
	go func() {
		defer wg.Done()
		var err error
		tasks, _, err = h.Service.GetTasks(status, search, page, limit)
		if err != nil {
			log.Printf("[X] Error fetching tasks: %v\n", err)
			errChan <- err
		}
	}()

	// Fetch total task count concurrently
	go func() {
		defer wg.Done()
		var err error
		_, total, err = h.Service.GetTasks(status, search, 1, 1) // Only count total
		if err != nil {
			log.Printf("[X] Error counting tasks: %v\n", err)
			errChan <- err
		}
	}()

	// Wait for Goroutines to finish
	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		if err != nil {
			log.Printf("[X] Failed to fetch tasks: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
			return
		}
	}

	totalPages := (total + limit - 1) / limit

	log.Printf("[V] Successfully fetched tasks (page %d, limit %d)\n", page, limit)
	c.JSON(http.StatusOK, gin.H{
		"tasks": presenters.FormatTaskList(tasks),
		"pagination": gin.H{
			"current_page": page,
			"total_pages":  totalPages,
			"total_tasks":  total,
		},
	})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		log.Printf("[X] Invalid request body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validateTask(&task); err != nil {
		log.Printf("[X] Task validation failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateTask(&task); err != nil {
		log.Printf("[X] Failed to create task: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	log.Printf("[V] Task created successfully: ID %d\n", task.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    presenters.FormatTask(&task),
	})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		log.Printf("[X] Invalid task ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		log.Printf("[X] Task not found (ID %d): %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	log.Printf("[V] Task retrieved: ID %d\n", id)
	c.JSON(http.StatusOK, presenters.FormatTask(task))
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		log.Printf("[X] Invalid task ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		log.Printf("[X] Invalid request body: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validateTask(&updatedTask); err != nil {
		log.Printf("[X] Task validation failed: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateTask(id, &updatedTask); err != nil {
		log.Printf("[X] Task update failed (ID %d): %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	log.Printf("[V] Task updated successfully: ID %d\n", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    presenters.FormatTask(&updatedTask),
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		log.Printf("[X] Invalid task ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := h.Service.DeleteTask(id); err != nil {
		log.Printf("[X] Task deletion failed (ID %d): %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	log.Printf("[V] Task deleted successfully: ID %d\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func parseQueryParams(c *gin.Context) (string, string, int, int) {
	status := c.Query("status")
	search := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return status, search, page, limit
}

func parseIDParam(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		return 0, err
	}
	return uint(id), nil
}

func validateTask(task *models.Task) error {
	if task.Title == "" {
		return fmt.Errorf("Title is required")
	}
	if task.Status != "pending" && task.Status != "completed" {
		return fmt.Errorf("Invalid status. Use 'pending' or 'completed'")
	}
	return nil
}
