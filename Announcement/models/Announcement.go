package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	ID        uint           `json:"id"`
	Title     string         `gorm:"not null;" json:"title"`
	Text      string         `gorm:"not null;" json:"text"`
	Image     string         `gorm: "not null;" json: "image"`
	Price     float64        `gorm: "not null;" json: "price"`
	UserRefer uint           `gorm: "not null" json: "user_refer"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type AnnouncementDTO struct {
	Title     string  `json:"title"`
	Text      string  `json:"text"`
	Image     string  `json: "image"`
	Price     float64 `json: "price"`
	UserRefer uint    `json: "user_refer"`
	IsAuthor  bool    `json: "is_author"`
}
