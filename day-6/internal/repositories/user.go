package repositories

import (
	"hexagonal-architecture/internal/middlewares"
	"hexagonal-architecture/internal/model"
	"strconv"

	"gorm.io/gorm"
)

type UserLib struct {
	DB *gorm.DB
}

type UserContract interface {
	CreateUser(model.User) (*model.User, error)
	GetUsers() ([]model.User, error)
	GetUserById(string) (*model.User, error)
	UpdateUser(string, model.User) (*model.User, error)
	DeletedUser(string) (*model.User, error)
	LoginUser(model.User) (*string, error)
}

func InitUser(DB *gorm.DB) UserContract {
	return &UserLib{DB}
}

func (u *UserLib) CreateUser(user model.User) (*model.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserLib) GetUsers() (data []model.User, err error) {
	if err := u.DB.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (u *UserLib) GetUserById(id string) (data *model.User, err error) {

	if err = u.DB.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserLib) UpdateUser(id string, dataUpdate model.User) (data *model.User, err error) {

	user, err := u.GetUserById(id)

	if err != nil {
		return nil, err
	}

	u.DB.Model(&user).Updates(&dataUpdate)

	return data, nil
}

func (u *UserLib) DeletedUser(id string) (data *model.User, err error) {
	user, err := u.GetUserById(id)

	if err != nil {
		return nil, err
	}

	u.DB.Delete(&user)

	return data, nil
}

func (u *UserLib) LoginUser(data model.User) (*string, error) {
	if err := u.DB.Where(`email=? AND password=?`, data.Email, data.Password).Error; err != nil {
		return nil, err
	}

	convID := strconv.Itoa(int(data.ID))

	token, err := middlewares.CreateToken(convID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
