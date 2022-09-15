package database

import (
	"alterra-agmc-day2/models"

	"gorm.io/gorm"
)

type BookLib struct {
	DB *gorm.DB
}

type BookContract interface {
	GetAllBooks() (*[]models.Book, error)
	GetBooksById(string) (*models.Book, error)
	PostBook(models.Book) (*models.Book, error)
	UpdateBook(string, models.Book) (*models.Book, error)
	DeleteBook(string) (*models.Book, error)
}

func InitBook(DB *gorm.DB) BookContract {
	return &BookLib{DB}
}

func (b *BookLib) GetAllBooks() (data *[]models.Book, err error) {
	if err := b.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookLib) GetBooksById(id string) (data *models.Book, err error) {
	if err := b.DB.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookLib) PostBook(book models.Book) (*models.Book, error) {
	if err := b.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *BookLib) UpdateBook(id string, bookUpdate models.Book) (data *models.Book, err error) {
	book, err := b.GetBooksById(id)

	if err != nil {
		return nil, err
	}

	b.DB.Model(&book).Updates(&bookUpdate)

	return data, nil
}

func (b *BookLib) DeleteBook(id string) (data *models.Book, err error) {
	book, err := b.GetBooksById(id)

	if err != nil {
		return nil, err
	}

	b.DB.Delete(&book)

	return data, nil
}
