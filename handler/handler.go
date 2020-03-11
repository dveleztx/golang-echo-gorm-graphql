package handler

import (
	"stack/domain/model"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u []*model.User

		if err := db.Find(&u).Error; err != nil {
			//error handling
			return err
		}

		return c.JSON(http.StatusOK, u)
	}
}
