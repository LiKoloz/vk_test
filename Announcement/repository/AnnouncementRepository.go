package repository

import (
	. "vk/Announcement/models"

	"gorm.io/gorm"
)

func GetPaginationCreateAt(g *gorm.DB, page int, id uint) ([]AnnouncementDTO, error) {

	var (
		announcements []Announcement
		res_ann       []AnnouncementDTO
	)

	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("CreatedAt")
	if result.Error != nil {
		return nil, result.Error
	}

	for _, v := range announcements {
		var is_author = v.UserRefer == id

		res_ann = append(res_ann, AnnouncementDTO{
			Title:    v.Text,
			Text:     v.Text,
			Image:    v.Image,
			Price:    v.Price,
			IsAuthor: is_author,
		})

	}
	return res_ann, nil
}

func GetPaginationCreateAtDesc(g *gorm.DB, page int, id uint) ([]AnnouncementDTO, error) {
	var (
		announcements []Announcement
		res_ann       []AnnouncementDTO
	)
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("created_at DESC")
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range announcements {
		var is_author = v.UserRefer == id

		res_ann = append(res_ann, AnnouncementDTO{
			Title:    v.Text,
			Text:     v.Text,
			Image:    v.Image,
			Price:    v.Price,
			IsAuthor: is_author,
		})

	}
	return res_ann, nil
}

func GetPaginationPrice(g *gorm.DB, page int, id uint) ([]AnnouncementDTO, error) {
	var (
		announcements []Announcement
		res_ann       []AnnouncementDTO
	)
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("price")
	if result.Error != nil {
		return nil, result.Error
	}

	for _, v := range announcements {
		var is_author = v.UserRefer == id

		res_ann = append(res_ann, AnnouncementDTO{
			Title:    v.Text,
			Text:     v.Text,
			Image:    v.Image,
			Price:    v.Price,
			IsAuthor: is_author,
		})

	}
	return res_ann, nil
}

func GetPaginationPriceAtDesc(g *gorm.DB, page int, id uint) ([]AnnouncementDTO, error) {
	var (
		announcements []Announcement
		res_ann       []AnnouncementDTO
	)
	result := g.Model(&Announcement{}).Offset(int((page - 1) * 5)).Limit(5).Find(&announcements).Order("price DESC")
	if result.Error != nil {
		return nil, result.Error
	}

	for _, v := range announcements {
		var is_author = v.UserRefer == id

		res_ann = append(res_ann, AnnouncementDTO{
			Title:    v.Text,
			Text:     v.Text,
			Image:    v.Image,
			Price:    v.Price,
			IsAuthor: is_author,
		})

	}
	return res_ann, nil
}

func AddAnnouncement(g *gorm.DB, an Announcement) (ann Announcement, err2 error) {
	result := g.Model(&Announcement{}).Create(&an)

	if result.Error != nil {
		return Announcement{}, result.Error
	}
	return an, nil
}
