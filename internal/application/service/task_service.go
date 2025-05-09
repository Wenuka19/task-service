package service

import (
	"github.com/Wenuka19/task-service/internal/domain/model"
	"github.com/Wenuka19/task-service/internal/domain/port"
)

type TaskService struct {
	Repo port.TaskRepository
}

func NewTaskService(repo port.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (service *TaskService) CreateTask(task *model.Task) error {
	return service.Repo.CreateTask(task)
}

func (service *TaskService) FindAll() ([]model.Task, error) {
	return service.Repo.FindAll()
}

func (service *TaskService) UpdateStatus(id uint, completed bool) error {
	return service.Repo.UpdateStatus(id, completed)
}
