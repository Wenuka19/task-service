package port

import "github.com/Wenuka19/task-service/internal/domain/model"

type TaskRepository interface {
	CreateTask(task *model.Task) error
	FindAll() ([]model.Task, error)
	UpdateStatus(id uint, completed bool) error
}
