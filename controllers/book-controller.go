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

//BookController is an interface of logic what can BookController do
type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	getUserIDByGivenToken(token string) string
}

type bookController struct {
	service    services.BookService
	jwtService services.JWTService
}

//NewBookController function in creating a new BookController instance
func NewBookController(service services.BookService, jwtService services.JWTService) BookController {
	return &bookController{
		service:    service,
		jwtService: jwtService,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entities.Book = c.service.All()
	response := entities.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, response)
}

func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := entities.BuildErrorResponse("No parameter id were found", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	var book entities.Book = c.service.FindByID(id)
	if book.ID == 0 {
		response := entities.BuildErrorResponse("Id not found", "No data with gived id", entities.EmptyObj{})
		context.JSON(http.StatusNotFound, response)
	} else {
		response := entities.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, response)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	err := context.ShouldBind(&bookCreateDTO)

	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data.", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByGivenToken(authHeader)
		userIDConverted, _err := strconv.ParseUint(userID, 10, 64)
		if _err == nil {
			bookCreateDTO.UserID = userIDConverted
		}
		res := c.service.Insert(bookCreateDTO)
		response := entities.BuildResponse(true, "OK", res)
		fmt.Printf("%v", response)
		context.JSON(http.StatusCreated, response)
	}

}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	err := context.ShouldBind(&bookUpdateDTO)
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
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.service.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		res := c.service.Update(bookUpdateDTO)
		response := entities.BuildResponse(true, "OK", res)
		context.JSON(http.StatusOK, response)
	} else {
		response := entities.BuildErrorResponse("You don't have permission", "You are not the owner of this book", entities.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) Delete(context *gin.Context) {
	var book entities.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to find your id", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	book.ID = id
	c.service.Delete(book)
	response := entities.BuildResponse(true, "Deleted", entities.EmptyObj{})
	context.JSON(http.StatusOK, response)
}

func (c *bookController) getUserIDByGivenToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
