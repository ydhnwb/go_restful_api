package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
}

type controller struct {
	service services.BookService
}

//NewBookController function in creating a new BookController instance
func NewBookController(service services.BookService) BookController {
	return &controller{
		service: service,
	}
}

func (c *controller) All(context *gin.Context) {
	var books []entities.Book = c.service.All()
	response := entities.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, response)
}

func (c *controller) FindByID(context *gin.Context) {
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

func (c *controller) Insert(context *gin.Context) {
	var book entities.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to process your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		c.service.Insert(book)
		response := entities.BuildResponse(true, "OK", book)
		context.JSON(http.StatusCreated, response)
	}

}

func (c *controller) Update(context *gin.Context) {
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

func (c *controller) Delete(context *gin.Context) {
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