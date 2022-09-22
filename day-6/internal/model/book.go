package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Writer string `json:"writer"`
}
