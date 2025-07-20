package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	. "vk/Announcement/models"
	. "vk/Announcement/repository"
	. "vk/User/models"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=postgres password=123 dbname=vk port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error to connect DB!")
	} else {
		fmt.Println("Db connect ready!")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Announcement{})

	GenerateMockData(db)

	e := echo.New()

	e.GET("announcement/get/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad page")
		}

		announcements, err := GetPaginationPriceAtDesc(db, page)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request")
		}
		return c.JSON(http.StatusOK, announcements)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

// GenerateMockData создает тестовых пользователей с объявлениями
func GenerateMockData(db *gorm.DB) {
	// Проверяем, есть ли уже данные
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем 5 пользователей
	for i := 1; i <= 5; i++ {
		user := User{
			Login:    fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("password%d", i),
		}

		db.Create(&user)

		// Создаем 3-5 объявлений для каждого пользователя
		numAnnouncements := 3 + rand.Intn(3) // От 3 до 5 объявлений
		for j := 1; j <= numAnnouncements; j++ {
			announcement := Announcement{
				Title:     fmt.Sprintf("Объявление %d-%d", i, j),
				Text:      fmt.Sprintf("Это текст объявления от пользователя %d, номер %d", i, j),
				Image:     fmt.Sprintf("image%d.jpg", j),
				Price:     float64(100*i + 10*j),
				UserRefer: user.ID,
				CreatedAt: time.Now().Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour),
			}
			db.Create(&announcement)
		}
	}
}
