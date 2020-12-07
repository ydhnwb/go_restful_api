package dto

import (
	"time"
)

//BookReadDTO is a model for general serializer of Book
type BookReadDTO struct {
	ID          uint64      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	User        UserReadDTO `json:"user" gorm:"foreignKey:ID"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

//BookUpdateDTO is a model for book update serializer
type BookUpdateDTO struct {
	ID          uint64 `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

//BookCreateDTO is a model serializer for creating Book
type BookCreateDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"`
}
