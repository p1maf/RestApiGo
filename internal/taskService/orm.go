package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
	UserId uint   `json:"userId"`
}
