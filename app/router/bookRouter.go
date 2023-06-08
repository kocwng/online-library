package router

import (
	_discussionsData "online-library/features/books/data"
	_discussionsHandler "online-library/features/books/delivery"
	_discussionsService "online-library/features/books/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookRouter(db *gorm.DB, e *echo.Echo) {
	
	data := _discussionsData.New(db)
	service := _discussionsService.New(data)
	handler := _discussionsHandler.New(service)

	g := e.Group("/discussions")
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	
}