package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//LoginController is a contract
type LoginController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
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
	var credentials dto.LoginDTO
	err := context.ShouldBind(&credentials)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	isAuthenticated := controller.loginService.VerifyCredential(credentials.Email, credentials.Password)
	if isAuthenticated {
		user := controller.loginService.FindByEmail(credentials.Email)
		generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(user.ID, 10), false)
		userResponse := entities.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Fullname: user.Fullname,
			Token:    generatedToken,
		}
		response := entities.BuildResponse(true, "OK!", userResponse)
		context.JSON(http.StatusOK, response)
	} else {
		response := entities.BuildErrorResponse("Cannot authenticate! Check again your credentials", "Invalid credentials", nil)
		context.JSON(http.StatusUnauthorized, response)
	}
}

//Register is creates a new user
func (controller *loginController) Register(context *gin.Context) {
	var user entities.User
	err := context.ShouldBind(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		if !controller.loginService.IsDuplicateEmail(user) {
			response := entities.BuildErrorResponse("Failed to process your data", "Duplicate email", nil)
			context.JSON(http.StatusConflict, response)
		} else {
			createdUser := controller.loginService.CreateUser(user)
			userResponse := entities.UserResponse{
				ID:       createdUser.ID,
				Fullname: createdUser.Fullname,
				Email:    createdUser.Email,
				Token:    createdUser.Token,
			}
			response := entities.BuildResponse(true, "OK!", userResponse)
			context.JSON(http.StatusCreated, response)
		}
	}
}
