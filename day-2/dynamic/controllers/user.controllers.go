package controllers

import (
	"alterra-agmc-day2/lib/database"
	"alterra-agmc-day2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	Lib database.UserContract
}

type UserControllerInterface interface {
	GetUsersControllers(echo.Context) error
	CreateUserControllers(echo.Context) error
	GetUserByIdControllers(c echo.Context) error
	UpdateUserControllers(c echo.Context) error
	DeletedUserControllers(c echo.Context) error
}

func Init(Lib database.UserContract) UserControllerInterface {
	return &UserControllers{Lib}
}

func (u UserControllers) CreateUserControllers(c echo.Context) error {
	var data models.User

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	u.Lib.CreateUser(data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func (u UserControllers) GetUsersControllers(c echo.Context) error {
	data, err := u.Lib.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func (u UserControllers) GetUserByIdControllers(c echo.Context) error {
	id := c.Param("id")
	data, err := u.Lib.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func (u UserControllers) UpdateUserControllers(c echo.Context) error {
	id := c.Param("id")
	var data models.User

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	u.Lib.UpdateUser(id, data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": data,
	})
}

func (u UserControllers) DeletedUserControllers(c echo.Context) error {
	id := c.Param("id")
	_, err := u.Lib.DeletedUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success delete user",
	})
}
