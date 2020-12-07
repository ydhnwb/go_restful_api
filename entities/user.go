package entities

//User is represents users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" binding:"-" json:"token,omitempty"`
}
