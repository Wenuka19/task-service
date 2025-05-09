package db

import (
	"fmt"
	"github.com/Wenuka19/task-service/internal/config"
	"github.com/Wenuka19/task-service/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	InitDatabase()
	//RunMigrations()
	log.Println("Database Initialized successfully.")
}

func InitDatabase() {
	cfg := config.AppConfig
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Database connection established.")
}

func RunMigrations() {
	err := DB.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully.")
}
