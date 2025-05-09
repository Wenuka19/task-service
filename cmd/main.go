package main

import (
	"github.com/Wenuka19/task-service/internal/application/service"
	"github.com/Wenuka19/task-service/internal/config"
	_ "github.com/Wenuka19/task-service/internal/infrastructure/db"
	"github.com/Wenuka19/task-service/internal/infrastructure/repository"
	"github.com/Wenuka19/task-service/internal/interfaces/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := http.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(r)

	err := r.Run(":" + config.AppConfig.Port)

	if err != nil {
		return
	}
}
