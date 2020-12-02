package dto

//UserUpdateDTO is a model for updating user
type UserUpdateDTO struct {
	ID       uint64 `json:"id" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty"`
}
