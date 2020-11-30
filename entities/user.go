package entities

//User is represents users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname" binding:"required"`
	Email    string `gorm:"type:varchar(255)" json:"email" binding:"required,email"`
	Password string `gorm:"->:false;<-;not null" json:"password" binding:"required"`
}
