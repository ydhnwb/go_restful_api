package services

// LoginService is an interface about something that this service can do
type LoginService interface {
	Login(email string, password string) bool
}

type loginService struct {
	authorizedEmail    string
	authorizedPassword string
}

// NewLoginService is creating a new instance of LoginService
func NewLoginService() LoginService {
	return &loginService{
		authorizedEmail:    "example@gmail.com",
		authorizedPassword: "example",
	}
}

func (service *loginService) Login(email string, password string) bool {
	status := service.authorizedEmail == email && service.authorizedPassword == password
	return status
}
