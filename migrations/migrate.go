package migrations

import (
	"fmt"

	"github.com/yasseryazid/technical-test/config"
	"github.com/yasseryazid/technical-test/models"
)

func RunMigration() {
	if err := config.DB.AutoMigrate(&models.Task{}, &models.User{}); err != nil {
		fmt.Println("[X] Migration failed:", err)
		return
	}
	fmt.Println("[V] Migration successful")

	insertDummyIntoTaskTable()
	insertDummyIntoUserTable()
}

func insertDummyIntoTaskTable() {
	var count int64
	config.DB.Model(&models.Task{}).Count(&count)

	if count > 0 {
		fmt.Println("[!] Dummy Tasks data already exists, skipping insertion")
		return
	}

	dummyTasks := []models.Task{
		{Title: "Task 1", Description: "Description for Task 1", Status: "pending", DueDate: "2025-03-10"},
		{Title: "Task 2", Description: "Description for Task 2", Status: "completed", DueDate: "2025-03-12"},
		{Title: "Task 3", Description: "Description for Task 3", Status: "pending", DueDate: "2025-03-15"},
	}

	if err := config.DB.Create(&dummyTasks).Error; err != nil {
		fmt.Println("[X] Failed to insert dummy data:", err)
		return
	}

	fmt.Println("[V] Dummy Tasks data inserted into 'tasks' table")
}

func insertDummyIntoUserTable() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)

	if count > 0 {
		fmt.Println("[!] Dummy User data already exists, skipping insertion")
		return
	}

	dummyUsers := []models.User{
		{Username: "admin", Password: "$2a$10$7QjtOH3oEj0PbTtrO7H7R.hHUpYV1I4L5fWfE3hXZl0R/RY3LfLKm"}, // Password: "password"
	}

	if err := config.DB.Create(&dummyUsers).Error; err != nil {
		fmt.Println("[X] Failed to insert dummy user data:", err)
		return
	}

	fmt.Println("[V] Dummy user data inserted into 'users' table")
}
