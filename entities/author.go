package entities

// Author struct is used for Video author
type Author struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname" binding:"required"`
}
