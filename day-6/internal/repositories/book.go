package repositories

import (
	"hexagonal-architecture/internal/model"

	"gorm.io/gorm"
)

type BookLib struct {
	DB *gorm.DB
}

type BookContract interface {
	GetAllBooks() (*[]model.Book, error)
	GetBooksById(string) (*model.Book, error)
	PostBook(model.Book) (*model.Book, error)
	UpdateBook(string, model.Book) (*model.Book, error)
	DeleteBook(string) (*model.Book, error)
}

func InitBook(DB *gorm.DB) BookContract {
	return &BookLib{DB}
}

func (b *BookLib) GetAllBooks() (data *[]model.Book, err error) {
	if err := b.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookLib) GetBooksById(id string) (data *model.Book, err error) {
	if err := b.DB.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookLib) PostBook(book model.Book) (*model.Book, error) {
	if err := b.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *BookLib) UpdateBook(id string, bookUpdate model.Book) (data *model.Book, err error) {
	book, err := b.GetBooksById(id)

	if err != nil {
		return nil, err
	}

	b.DB.Model(&book).Updates(&bookUpdate)

	return data, nil
}

func (b *BookLib) DeleteBook(id string) (data *model.Book, err error) {
	book, err := b.GetBooksById(id)

	if err != nil {
		return nil, err
	}

	b.DB.Delete(&book)

	return data, nil
}
