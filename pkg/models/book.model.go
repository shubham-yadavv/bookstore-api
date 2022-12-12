package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}
