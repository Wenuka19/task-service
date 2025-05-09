package http

import (
	"fmt"
	"github.com/Wenuka19/task-service/internal/application/service"
	"github.com/Wenuka19/task-service/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: taskService}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/tasks")
	{
		group.POST("/", h.CreateTask)
		group.GET("/", h.GetTasks)
		group.PUT("/:id/status", h.UpdateTaskStatus)
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.TaskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.TaskService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	var input struct {
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := h.TaskService.UpdateStatus(id, input.Completed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task status updated"})
}
