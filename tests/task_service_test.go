package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yasseryazid/technical-test/models"
	"github.com/yasseryazid/technical-test/usecases"
)

// Mock repository
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTasks(status, search string, page, limit int) ([]models.Task, int, error) {
	args := m.Called(status, search, page, limit)
	return args.Get(0).([]models.Task), args.Int(1), args.Error(2)
}

func (m *MockTaskRepository) CreateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTaskByID(id uint) (*models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id uint, updatedTask *models.Task) error {
	args := m.Called(id, updatedTask)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// ✅ Test Create Task
func Test_CreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := usecases.NewTaskService(mockRepo)

	task := &models.Task{
		Title:       "Test Task",
		Description: "Testing task creation",
		Status:      "pending",
		DueDate:     time.Now().Format("2006-01-02"),
	}

	mockRepo.On("CreateTask", task).Return(nil)

	err := service.CreateTask(task)
	assert.Nil(t, err, "Expected no error when creating task")
	mockRepo.AssertExpectations(t)
}

// ✅ Test Get Task by ID
func Test_GetTaskByID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := usecases.NewTaskService(mockRepo)

	task := &models.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Testing",
		Status:      "pending",
		DueDate:     "2025-03-07",
	}

	mockRepo.On("GetTaskByID", uint(1)).Return(task, nil)

	result, err := service.GetTaskByID(1)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, task, result, "Task should match the expected value")
	mockRepo.AssertExpectations(t)
}

// ✅ Test Update Task
func Test_UpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := usecases.NewTaskService(mockRepo)

	taskID := uint(1)
	updatedTask := &models.Task{
		Title:       "Updated Task",
		Description: "Updated description",
		Status:      "completed",
		DueDate:     "2025-04-01",
	}

	mockRepo.On("UpdateTask", taskID, updatedTask).Return(nil)

	err := service.UpdateTask(taskID, updatedTask)
	assert.Nil(t, err, "Expected no error when updating task")
	mockRepo.AssertExpectations(t)
}

// ✅ Test Delete Task
func Test_DeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := usecases.NewTaskService(mockRepo)

	taskID := uint(1)
	mockRepo.On("DeleteTask", taskID).Return(nil)

	err := service.DeleteTask(taskID)
	assert.Nil(t, err, "Expected no error when deleting task")
	mockRepo.AssertExpectations(t)
}

// ✅ Test Get All Tasks
func TestGetTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := usecases.NewTaskService(mockRepo)

	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Description: "Task 1 Desc", Status: "pending", DueDate: "2025-03-10"},
		{ID: 2, Title: "Task 2", Description: "Task 2 Desc", Status: "completed", DueDate: "2025-03-12"},
	}

	mockRepo.On("GetTasks", "", "", 1, 5).Return(tasks, len(tasks), nil)

	result, total, err := service.GetTasks("", "", 1, 5)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, len(tasks), total, "Total tasks should match")
	assert.Equal(t, tasks, result, "Returned tasks should match the expected value")
	mockRepo.AssertExpectations(t)
}

// ✅ Test Error Handling
func TestGetTaskByID_NotFound(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := usecases.NewTaskService(mockRepo)

	mockRepo.On("GetTaskByID", uint(99)).Return((*models.Task)(nil), errors.New("record not found"))

	result, err := service.GetTaskByID(99)
	assert.Nil(t, result, "Result should be nil when task is not found")
	assert.NotNil(t, err, "Error should not be nil when task is not found")
	assert.Equal(t, "record not found", err.Error(), "Error message should match")
	mockRepo.AssertExpectations(t)
}
