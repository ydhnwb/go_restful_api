package entities

import "time"

// Person struct is used for Video author
type Person struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Firstname string `gorm:"type:varchar(100)" json:"firstname" binding:"required"`
	Lastname  string `gorm:"type:varchar(100)" json:"lastname"`
	Email     string `gorm:"type:varchar(255)" json:"email" binding:"required,email"`
}

// Video is a data class
type Video struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title" binding:"min=1,max=255"`
	Description string    `gorm:"type:text" json:"description" binding:"max=255"`
	URL         string    `gorm:"type:text" json:"url" binding:"required,url"`
	Author      Person    `gorm:"foreignkey:PersonID" json:"author" binding:"required"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
