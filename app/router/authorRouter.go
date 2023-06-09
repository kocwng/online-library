package router

import (
	_authorsData "online-library/features/authors/data"
	_authorsHandler "online-library/features/authors/delivery"
	_authorsService "online-library/features/authors/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthorRouter(db *gorm.DB, e *echo.Echo) {
	
	data := _authorsData.New(db)
	service := _authorsService.New(data)
	handler := _authorsHandler.New(service)

	g := e.Group("/authors")
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
	g.GET("/:id", handler.GetAuthorsByBook)
}