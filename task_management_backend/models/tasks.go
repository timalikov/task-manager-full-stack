package models

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:500"`
	Status      string `gorm:"size:50;default:'pending'"`
	Priority    int    `gorm:"not null;default:3"`
	Deadline    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
