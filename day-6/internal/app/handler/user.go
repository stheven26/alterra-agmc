package handler

import (
	"hexagonal-architecture/internal/model"
	"hexagonal-architecture/internal/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	Lib repositories.UserContract
}

type UserControllerInterface interface {
	GetUsersControllers(echo.Context) error
	CreateUserControllers(echo.Context) error
	GetUserByIdControllers(c echo.Context) error
	UpdateUserControllers(c echo.Context) error
	DeletedUserControllers(c echo.Context) error
	LoginUserControllers(c echo.Context) error
}

func InitUser(Lib repositories.UserContract) UserControllerInterface {
	return &UserControllers{Lib}
}

func (u *UserControllers) CreateUserControllers(c echo.Context) error {
	var data model.User

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

func (u *UserControllers) GetUsersControllers(c echo.Context) error {
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

func (u *UserControllers) GetUserByIdControllers(c echo.Context) error {
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

func (u *UserControllers) UpdateUserControllers(c echo.Context) error {
	id := c.Param("id")
	var data model.User

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	_, err := u.Lib.UpdateUser(id, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

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

func (u UserControllers) LoginUserControllers(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	users, err := u.Lib.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":   "Success Login",
		"users": users,
	})
}
