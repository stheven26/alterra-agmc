package database

import (
	"alterra-agmc-day2/middlewares"
	"alterra-agmc-day2/models"
	"strconv"

	"gorm.io/gorm"
)

type UserLib struct {
	DB *gorm.DB
}

type UserContract interface {
	CreateUser(models.User) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUserById(string) (*models.User, error)
	UpdateUser(string, models.User) (*models.User, error)
	DeletedUser(string) (*models.User, error)
	LoginUser(models.User) (*string, error)
}

func InitUser(DB *gorm.DB) UserContract {
	return &UserLib{DB}
}

func (u *UserLib) CreateUser(user models.User) (*models.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserLib) GetUsers() (data []models.User, err error) {
	if err := u.DB.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (u *UserLib) GetUserById(id string) (data *models.User, err error) {

	if err = u.DB.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserLib) UpdateUser(id string, dataUpdate models.User) (data *models.User, err error) {

	user, err := u.GetUserById(id)

	if err != nil {
		return nil, err
	}

	u.DB.Model(&user).Updates(&dataUpdate)

	return data, nil
}

func (u *UserLib) DeletedUser(id string) (data *models.User, err error) {
	user, err := u.GetUserById(id)

	if err != nil {
		return nil, err
	}

	u.DB.Delete(&user)

	return data, nil
}

func (u *UserLib) LoginUser(data models.User) (*string, error) {
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
