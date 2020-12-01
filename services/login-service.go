package services

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

// LoginService is an interface about something that this service can do
type LoginService interface {
	VerifyCredential(email string, password string) bool
	CreateUser(user entities.User) entities.User
	FindByEmail(email string) entities.User
	IsDuplicateEmail(user entities.User) bool
}

type loginService struct {
	userRepository repositories.UserRepository
}

// NewLoginService is creating a new instance of LoginService
func NewLoginService(userRepo repositories.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepo,
	}
}

func (service *loginService) VerifyCredential(email string, password string) bool {
	return service.userRepository.VerifyCredential(email, password)

}

func (service *loginService) CreateUser(user entities.User) entities.User {
	res := service.userRepository.InsertUser(user)
	return res
}

func (service *loginService) IsDuplicateEmail(user entities.User) bool {
	return service.userRepository.IsDuplicateEmail(user)
}

func (service *loginService) FindByEmail(email string) entities.User {
	return service.userRepository.FindByEmail(email)
}
