package entities

import (
	"time"
)

// Book represents Books table from database
// User        User      `gorm:"->;<-:false;foreignkey:UserID" json:"-"`
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title" binding:"min=1,max=255,required"`
	Description string `gorm:"type:text" json:"description" binding:"required"`
	// UserID      uint64    `gorm:"->;<-:create;foreignkey:ID;constraint:OnDelete:CASCADE;" json:"-"`
	UserID uint64 `gorm:"not null" json:"-"`
	User   User   `json:"-,omitempty"`
	// User        User      `json:"user"`
	CreatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
