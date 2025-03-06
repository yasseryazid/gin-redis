package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"type:varchar(50);default:'pending'" json:"status"`
	DueDate     string    `gorm:"type:date" json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
}
