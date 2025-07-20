package services

import (
	. "vk/User/models"
	. "vk/User/repository"

	"gorm.io/gorm"
)

func Registration(g *gorm.DB, login string, password string) (User, string) {
	_, err := GetUserByLogin(g, login, password)
	if err == nil {
		return User{}, "Already exists!"
	}
	user := User{
		Login:    login,
		Password: password,
	}

	g.Create(&user)

	return user, ""
}

func SignIn(g *gorm.DB, login string, password string) (User, string) {
	u, err := GetUserByLoginAndPassword(g, login, password)
	if err != nil {
		return User{}, "User doesn't exist!"
	}
	return u, ""
}
