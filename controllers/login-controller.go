package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//LoginController is a contract
type LoginController interface {
	Login(context *gin.Context)
}

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

//NewLoginController creates an instane LoginController
func NewLoginController(loginService services.LoginService, jwtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(context *gin.Context) {
	//mocking from database login
	// var credentials dto.Credentials
	// err := ctx.ShouldBind(&credentials)
	// if err != nil {
	// 	return ""
	// }
	email := "example@gmail.com"
	password := "example"
	isAuthenticated := controller.loginService.Login(email, password)
	if isAuthenticated {
		generatedToken := controller.jwtService.GenerateToken(email, false)
		user := entities.User{
			ID:    1,
			Email: email,
			Name:  "Udin gambut",
			Token: generatedToken,
		}
		response := entities.BuildResponse(true, "OK!", user)
		context.JSON(http.StatusOK, response)
	} else {
		response := entities.BuildErrorResponse("Cannot authenticate! Check again your credentials", "Error\nErro", nil)
		context.JSON(http.StatusUnauthorized, response)
	}
}
