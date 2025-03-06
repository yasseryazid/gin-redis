package usecases

import (
	"github.com/yasseryazid/technical-test/models"
	"github.com/yasseryazid/technical-test/repositories"
)

type TaskService struct {
	Repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetTasks(status, search string, page, limit int) ([]models.Task, int, error) {
	return s.Repo.GetTasks(status, search, page, limit)
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.Repo.CreateTask(task)
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.Repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(id uint, updatedTask *models.Task) error {
	return s.Repo.UpdateTask(id, updatedTask)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.Repo.DeleteTask(id)
}
