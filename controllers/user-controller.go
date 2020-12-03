package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/dto"
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
	var user dto.UserUpdateDTO
	err := context.ShouldBind(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["name"] != strconv.FormatUint(user.ID, 10) {
		response := entities.BuildErrorResponse("Your id or token didn't match", "Token or ID didnt match", nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		var userToUpdate entities.User = entities.User{
			Email:    user.Email,
			Fullname: user.Fullname,
			ID:       user.ID,
		}
		if user.Password != "" {
			userToUpdate.Password = user.Password
		}
		c.userService.Update(userToUpdate)
		userToUpdate.Password = ""
		response := entities.BuildResponse(true, "OK", userToUpdate)
		context.JSON(http.StatusOK, response)
	}

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
