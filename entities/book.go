package entities

import (
	"time"
)

// Book represents Books table from database
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title" binding:"min=1,max=255"`
	Description string `gorm:"type:text" json:"description" binding:"max=255"`
	// User        User      `gorm:"->;foreignkey:UserID" json:"user" binding:"-"`
	// UserID      uint64    `gorm:"->:false;<-" json:"user_id" binding:"required"`
	CreatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
