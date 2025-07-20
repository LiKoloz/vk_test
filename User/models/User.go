package models

import (
	"time"

	"gorm.io/gorm"

	"vk/Announcement/models"
)

type User struct {
	ID            uint                  `json:"id"`
	Login         string                `gorm:"uniqueIndex;not null;size:100" json:"login"`
	Password      string                `gorm:"not null;" json:"password"`
	Announcements []models.Announcement `gorm:"foreignKey:UserRefer"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	DeletedAt     gorm.DeletedAt        `gorm:"index" json:"deleted_at,omitempty"`
}

type UserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
