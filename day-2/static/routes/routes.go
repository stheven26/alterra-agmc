package routes

import (
	"static/controllers"

	"github.com/labstack/echo/v4"
)

func App() *echo.Echo {
	r := echo.New()

	books := r.Group("/books")

	books.GET("", controllers.GetAllBook)
	books.GET("/:id", controllers.GetBookById)
	books.POST("", controllers.CreateBook)
	books.PUT("/:id", controllers.UpdateBook)
	books.DELETE("/:id", controllers.DeleteBook)

	return r
}
