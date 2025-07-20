package repository

import (
	. "vk/User/models"

	"gorm.io/gorm"
)

func GetUserByLogin(g *gorm.DB, login string, password string) (User, error) {
	var user User
	result := g.First(&user, "login = ?", login)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func GetUserByLoginAndPassword(g *gorm.DB, login string, password string) (User, error) {
	var user User
	result := g.First(&user, "login = ? and password = ?", login, password)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
