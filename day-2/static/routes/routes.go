package routes

import (
	"os"
	"static/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func App() *echo.Echo {
	r := echo.New()

	r.POST("/generateToken", controllers.GenerateJWT)
	r.GET("/books", controllers.GetAllBook)
	r.GET("/books/:id", controllers.GetBookById)

	jwt := r.Group("/jwt")
	jwt.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	jwt.POST("/books", controllers.CreateBook)
	jwt.PUT("/books/:id", controllers.UpdateBook)
	jwt.DELETE("/books/:id", controllers.DeleteBook)

	return r
}
