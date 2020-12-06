package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
	"golang.org/x/crypto/bcrypt"
)

// LoginService is an interface about something that this service can do
type LoginService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entities.User
	FindByEmail(email string) entities.User
	IsDuplicateEmail(email string) bool
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

func (service *loginService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entities.User); ok {
		comparedPassword := comparePasswords(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *loginService) CreateUser(user dto.UserCreateDTO) entities.User {
	userToCreate := entities.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *loginService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *loginService) FindByEmail(email string) entities.User {
	return service.userRepository.FindByEmail(email)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
