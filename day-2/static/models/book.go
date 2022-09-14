package models

import "time"

type Book struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	ISBN      string    `json:"isbn"`
	Writer    string    `json:"writer"`
}
