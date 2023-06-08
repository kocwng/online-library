package router

import (
	_booksData "online-library/features/books/data"
	_booksHandler "online-library/features/books/delivery"
	_booksService "online-library/features/books/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BookRouter(db *gorm.DB, e *echo.Echo) {
	
	data := _booksData.New(db)
	service := _booksService.New(data)
	handler := _booksHandler.New(service)

	g := e.Group("/books")
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}