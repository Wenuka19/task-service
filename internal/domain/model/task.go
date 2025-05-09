package model

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`
	AddedBy   string    `json:"added_by"`
	Completed bool      `gorm:"default:false" json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
