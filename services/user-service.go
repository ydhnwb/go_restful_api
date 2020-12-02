package services

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

//UserService is a contract interface
type UserService interface {
	Insert(user entities.User) entities.User
	Update(user entities.User) entities.User
	Delete(user entities.User)
	Profile(userId string) entities.User
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

func (service *userService) Insert(user entities.User) entities.User {
	service.userRepository.InsertUser(user)
	return user
}

func (service *userService) Update(user entities.User) entities.User {
	service.userRepository.UpdateUser(user)
	return user
}

func (service *userService) Delete(user entities.User) {
	service.userRepository.DeleteUser(user)
}

func (service *userService) Profile(userID string) entities.User {
	return service.userRepository.ProfileUser(userID)
}
