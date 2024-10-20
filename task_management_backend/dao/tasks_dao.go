package dao

import (
	"fmt"
	"gorm.io/gorm"
	"task_management_backend/models"
)

type TaskDAO struct {
	DB *gorm.DB
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{DB: db}
}

func (dao *TaskDAO) CreateTask(task *models.Task) error {
	result := dao.DB.Create(task)
	if result.Error != nil {
		return fmt.Errorf("failed to create task: %w", result.Error)
	}
	return nil
}

func (dao *TaskDAO) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	result := dao.DB.First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task not found")
		}
		return nil, fmt.Errorf("failed to retrieve task: %w", result.Error)
	}
	return &task, nil
}

func (dao *TaskDAO) GetAllTasks(limit int, offset int) ([]models.Task, error) {
	var tasks []models.Task

	result := dao.DB.Limit(limit).Offset(offset).Find(&tasks)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %w", result.Error)
	}

	return tasks, nil
}

func (dao *TaskDAO) GetTotalTaskCount() (int64, error) {
	var count int64
	result := dao.DB.Model(&models.Task{}).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to count tasks: %w", result.Error)
	}
	return count, nil
}

func (dao *TaskDAO) UpdateTask(task *models.Task) error {
	result := dao.DB.Save(task)
	if result.Error != nil {
		return fmt.Errorf("failed to update task: %w", result.Error)
	}
	return nil
}

func (dao *TaskDAO) DeleteTask(id uint) error {
	result := dao.DB.Delete(&models.Task{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete task: %w", result.Error)
	}
	return nil
}
