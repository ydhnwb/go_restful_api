package entities

import (
	"time"

	"github.com/ydhnwb/go_restful_api/dto"
)

// Book represents Books table from database
type Book struct {
	ID          uint64          `gorm:"primary_key:auto_increment" json:"id"`
	Title       string          `gorm:"type:varchar(255)" json:"title"`
	Description string          `gorm:"type:text" json:"description"`
	UserID      uint64          `json:"-,omitempty"`
	User        dto.UserReadDTO `json:"-" gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time       `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
