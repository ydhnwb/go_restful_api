package entities

//User is represents users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname" binding:"required"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email" binding:"required,email"`
	Password string `gorm:"->;<-:create;not null" json:"password" binding:"required"`
	Token    string `gorm:"-" binding:"-" json:"-"`
}

//UserResponse is a different shape to hide some important value
type UserResponse struct {
	ID       uint64 `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
