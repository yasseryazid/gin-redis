package migrations

import (
	"fmt"

	"github.com/yasseryazid/technical-test/config"
	"github.com/yasseryazid/technical-test/models"
)

func RunMigration() {
	if err := config.DB.AutoMigrate(&models.Task{}); err != nil {
		fmt.Println("❌ Migration failed:", err)
		return
	}
	fmt.Println("✅ Migration successful")

	insertDummyIntoTaskTable()
}

func insertDummyIntoTaskTable() {
	var count int64
	config.DB.Model(&models.Task{}).Count(&count)

	if count > 0 {
		fmt.Println("⚠️ Dummy data already exists, skipping insertion")
		return
	}

	dummyTasks := []models.Task{
		{Title: "Task 1", Description: "Description for Task 1", Status: "pending", DueDate: "2025-03-10"},
		{Title: "Task 2", Description: "Description for Task 2", Status: "completed", DueDate: "2025-03-12"},
		{Title: "Task 3", Description: "Description for Task 3", Status: "pending", DueDate: "2025-03-15"},
	}

	if err := config.DB.Create(&dummyTasks).Error; err != nil {
		fmt.Println("❌ Failed to insert dummy data:", err)
		return
	}

	fmt.Println("✅ Dummy data inserted into 'tasks' table")
}
