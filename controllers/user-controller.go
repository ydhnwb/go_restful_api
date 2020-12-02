package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//UserController ...
type UserController interface {
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

//NewUserController is a
func NewUserController(service services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: service,
		jwtService:  jwtService,
	}
}

func (c *userController) Insert(context *gin.Context) {
	var user entities.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		//also handle hash password here
		c.userService.Insert(user)
		response := entities.BuildResponse(true, "OK", user)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *userController) Update(context *gin.Context) {
	var user entities.User
	err := context.ShouldBind(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	response := entities.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, response)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["name"]))
	user.Password = ""
	response := entities.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, response)
}
