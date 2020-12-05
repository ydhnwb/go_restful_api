package services

import (
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

//UserService is a contract interface
type UserService interface {
	Insert(user dto.UserCreateDTO) entities.User
	Update(user dto.UserUpdateDTO) dto.UserUpdateDTO
	Delete(user entities.User)
	Profile(userID string) entities.User
}

type userService struct {
	userRepository repositories.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Insert(user dto.UserCreateDTO) entities.User {
	createdUser := service.userRepository.InsertUser(user)
	return createdUser
}

func (service *userService) Update(user dto.UserUpdateDTO) dto.UserUpdateDTO {
	service.userRepository.UpdateUser(user)
	return user
}

func (service *userService) Delete(user entities.User) {
	service.userRepository.DeleteUser(user)
}

func (service *userService) Profile(userID string) entities.User {
	return service.userRepository.ProfileUser(userID)
}
