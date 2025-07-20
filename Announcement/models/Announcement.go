package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	ID        uint           `json:"id"`
	Title     string         `gorm:"uniqueIndex;not null;" json:"title"`
	Text      string         `gorm:"not null;" json:"text"`
	Image     string         `gorm: "not null;" json: "image"`
	Price     float64        `gorm: "not null;" json: "price"`
	UserRefer uint           `gorm: "not null" json: "user_refer"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
