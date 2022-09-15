package routes

import (
	"alterra-agmc-day2/config"
	"alterra-agmc-day2/controllers"
	"alterra-agmc-day2/lib/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Connection = config.InitDB()
	LibUser    = database.InitUser(Connection)
	LibBook    = database.InitBook(Connection)
)

func Init() *echo.Echo {
	r := echo.New()

	userControllers := controllers.InitUser(LibUser)
	bookControllers := controllers.InitBook(LibBook)

	// no need authorization
	r.POST("/login", userControllers.LoginUserControllers)
	r.POST("/users", userControllers.CreateUserControllers)
	r.GET("/books", bookControllers.GetAllBookControllers)
	r.GET("/books/:id", bookControllers.GetBookByIdControllers)

	//need authorization
	jwt := r.Group("/jwt")
	jwt.Use(middleware.JWT([]byte(config.LoadEnv().GetString("JWT_KEY"))))
	jwt.POST("/books", bookControllers.PostBookControllers)
	jwt.GET("/users", userControllers.GetUsersControllers)
	jwt.GET("/users/:id", userControllers.GetUserByIdControllers)
	jwt.PUT("/users/:id", userControllers.UpdateUserControllers)
	jwt.PUT("/books/:id", bookControllers.UpdateBookControllers)
	jwt.DELETE("/users/:id", userControllers.DeletedUserControllers)
	jwt.DELETE("/books/:id", bookControllers.DeleteBookControllers)

	return r
}
