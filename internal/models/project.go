package models

import (
	"time"

	"gorm.io/gorm"
)

// Project represents a portfolio project
type Project struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	ImageURL    string         `json:"image_url"`
	TechStack   string         `gorm:"type:text" json:"tech_stack"` // JSON array as string
	DemoLink    string         `json:"demo_link"`
	GithubLink  string         `json:"github_link"`
}

