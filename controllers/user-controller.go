package controllers

import (
	"net/http"

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
}

//NewUserController is a
func NewUserController(service services.UserService) UserController {
	return &userController{
		userService: service,
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
	err := context.ShouldBindJSON(&user)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	response := entities.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, response)
}

func (c *userController) Profile(context *gin.Context) {
	println("Get user profile")
}
