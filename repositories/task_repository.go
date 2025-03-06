package repositories

import (
	"github.com/yasseryazid/technical-test/config"
	"github.com/yasseryazid/technical-test/models"
)

type TaskRepository interface {
	GetTasks(status, search string, page, limit int) ([]models.Task, int, error)
	CreateTask(task *models.Task) error
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(id uint, updatedTask *models.Task) error
	DeleteTask(id uint) error
}

type taskRepository struct{}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (r *taskRepository) GetTasks(status, search string, page, limit int) ([]models.Task, int, error) {
	offset := (page - 1) * limit
	var tasks []models.Task
	query := config.DB.Model(&models.Task{})

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	result := query.Limit(limit).Offset(offset).Order("id DESC").Find(&tasks)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var total int64
	query.Count(&total)

	return tasks, int(total), nil
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func (r *taskRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	result := config.DB.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *taskRepository) UpdateTask(id uint, updatedTask *models.Task) error {
	var task models.Task

	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	task.DueDate = updatedTask.DueDate

	return config.DB.Save(&task).Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	}

	return config.DB.Delete(&task).Error
}
