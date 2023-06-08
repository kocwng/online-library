package delivery

import (
	"net/http"
	"strconv"
	// "online-library/features/authors"
	"online-library/features/books"
	"online-library/utils/helpers"

	"github.com/labstack/echo/v4"
)

// type BookHandler struct {
// 	bookService   books.BookServiceInterface
// 	authorService authors.AuthorServiceInterface
// }

// func New(srv books.BookServiceInterface, srv2 authors.AuthorServiceInterface) *BookHandler {
// 	return &BookHandler{
// 		bookService:   srv,
// 		authorService: srv2,
// 	}
// }

type BookHandler struct {
	bookService   books.BookServiceInterface
}

func New(srv books.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService:   srv,
	}
}

func (h *BookHandler) Create(c echo.Context) error {
	var formInput BookRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	book, err := h.bookService.Create(ToEntity(formInput))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	bookEntity, _ := h.bookService.GetById(book.Id)
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ToResponse(bookEntity)))
}

func (h *BookHandler) Update(c echo.Context) error {
	var formInput BookRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	feedback, err := h.bookService.Update(ToEntity(formInput), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", ToResponse(feedback)))
}

// func (h *BookHandler) Create(c echo.Context) error {
// 	var formInput BookRequest
// 	if err := c.Bind(&formInput); err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
// 	}

// 	bookEntity := ToEntity(formInput)
// 	book, err := h.bookService.Create(bookEntity)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
// 	}

// 	bookEntity, err = h.bookService.GetById(book.Id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
// 	}

// 	// Store authors in the authors table
// 	for _, author := range formInput.Author {
// 		authorEntity := authors.AuthorEntity{
// 			Name:    author.Name,
// 			Country: author.Country,
// 		}
// 		_, err := h.authorService.Create(authorEntity)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
// 		}
// 	}

// 	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ToResponse(bookEntity)))
// }
