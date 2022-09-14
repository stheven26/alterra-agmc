package database

import (
	"alterra-agmc-day2/config"
	"alterra-agmc-day2/models"

	"gorm.io/gorm"
)

type UserLib struct {
	DB *gorm.DB
}

type UserContract interface {
	CreateUser(models.User) (*models.User, error)
	GetUsers() (*[]models.User, error)
	GetUserById(string) (*models.User, error)
	UpdateUser(string, models.User) (*models.User, error)
	DeletedUser(string) (*models.User, error)
}

func Init(DB *gorm.DB) UserContract {
	return &UserLib{DB}
}

func (u *UserLib) CreateUser(user models.User) (*models.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserLib) GetUsers() (data *[]models.User, err error) {
	if err := u.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserLib) GetUserById(id string) (data *models.User, err error) {
	conn := config.InitDB()

	if err = conn.Where(`id=?`, id).First(&data).Error; err != nil {
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

	conn := config.InitDB()

	if err = conn.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	conn.Delete(&data)

	return data, nil
}
