package handler

import (
	"hexagonal-architecture/internal/model"
	"hexagonal-architecture/internal/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookControllers struct {
	Lib repositories.BookContract
}

type BookControllersInterface interface {
	GetAllBookControllers(c echo.Context) error
	GetBookByIdControllers(c echo.Context) error
	PostBookControllers(c echo.Context) error
	UpdateBookControllers(c echo.Context) error
	DeleteBookControllers(c echo.Context) error
}

func InitBook(Lib repositories.BookContract) BookControllersInterface {
	return &BookControllers{Lib}
}

func (b *BookControllers) GetAllBookControllers(c echo.Context) error {
	data, err := b.Lib.GetAllBooks()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func (b *BookControllers) GetBookByIdControllers(c echo.Context) error {
	id := c.Param("id")

	data, err := b.Lib.GetBooksById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func (b *BookControllers) PostBookControllers(c echo.Context) error {
	var data model.Book

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	book, err := b.Lib.PostBook(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create book",
		"msg":    book,
	})
}

func (b *BookControllers) UpdateBookControllers(c echo.Context) error {
	id := c.Param("id")
	var data model.Book

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}

	_, err := b.Lib.UpdateBook(id, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update book",
		"msg":    data,
	})
}

func (b *BookControllers) DeleteBookControllers(c echo.Context) error {
	id := c.Param("id")

	_, err := b.Lib.DeleteBook(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete book",
	})
}
