package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/services"
)

//UserController ...
type UserController interface {
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
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

}

func (c *userController) Update(context *gin.Context) {

}

func (c *userController) Delete(context *gin.Context) {

}

func (c *userController) Profile(context *gin.Context) {

}
