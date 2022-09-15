package controllers

import (
	"log"
	"net/http"
	"static/middleware"
	"static/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var books []models.Book

func GetAllBook(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": books,
	})
}

func GetBookById(c echo.Context) error {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Can't convert to int")
	}

	for i, v := range books {
		if v.Id == uint(idConv) {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"msg": books[i],
			})
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"msg": "Data not found",
	})
}

func CreateBook(c echo.Context) error {
	var data models.Book

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	books = append(books, data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Can't convert to int")
	}

	var data models.Book

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	for i, v := range books {
		if v.Id == uint(idConv) {
			books[i].UpdatedAt = time.Now()
			books[i].Title = data.Title
			books[i].ISBN = data.ISBN
			books[i].Writer = data.Writer
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update books",
		"msg":    books,
	})

}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Can't convert to int")
	}

	for i, v := range books {
		if v.Id == uint(idConv) {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"msg": "success delete book",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"msg": "book not found",
	})
}

func GenerateJWT(c echo.Context) error {
	var id uint
	for i := range books {
		id = books[i].Id
	}

	convId := strconv.Itoa(int(id))

	token, err := middleware.CreateToken(convId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
