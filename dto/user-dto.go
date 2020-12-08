package dto

//UserUpdateDTO is a model for updating user
type UserUpdateDTO struct {
	ID       uint64 `json:"id" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty"`
}

//UserCreateDTO is a model for creating user
type UserCreateDTO struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
