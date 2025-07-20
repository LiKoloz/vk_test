package repository

import (
	. "vk/Announcement/models"

	"gorm.io/gorm"
)

func GetPaginationCreateAt(g *gorm.DB, page int) (announcements []Announcement, err2 error) {
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("CreatedAt")
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}

func GetPaginationCreateAtDesc(g *gorm.DB, page int) (announcements []Announcement, err2 error) {
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("created_at DESC")
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}

func GetPaginationPrice(g *gorm.DB, page int) (announcements []Announcement, err2 error) {
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("price")
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}

func GetPaginationPriceAtDesc(g *gorm.DB, page int) (announcements []Announcement, err2 error) {
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("price DESC")
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}

func AddAnnouncement(g *gorm.DB, an Announcement) (ann Announcement, err2 error) {
	result := g.Model(&Announcement{}).Create(&an)

	if result.Error != nil {
		return Announcement{}, result.Error
	}
	return an, nil
}
