package delivery

import (
	"net/http"
	"strconv"

	// "online-library/features/authors"

	// bookauthors "online-library/features/bookAuthors/data"
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
	bookService books.BookServiceInterface
}

func New(srv books.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: srv,
	}
}

// func (h *BookHandler) CreateBook(c echo.Context) error {
// 	var request BookRequest
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
// 	}

// 	if err := c.Validate(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
// 	}

// 	bookEntity := request.ToEntity()
// 	id, err := h.bookService.Create(bookEntity)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create book"})
// 	}

// 	return c.JSON(http.StatusCreated, map[string]uint{"id": id})
// }

// func (h *BookHandler) Create(c echo.Context) error {
// 	var formInput BookRequest
// 	if err := c.Bind(&formInput); err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
// 	}

// 	book, err := h.bookService.Create(ToEntity(formInput))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
// 	}

// 	bookEntity, _ := h.bookService.GetById(book.Id)
// 	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ToResponse(bookEntity)))
// }

// func (h *BookHandler) Create(c echo.Context) error {
// 	var formInput BookRequest
// 	if err := c.Bind(&formInput); err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
// 	}

// 	// Store book data
// 	book, err := h.bookService.Create(ToEntity(formInput))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
// 	}

// 	// bookAuthor := bookauthors.ToEntity(&formInput)

// 	// Store author associations
// 	for _, authorID := range formInput.AuthorId {
// 		if err := h.bookService.AddAuthorAssociation(book.Id, authorID); err != nil {
// 			return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
// 		}
// 	}

// 	bookEntity, _ := h.bookService.GetById(book.Id)
// 	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ToResponse(bookEntity)))
// }

func (h *BookHandler) Create(c echo.Context) error {
    var formInput BookRequest
    if err := c.Bind(&formInput); err != nil {
        return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
    }

    // Store book data
    book, err := h.bookService.Create(ToEntity(formInput))
    if err != nil {
        return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
    }

    // Store author associations
    for _, authorID := range formInput.AuthorId {
        if err := h.bookService.AddAuthorAssociation(book.Id, authorID); err != nil {
            return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
        }
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

func (h *BookHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, errId := h.bookService.GetById(uint(id))
	if errId != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(errId.Error()))
	}

	if err := h.bookService.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}

func (h *BookHandler) GetBooksByAuthor(c echo.Context) error {
	authorID,_ := strconv.Atoi(c.Param("id"))

	books, err := h.bookService.GetBooksByAuthor(uint(authorID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Get Books By Author Success", books))
}