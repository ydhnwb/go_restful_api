package dto

//RegisterDTO is a serializer when user creates an account
type RegisterDTO struct {
	Fullname string `form:"fullname" json:"fullname" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
