package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
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
	authResult := controller.loginService.VerifyCredential(credentials.Email, credentials.Password)
	if v, ok := authResult.(entities.User); ok {
		generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10), false)
		userToReturn := dto.UserReadDTO{}
		err = smapping.FillStruct(&userToReturn, smapping.MapFields(&v))
		userToReturn.Token = generatedToken
		response := entities.BuildResponse(true, "OK!", userToReturn)
		context.JSON(http.StatusOK, response)
		return
	}
	response := entities.BuildErrorResponse("Cannot authenticate! Check again your credentials", "Invalid credentials", nil)
	context.JSON(http.StatusUnauthorized, response)
}

//Register is creates a new user
func (controller *loginController) Register(context *gin.Context) {
	var user dto.UserCreateDTO
	err := context.ShouldBind(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		if !controller.loginService.IsDuplicateEmail(user.Email) {
			response := entities.BuildErrorResponse("Failed to process your data", "Duplicate email", nil)
			context.JSON(http.StatusConflict, response)
		} else {
			createdUser := controller.loginService.CreateUser(user)
			token := controller.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10), false)
			userToReturn := dto.UserReadDTO{}
			err = smapping.FillStruct(&userToReturn, smapping.MapFields(&createdUser))
			userToReturn.Token = token
			if err != nil {
				log.Fatalf("failed map: %v", err)
			}
			response := entities.BuildResponse(true, "OK!", userToReturn)
			context.JSON(http.StatusCreated, response)
		}
	}
}
