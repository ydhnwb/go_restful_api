package entities

//User is represents users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname" binding:"required"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email" binding:"required,email"`
	Password string `gorm:"->;<-;not null" json:"password,omitempty" binding:"required"`
	Token    string `gorm:"-" binding:"-" json:"token,omitempty"`
}

//UserResponse is a different shape to hide some important value
type UserResponse struct {
	ID       uint64 `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token,omitempty"`
}

//UserCreate is a serializer to used when creating/register a new User
type UserCreate struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
}
