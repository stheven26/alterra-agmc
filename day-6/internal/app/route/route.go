package route

import (
	"hexagonal-architecture/database"
	"hexagonal-architecture/internal/app/handler"
	"hexagonal-architecture/internal/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Connection = database.InitDB()
	LibUser    = repositories.InitUser(Connection)
	LibBook    = repositories.InitBook(Connection)
)

func Init() *echo.Echo {
	r := echo.New()
	jwtMiddleware := middleware.JWT([]byte(database.LoadEnv().GetString("JWT_SECRET")))
	userControllers := handler.InitUser(LibUser)
	bookControllers := handler.InitBook(LibBook)

	// no need authorization
	r.POST("/login", userControllers.LoginUserControllers)
	r.POST("/users", userControllers.CreateUserControllers)
	r.GET("/books", bookControllers.GetAllBookControllers)
	r.GET("/books/:id", bookControllers.GetBookByIdControllers)

	//need authorization
	jwt := r.Group("/jwt", jwtMiddleware)
	jwt.POST("/books", bookControllers.PostBookControllers)
	jwt.GET("/users", userControllers.GetUsersControllers)
	jwt.GET("/users/:id", userControllers.GetUserByIdControllers)
	jwt.PUT("/users/:id", userControllers.UpdateUserControllers)
	jwt.PUT("/books/:id", bookControllers.UpdateBookControllers)
	jwt.DELETE("/users/:id", userControllers.DeletedUserControllers)
	jwt.DELETE("/books/:id", bookControllers.DeleteBookControllers)

	return r
}
