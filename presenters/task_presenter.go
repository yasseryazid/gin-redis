package presenters

import (
	"strconv"

	"github.com/yasseryazid/technical-test/models"
)

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

type TaskDetailResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

func FormatTask(task *models.Task) TaskResponse {
	return TaskResponse{
		ID:          strconv.FormatUint(uint64(task.ID), 10),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}
}

func FormatTaskList(tasks []models.Task) []TaskResponse {
	formattedTasks := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		formattedTasks[i] = FormatTask(&task)
	}
	return formattedTasks
}

func FormatTaskDetail(task *models.Task) TaskDetailResponse {
	return TaskDetailResponse{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}
}
