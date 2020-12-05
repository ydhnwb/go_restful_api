package dto

//LoginDTO converts request body data to an object
type LoginDTO struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
