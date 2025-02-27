package taskService

import (
	"fmt"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskById(id uint, task Task) (Task, error)
	DeleteTaskById(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTaskById(id uint, task Task) (Task, error) {
	result := r.db.Model(&Task{}).Where("id = ?", id).Updates(task)

	if result.RowsAffected == 0 {
		return Task{}, result.Error
	}

	if result.Error != nil {
		return Task{}, result.Error
	}

	var updatedTask Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepository) DeleteTaskById(id uint) error {
	result := r.db.Unscoped().Delete(&Task{}, id)

	if result.RowsAffected == 0 {
		return fmt.Errorf("task not found")
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}
