package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
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
	// booksDTO := convertBookEntitiesToDTO(books)
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
	if (entities.Book{}) == book {
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
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByGivenToken(authHeader)
	userIDConverted, _err := strconv.ParseUint(userID, 10, 64)
	if _err == nil {
		bookCreateDTO.UserID = userIDConverted
	}
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data.", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		res := c.service.Insert(bookCreateDTO)
		response := entities.BuildResponse(true, "OK", res)
		context.JSON(http.StatusCreated, response)
	}

}

func (c *bookController) Update(context *gin.Context) {
	var book entities.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to find your id", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	book.ID = id
	c.service.Update(book)
	response := entities.BuildResponse(true, "OK", book)
	context.JSON(http.StatusOK, response)
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
	response := entities.BuildResponse(true, "Deleted", nil)
	context.JSON(http.StatusOK, response)
}

func (c *bookController) getUserIDByGivenToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["name"])
}

func convertBookEntitiesToDTO(books []entities.Book) []dto.BookReadDTO {
	var booksDTO []dto.BookReadDTO
	for i, v := range books {
		print(i)
		user := dto.UserReadDTO{}
		smapping.FillStruct(&user, smapping.MapFields(&v.User))
		book := dto.BookReadDTO{}
		smapping.FillStruct(&book, smapping.MapFields(&v))
		book.User = user
		booksDTO = append(booksDTO, book)
		fmt.Printf("%v \n", v)
	}
	return booksDTO
}
