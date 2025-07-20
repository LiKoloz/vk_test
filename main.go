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
	. "vk/User/services"

	// "github.com/labstack/echo-jwt/v4"
	"github.com/go-playground/validator"
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
	// e.Use(echojwt.JWT([]byte("secret")))
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("announcement/get/create_at/:uid/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad page")
		}

		uid, err := strconv.Atoi(c.Param("uid"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad page")
		}

		announcements, err := GetPaginationCreateAt(db, page, uint(uid))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request")
		}
		return c.JSON(http.StatusOK, announcements)
	})

	e.GET("announcement/get/create_at/desc/:uid/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad uid")
		}

		uid, err := strconv.Atoi(c.Param("uid"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad page")
		}

		announcements, err := GetPaginationCreateAt(db, page, uint(uid))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request")
		}
		return c.JSON(http.StatusOK, announcements)
	})

	e.GET("announcement/get/price/:uid/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad uid")
		}

		uid, err := strconv.Atoi(c.Param("uid"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad page")
		}

		announcements, err := GetPaginationPrice(db, page, uint(uid))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request")
		}
		return c.JSON(http.StatusOK, announcements)
	})

	e.GET("announcement/get/price/desc/:uid/:page", func(c echo.Context) error {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad uid")
		}

		uid, err := strconv.Atoi(c.Param("uid"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad page")
		}

		announcements, err := GetPaginationPriceAtDesc(db, page, uint(uid))
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request")
		}
		return c.JSON(http.StatusOK, announcements)
	})

	e.POST("announcement/create", func(c echo.Context) error {
		a := new(AnnouncementDTO)
		if err = c.Bind(a); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(a); err != nil {
			return err
		}
		fmt.Println(a.UserRefer)
		res, err := AddAnnouncement(db, Announcement{
			Title:     a.Title,
			Text:      a.Text,
			Image:     a.Image,
			Price:     a.Price,
			UserRefer: a.UserRefer,
		})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusBadRequest, res)
	})

	e.POST("user/registration", func(c echo.Context) error {

		u := new(UserDTO)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		res, err := Registration(db, u.Login, u.Password)
		if err != "" {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("user/sign_in/:login/:password", func(c echo.Context) error {
		login := c.Param("login")
		password := c.Param("password")

		res, err := SignIn(db, login, password)

		if err != "" {
			return c.String(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, res)
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
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type CustomValidator struct {
	validator *validator.Validate
}
