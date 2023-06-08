package router

import (
	"net/http"
	"online-library/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	e.GET("/", index)
	BookRouter(db, e)

}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to Starter Pack Echo",
		"documentation": "/swagger/index.html",
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", data))
}