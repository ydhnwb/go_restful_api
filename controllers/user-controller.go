package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//UserController ...
type UserController interface {
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

func (c *userController) Update(context *gin.Context) {
	var user dto.UserUpdateDTO
	err := context.ShouldBind(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["name"]), 10, 64)
	user.ID = id
	u := c.userService.Update(user)
	userToReturn := dto.UserReadDTO{}
	errMap := smapping.FillStruct(&userToReturn, smapping.MapFields(&u))
	if errMap != nil {
		log.Fatalf("failed map: %v", err)
	}
	response := entities.BuildResponse(true, "OK", userToReturn)
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
	userToReturn := dto.UserReadDTO{}
	errMap := smapping.FillStruct(&userToReturn, smapping.MapFields(&user))
	if errMap != nil {
		log.Fatalf("failed map: %v", err)
	}
	response := entities.BuildResponse(true, "OK", userToReturn)
	context.JSON(http.StatusOK, response)
}
