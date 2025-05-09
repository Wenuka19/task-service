package repository

import (
	"github.com/Wenuka19/task-service/internal/domain/model"
	"github.com/Wenuka19/task-service/internal/domain/port"
	"github.com/Wenuka19/task-service/internal/infrastructure/db"
)

type TaskRepositoryImpl struct{}

func NewTaskRepository() port.TaskRepository {
	return &TaskRepositoryImpl{}
}

func (t TaskRepositoryImpl) CreateTask(task *model.Task) error {
	return db.DB.Create(task).Error
}

func (t TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := db.DB.Find(&tasks).Error
	return tasks, err
}

func (t TaskRepositoryImpl) UpdateStatus(id uint, completed bool) error {
	return db.DB.Model(&model.Task{}).Where("id = ?", id).Update("completed", completed).Error
}
