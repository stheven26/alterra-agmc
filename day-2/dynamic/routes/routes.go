package routes

import (
	"alterra-agmc-day2/config"
	"alterra-agmc-day2/controllers"
	"alterra-agmc-day2/lib/database"

	"github.com/labstack/echo/v4"
)

var (
	Connection = config.InitDB()
	LibUser    = database.Init(Connection)
)

func Init() *echo.Echo {
	e := echo.New()
	routes := e.Group("/api/v1/users")

	userController := controllers.Init(LibUser)
	routes.POST("", userController.CreateUserControllers)
	routes.GET("", userController.GetUsersControllers)
	routes.GET("/:id", userController.GetUserByIdControllers)
	routes.PUT("/:id", userController.UpdateUserControllers)
	routes.DELETE("/:id", userController.DeletedUserControllers)

	return e
}
