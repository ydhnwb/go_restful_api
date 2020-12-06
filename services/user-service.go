package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

//UserService is a contract interface
type UserService interface {
	Update(user dto.UserUpdateDTO) entities.User
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

func (service *userService) Update(user dto.UserUpdateDTO) entities.User {
	userToUpdate := entities.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Delete(user entities.User) {
	service.userRepository.DeleteUser(user)
}

func (service *userService) Profile(userID string) entities.User {
	return service.userRepository.ProfileUser(userID)
}
