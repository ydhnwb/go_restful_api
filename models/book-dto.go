package dto

import (
	"time"

	"github.com/ydhnwb/go_restful_api/entities"
)

//BookDTO is a model for general serializer of Book
type BookDTO struct {
	ID          uint64        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	UserID      uint64        `json:"-"`
	User        entities.User `json:"user"`
	CreatedAt   time.Time     `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
