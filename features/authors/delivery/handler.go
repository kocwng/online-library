package delivery

import (
	"net/http"
	"online-library/features/authors"
	"online-library/utils/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	authorService authors.AuthorServiceInterface
}

func New(srv authors.AuthorServiceInterface) *AuthorHandler {
	return &AuthorHandler{
		authorService: srv,
	}
}

func (h *AuthorHandler) Create(c echo.Context) error {
	var formInput AuthorRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	author, err := h.authorService.Create(ToEntity(formInput))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	authorEntity, _ := h.authorService.GetById(author.Id)
	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ToResponse(authorEntity)))
}

func (h *AuthorHandler) Update(c echo.Context) error {
	var formInput AuthorRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	feedback, err := h.authorService.Update(ToEntity(formInput), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", ToResponse(feedback)))
}

func (h *AuthorHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, errId := h.authorService.GetById(uint(id))
	if errId != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(errId.Error()))
	}

	if err := h.authorService.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}